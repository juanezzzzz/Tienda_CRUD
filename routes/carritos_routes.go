package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupCarritosRoutes(r *mux.Router) {
	r.HandleFunc("/carritos", controllers.GetAllCarritos).Methods("GET")
	r.HandleFunc("/carritos/{id}", controllers.GetCarritosByID).Methods("GET")
	r.HandleFunc("/carritos", controllers.CreateCarritos).Methods("POST")
	r.HandleFunc("/carritos/{id}", controllers.UpdateCarritos).Methods("PUT")
	r.HandleFunc("/carritos/{id}", controllers.DeleteCarritos).Methods("DELETE")
}
