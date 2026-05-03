package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupPagosRoutes(r *mux.Router) {
	r.HandleFunc("/pagos", controllers.GetAllPagos).Methods("GET")
	r.HandleFunc("/pagos/{id}", controllers.GetPagosByID).Methods("GET")
	r.HandleFunc("/pagos", controllers.CreatePagos).Methods("POST")
	r.HandleFunc("/pagos/{id}", controllers.UpdatePagos).Methods("PUT")
	r.HandleFunc("/pagos/{id}", controllers.DeletePagos).Methods("DELETE")
}
