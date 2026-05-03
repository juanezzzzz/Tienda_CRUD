package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupMetodosPagoGuardadosRoutes(r *mux.Router) {
	r.HandleFunc("/metodos-pago-guardados", controllers.GetAllMetodosPagoGuardados).Methods("GET")
	r.HandleFunc("/metodos-pago-guardados/{id}", controllers.GetMetodosPagoGuardadosByID).Methods("GET")
	r.HandleFunc("/metodos-pago-guardados", controllers.CreateMetodosPagoGuardados).Methods("POST")
	r.HandleFunc("/metodos-pago-guardados/{id}", controllers.UpdateMetodosPagoGuardados).Methods("PUT")
	r.HandleFunc("/metodos-pago-guardados/{id}", controllers.DeleteMetodosPagoGuardados).Methods("DELETE")
}
