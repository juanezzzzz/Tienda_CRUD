package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupTransaccionesRoutes(r *mux.Router) {
	r.HandleFunc("/transacciones", controllers.GetAllTransacciones).Methods("GET")
	r.HandleFunc("/transacciones/{id}", controllers.GetTransaccionesByID).Methods("GET")
	r.HandleFunc("/transacciones", controllers.CreateTransacciones).Methods("POST")
	r.HandleFunc("/transacciones/{id}", controllers.UpdateTransacciones).Methods("PUT")
	r.HandleFunc("/transacciones/{id}", controllers.DeleteTransacciones).Methods("DELETE")
}
