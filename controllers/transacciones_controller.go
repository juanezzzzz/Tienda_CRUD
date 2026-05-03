package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllTransacciones(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_transacciones, id_usuarios, id_pagos, id_tipo_transaccion,
		       monto, descripcion, referencia, activo
		FROM transacciones`)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.Transacciones
	for rows.Next() {
		var t models.Transacciones
		rows.Scan(&t.ID_Transacciones, &t.ID_Usuarios, &t.ID_Pagos, &t.ID_Tipo_Transaccion,
			&t.Monto, &t.Descripcion, &t.Referencia, &t.Activo)
		list = append(list, t)
	}
	ResponseJSON(w, 200, list)
}

func GetTransaccionesByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var t models.Transacciones
	err = config.DB.QueryRow(`
		SELECT id_transacciones, id_usuarios, id_pagos, id_tipo_transaccion,
		       monto, descripcion, referencia, activo
		FROM transacciones WHERE id_transacciones = $1`, id).
		Scan(&t.ID_Transacciones, &t.ID_Usuarios, &t.ID_Pagos, &t.ID_Tipo_Transaccion,
			&t.Monto, &t.Descripcion, &t.Referencia, &t.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, t)
}

func CreateTransacciones(w http.ResponseWriter, r *http.Request) {
	var t models.Transacciones
	json.NewDecoder(r.Body).Decode(&t)
	err := config.DB.QueryRow(`
		INSERT INTO transacciones (id_usuarios, id_pagos, id_tipo_transaccion, monto, descripcion, referencia, activo)
		VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id_transacciones`,
		t.ID_Usuarios, t.ID_Pagos, t.ID_Tipo_Transaccion, t.Monto, t.Descripcion, t.Referencia, t.Activo).
		Scan(&t.ID_Transacciones)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, t)
}

func UpdateTransacciones(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var t models.Transacciones
	json.NewDecoder(r.Body).Decode(&t)
	_, err := config.DB.Exec(`
		UPDATE transacciones SET id_usuarios=$1, id_pagos=$2, id_tipo_transaccion=$3,
		    monto=$4, descripcion=$5, referencia=$6, activo=$7
		WHERE id_transacciones=$8`,
		t.ID_Usuarios, t.ID_Pagos, t.ID_Tipo_Transaccion, t.Monto, t.Descripcion, t.Referencia, t.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Transacciones updated successfully"})
}

func DeleteTransacciones(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM transacciones WHERE id_transacciones = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Transacciones deleted successfully"})
}
