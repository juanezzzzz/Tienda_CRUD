package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupSaldoJobsyRoutes(r *mux.Router) {
	r.HandleFunc("/saldo-jobsy", controllers.GetAllSaldoJobsy).Methods("GET")
	r.HandleFunc("/saldo-jobsy/{id}", controllers.GetSaldoJobsyByID).Methods("GET")
	r.HandleFunc("/saldo-jobsy", controllers.CreateSaldoJobsy).Methods("POST")
	r.HandleFunc("/saldo-jobsy/{id}", controllers.UpdateSaldoJobsy).Methods("PUT")
	r.HandleFunc("/saldo-jobsy/{id}", controllers.DeleteSaldoJobsy).Methods("DELETE")
}
