package services

import (
	"context"
	"sync"
)

// EventService maneja eventos de productos
type EventService struct {
	mu       sync.Mutex
	products []Product
	waiters  []chan []Product
}

// Product estructura para productos
type Product struct {
	Nombre string
	Precio int16
}

// NewEventService crea una nueva instancia del servicio
func NewEventService() *EventService {
	return &EventService{
		products: []Product{},
		waiters:  []chan []Product{},
	}
}

// AddProduct agrega un producto y notifica a los clientes en espera
func (es *EventService) AddProduct(nombre string, precio int16) {
	es.mu.Lock()
	defer es.mu.Unlock()

	product := Product{Nombre: nombre, Precio: precio}
	es.products = append(es.products, product)

	// Notificar solo a los clientes en espera
	for _, ch := range es.waiters {
		ch <- []Product{product} // Solo envía el nuevo producto
	}
	es.waiters = nil // Limpiar la lista de clientes en espera
}

// WaitForProducts espera hasta que haya un nuevo producto disponible
func (es *EventService) WaitForProducts(ctx context.Context) ([]Product, bool) {
	es.mu.Lock()

	// Si ya hay productos, retornarlos de inmediato
	if len(es.products) > 0 {
		products := es.products
		es.products = nil
		es.mu.Unlock()
		return products, true
	}

	// Si no hay productos, se debe esperar un nuevo evento
	ch := make(chan []Product, 1)
	es.waiters = append(es.waiters, ch)
	es.mu.Unlock()

	select {
	case products := <-ch:
		return products, true
	case <-ctx.Done():
		return nil, false
	}
}
