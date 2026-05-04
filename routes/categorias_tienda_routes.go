package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupCategoriasTiendaRoutes(r *mux.Router) {
	r.HandleFunc("/categorias-tienda", controllers.GetAllCategoriasTienda).Methods("GET")
	r.HandleFunc("/categorias-tienda/{id}", controllers.GetCategoriasTiendaByID).Methods("GET")
	r.HandleFunc("/categorias-tienda", controllers.CreateCategoriasTienda).Methods("POST")
	r.HandleFunc("/categorias-tienda/{id}", controllers.UpdateCategoriasTienda).Methods("PUT")
	r.HandleFunc("/categorias-tienda/{id}", controllers.DeleteCategoriasTienda).Methods("DELETE")
}
