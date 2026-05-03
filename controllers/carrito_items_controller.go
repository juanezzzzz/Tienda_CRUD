package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllCarritoItems(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_tr_carrito_items, id_carritos, id_productos, cantidad, precio_unitario, seleccionado, activo
		FROM "TR_carrito_items"`)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.CarritoItems
	for rows.Next() {
		var ci models.CarritoItems
		rows.Scan(&ci.ID_TR_Carrito_Items, &ci.ID_Carritos, &ci.ID_Productos, &ci.Cantidad, &ci.Precio_Unitario, &ci.Seleccionado, &ci.Activo)
		list = append(list, ci)
	}
	ResponseJSON(w, 200, list)
}

func GetCarritoItemsByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var ci models.CarritoItems
	err = config.DB.QueryRow(`
		SELECT id_tr_carrito_items, id_carritos, id_productos, cantidad, precio_unitario, seleccionado, activo
		FROM "TR_carrito_items" WHERE id_tr_carrito_items = $1`, id).
		Scan(&ci.ID_TR_Carrito_Items, &ci.ID_Carritos, &ci.ID_Productos, &ci.Cantidad, &ci.Precio_Unitario, &ci.Seleccionado, &ci.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, ci)
}

func CreateCarritoItems(w http.ResponseWriter, r *http.Request) {
	var ci models.CarritoItems
	json.NewDecoder(r.Body).Decode(&ci)
	err := config.DB.QueryRow(`
		INSERT INTO "TR_carrito_items" (id_carritos, id_productos, cantidad, precio_unitario, seleccionado, activo)
		VALUES ($1,$2,$3,$4,$5,$6) RETURNING id_tr_carrito_items`,
		ci.ID_Carritos, ci.ID_Productos, ci.Cantidad, ci.Precio_Unitario, ci.Seleccionado, ci.Activo).
		Scan(&ci.ID_TR_Carrito_Items)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, ci)
}

func UpdateCarritoItems(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var ci models.CarritoItems
	json.NewDecoder(r.Body).Decode(&ci)
	_, err := config.DB.Exec(`
		UPDATE "TR_carrito_items" SET id_carritos=$1, id_productos=$2, cantidad=$3,
		    precio_unitario=$4, seleccionado=$5, activo=$6
		WHERE id_tr_carrito_items=$7`,
		ci.ID_Carritos, ci.ID_Productos, ci.Cantidad, ci.Precio_Unitario, ci.Seleccionado, ci.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "CarritoItems updated successfully"})
}

func DeleteCarritoItems(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec(`DELETE FROM "TR_carrito_items" WHERE id_tr_carrito_items = $1`, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "CarritoItems deleted successfully"})
}
