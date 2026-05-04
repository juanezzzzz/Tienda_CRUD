package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupMetodoPagoRoutes(r *mux.Router) {
	r.HandleFunc("/metodo-pago", controllers.GetAllMetodoPago).Methods("GET")
	r.HandleFunc("/metodo-pago/{id}", controllers.GetMetodoPagoByID).Methods("GET")
	r.HandleFunc("/metodo-pago", controllers.CreateMetodoPago).Methods("POST")
	r.HandleFunc("/metodo-pago/{id}", controllers.UpdateMetodoPago).Methods("PUT")
	r.HandleFunc("/metodo-pago/{id}", controllers.DeleteMetodoPago).Methods("DELETE")
}
