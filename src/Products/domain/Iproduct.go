package domain




type IProductRepository interface {
    Save(P *Product) error
    GetAll() ([]Product, error)
    Delete(name string) error  
    Update(nombre string, p *Product) error  
}