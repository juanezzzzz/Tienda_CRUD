package main

import (
	"log"
	"net/http"

	"tienda/config"
	"tienda/routes"

	"github.com/gorilla/mux"
)

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	config.ConnectDB() // conectar a la base de datos

	r := mux.NewRouter() // Crear un nuevo router

	r.Use(EnableCORS) // Habilitar CORS

	routes.SetupMetodoPagoRoutes(r)
	routes.SetupTipoTransaccionRoutes(r)
	routes.SetupCategoriasTiendaRoutes(r)
	routes.SetupProductosRoutes(r)
	routes.SetupCuponesRoutes(r)
	routes.SetupCarritosRoutes(r)
	routes.SetupCarritoItemsRoutes(r)
	routes.SetupFacturacionRoutes(r)
	routes.SetupPedidoItemsRoutes(r)
	routes.SetupPagosRoutes(r)
	routes.SetupTransaccionesRoutes(r)
	routes.SetupSaldoJobsyRoutes(r)
	routes.SetupMetodosPagoGuardadosRoutes(r)

	log.Println("Tienda API started on port 8083")
	http.ListenAndServe(":8083", r)
}
