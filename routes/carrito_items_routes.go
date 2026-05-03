package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupCarritoItemsRoutes(r *mux.Router) {
	r.HandleFunc("/carrito-items", controllers.GetAllCarritoItems).Methods("GET")
	r.HandleFunc("/carrito-items/{id}", controllers.GetCarritoItemsByID).Methods("GET")
	r.HandleFunc("/carrito-items", controllers.CreateCarritoItems).Methods("POST")
	r.HandleFunc("/carrito-items/{id}", controllers.UpdateCarritoItems).Methods("PUT")
	r.HandleFunc("/carrito-items/{id}", controllers.DeleteCarritoItems).Methods("DELETE")
}
