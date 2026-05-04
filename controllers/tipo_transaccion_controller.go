package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllTipoTransaccion(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_tipo_transaccion, nombre, descripcion, activo FROM tipo_transaccion")
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.TipoTransaccion
	for rows.Next() {
		var t models.TipoTransaccion
		rows.Scan(&t.ID_Tipo_Transaccion, &t.Nombre, &t.Descripcion, &t.Activo)
		list = append(list, t)
	}
	ResponseJSON(w, 200, list)
}

func GetTipoTransaccionByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var t models.TipoTransaccion
	err = config.DB.QueryRow(`
		SELECT id_tipo_transaccion, nombre, descripcion, activo
		FROM tipo_transaccion WHERE id_tipo_transaccion = $1`, id).
		Scan(&t.ID_Tipo_Transaccion, &t.Nombre, &t.Descripcion, &t.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, t)
}

func CreateTipoTransaccion(w http.ResponseWriter, r *http.Request) {
	var t models.TipoTransaccion
	json.NewDecoder(r.Body).Decode(&t)
	err := config.DB.QueryRow(
		"INSERT INTO tipo_transaccion (nombre, descripcion, activo) VALUES ($1, $2, $3) RETURNING id_tipo_transaccion",
		t.Nombre, t.Descripcion, t.Activo).Scan(&t.ID_Tipo_Transaccion)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, t)
}

func UpdateTipoTransaccion(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var t models.TipoTransaccion
	json.NewDecoder(r.Body).Decode(&t)
	_, err := config.DB.Exec(
		"UPDATE tipo_transaccion SET nombre = $1, descripcion = $2, activo = $3 WHERE id_tipo_transaccion = $4",
		t.Nombre, t.Descripcion, t.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "TipoTransaccion updated successfully"})
}

func DeleteTipoTransaccion(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM tipo_transaccion WHERE id_tipo_transaccion = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "TipoTransaccion deleted successfully"})
}
