package routes

import (
	"tienda/controllers"
	"github.com/gorilla/mux"
)

func SetupPedidoItemsRoutes(r *mux.Router) {
	r.HandleFunc("/pedido-items", controllers.GetAllPedidoItems).Methods("GET")
	r.HandleFunc("/pedido-items/{id}", controllers.GetPedidoItemsByID).Methods("GET")
	r.HandleFunc("/pedido-items", controllers.CreatePedidoItems).Methods("POST")
	r.HandleFunc("/pedido-items/{id}", controllers.UpdatePedidoItems).Methods("PUT")
	r.HandleFunc("/pedido-items/{id}", controllers.DeletePedidoItems).Methods("DELETE")
}
