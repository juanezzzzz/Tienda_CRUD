package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupCuponesRoutes(r *mux.Router) {
	r.HandleFunc("/cupones", controllers.GetAllCupones).Methods("GET")
	r.HandleFunc("/cupones/{id}", controllers.GetCuponesByID).Methods("GET")
	r.HandleFunc("/cupones", controllers.CreateCupones).Methods("POST")
	r.HandleFunc("/cupones/{id}", controllers.UpdateCupones).Methods("PUT")
	r.HandleFunc("/cupones/{id}", controllers.DeleteCupones).Methods("DELETE")
}
