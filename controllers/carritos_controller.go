package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllCarritos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_carritos, id_usuarios, id_estado, id_cupones, activo FROM carritos")
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.Carritos
	for rows.Next() {
		var c models.Carritos
		rows.Scan(&c.ID_Carritos, &c.ID_Usuarios, &c.ID_Estado, &c.ID_Cupones, &c.Activo)
		list = append(list, c)
	}
	ResponseJSON(w, 200, list)
}

func GetCarritosByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var c models.Carritos
	err = config.DB.QueryRow(`
		SELECT id_carritos, id_usuarios, id_estado, id_cupones, activo
		FROM carritos WHERE id_carritos = $1`, id).
		Scan(&c.ID_Carritos, &c.ID_Usuarios, &c.ID_Estado, &c.ID_Cupones, &c.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, c)
}

func CreateCarritos(w http.ResponseWriter, r *http.Request) {
	var c models.Carritos
	json.NewDecoder(r.Body).Decode(&c)
	err := config.DB.QueryRow(
		"INSERT INTO carritos (id_usuarios, id_estado, id_cupones, activo) VALUES ($1,$2,$3,$4) RETURNING id_carritos",
		c.ID_Usuarios, c.ID_Estado, c.ID_Cupones, c.Activo).Scan(&c.ID_Carritos)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, c)
}

func UpdateCarritos(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var c models.Carritos
	json.NewDecoder(r.Body).Decode(&c)
	_, err := config.DB.Exec(
		"UPDATE carritos SET id_usuarios=$1, id_estado=$2, id_cupones=$3, activo=$4 WHERE id_carritos=$5",
		c.ID_Usuarios, c.ID_Estado, c.ID_Cupones, c.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Carritos updated successfully"})
}

func DeleteCarritos(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM carritos WHERE id_carritos = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Carritos deleted successfully"})
}
