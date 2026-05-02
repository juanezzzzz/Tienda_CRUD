package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupProductosRoutes(r *mux.Router) {
	r.HandleFunc("/productos", controllers.GetAllProductos).Methods("GET")
	r.HandleFunc("/productos/{id}", controllers.GetProductosByID).Methods("GET")
	r.HandleFunc("/productos", controllers.CreateProductos).Methods("POST")
	r.HandleFunc("/productos/{id}", controllers.UpdateProductos).Methods("PUT")
	r.HandleFunc("/productos/{id}", controllers.DeleteProductos).Methods("DELETE")
}
