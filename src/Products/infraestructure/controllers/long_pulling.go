package controllers

import (
	"net/http"
	"time"
)


var eventChannel = make(chan string)


func triggerEvent() {
	time.Sleep(5 * time.Second) // Espera 5 segundos
	eventChannel <- "Nuevo evento disponible!" // Envía un evento
}

// Long Polling Handler
func LongPollingHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "text/plain")
	
	// El servidor espera recibir un evento o agotar el tiempo
	select {
	case event := <-eventChannel: // Espera un evento
		w.Write([]byte(event))  // Envía el evento al cliente
	case <-time.After(30 * time.Second): // Timeout de 30 segundos
		w.Write([]byte("Tiempo de espera agotado"))
	}
}
