package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupTipoTransaccionRoutes(r *mux.Router) {
	r.HandleFunc("/tipo-transaccion", controllers.GetAllTipoTransaccion).Methods("GET")
	r.HandleFunc("/tipo-transaccion/{id}", controllers.GetTipoTransaccionByID).Methods("GET")
	r.HandleFunc("/tipo-transaccion", controllers.CreateTipoTransaccion).Methods("POST")
	r.HandleFunc("/tipo-transaccion/{id}", controllers.UpdateTipoTransaccion).Methods("PUT")
	r.HandleFunc("/tipo-transaccion/{id}", controllers.DeleteTipoTransaccion).Methods("DELETE")
}
