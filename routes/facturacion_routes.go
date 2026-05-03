package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupFacturacionRoutes(r *mux.Router) {
	r.HandleFunc("/facturacion", controllers.GetAllFacturacion).Methods("GET")
	r.HandleFunc("/facturacion/{id}", controllers.GetFacturacionByID).Methods("GET")
	r.HandleFunc("/facturacion", controllers.CreateFacturacion).Methods("POST")
	r.HandleFunc("/facturacion/{id}", controllers.UpdateFacturacion).Methods("PUT")
	r.HandleFunc("/facturacion/{id}", controllers.DeleteFacturacion).Methods("DELETE")
}
