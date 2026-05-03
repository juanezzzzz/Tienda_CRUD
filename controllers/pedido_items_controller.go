package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllPedidoItems(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_pedido_items, id_pedidos, id_productos, cantidad, precio_unitario, subtotal, activo
		FROM pedido_items`)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.PedidoItems
	for rows.Next() {
		var pi models.PedidoItems
		rows.Scan(&pi.ID_Pedido_Items, &pi.ID_Pedidos, &pi.ID_Productos, &pi.Cantidad, &pi.Precio_Unitario, &pi.Subtotal, &pi.Activo)
		list = append(list, pi)
	}
	ResponseJSON(w, 200, list)
}

func GetPedidoItemsByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var pi models.PedidoItems
	err = config.DB.QueryRow(`
		SELECT id_pedido_items, id_pedidos, id_productos, cantidad, precio_unitario, subtotal, activo
		FROM pedido_items WHERE id_pedido_items = $1`, id).
		Scan(&pi.ID_Pedido_Items, &pi.ID_Pedidos, &pi.ID_Productos, &pi.Cantidad, &pi.Precio_Unitario, &pi.Subtotal, &pi.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, pi)
}

func CreatePedidoItems(w http.ResponseWriter, r *http.Request) {
	var pi models.PedidoItems
	json.NewDecoder(r.Body).Decode(&pi)
	err := config.DB.QueryRow(`
		INSERT INTO pedido_items (id_pedidos, id_productos, cantidad, precio_unitario, subtotal, activo)
		VALUES ($1,$2,$3,$4,$5,$6) RETURNING id_pedido_items`,
		pi.ID_Pedidos, pi.ID_Productos, pi.Cantidad, pi.Precio_Unitario, pi.Subtotal, pi.Activo).
		Scan(&pi.ID_Pedido_Items)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, pi)
}

func UpdatePedidoItems(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var pi models.PedidoItems
	json.NewDecoder(r.Body).Decode(&pi)
	_, err := config.DB.Exec(`
		UPDATE pedido_items SET id_pedidos=$1, id_productos=$2, cantidad=$3,
		    precio_unitario=$4, subtotal=$5, activo=$6
		WHERE id_pedido_items=$7`,
		pi.ID_Pedidos, pi.ID_Productos, pi.Cantidad, pi.Precio_Unitario, pi.Subtotal, pi.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "PedidoItems updated successfully"})
}

func DeletePedidoItems(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM pedido_items WHERE id_pedido_items = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "PedidoItems deleted successfully"})
}
