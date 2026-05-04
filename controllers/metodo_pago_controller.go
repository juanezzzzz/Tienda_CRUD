package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllMetodoPago(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_metodo_pago, nombre, descripcion, activo FROM metodo_pago")
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.MetodoPago
	for rows.Next() {
		var m models.MetodoPago
		rows.Scan(&m.ID_Metodo_Pago, &m.Nombre, &m.Descripcion, &m.Activo)
		list = append(list, m)
	}
	ResponseJSON(w, 200, list)
}

func GetMetodoPagoByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var m models.MetodoPago
	err = config.DB.QueryRow(`
		SELECT id_metodo_pago, nombre, descripcion, activo
		FROM metodo_pago WHERE id_metodo_pago = $1`, id).
		Scan(&m.ID_Metodo_Pago, &m.Nombre, &m.Descripcion, &m.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, m)
}

func CreateMetodoPago(w http.ResponseWriter, r *http.Request) {
	var m models.MetodoPago
	json.NewDecoder(r.Body).Decode(&m)
	err := config.DB.QueryRow(
		"INSERT INTO metodo_pago (nombre, descripcion, activo) VALUES ($1, $2, $3) RETURNING id_metodo_pago",
		m.Nombre, m.Descripcion, m.Activo).Scan(&m.ID_Metodo_Pago)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, m)
}

func UpdateMetodoPago(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var m models.MetodoPago
	json.NewDecoder(r.Body).Decode(&m)
	_, err := config.DB.Exec(
		"UPDATE metodo_pago SET nombre = $1, descripcion = $2, activo = $3 WHERE id_metodo_pago = $4",
		m.Nombre, m.Descripcion, m.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "MetodoPago updated successfully"})
}

func DeleteMetodoPago(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM metodo_pago WHERE id_metodo_pago = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "MetodoPago deleted successfully"})
}
