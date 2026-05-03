package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllCategoriasTienda(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_categorias_tienda, nombre, icono, activa FROM categorias_tienda")
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.CategoriasTienda
	for rows.Next() {
		var c models.CategoriasTienda
		rows.Scan(&c.ID_Categorias_Tienda, &c.Nombre, &c.Icono, &c.Activa)
		list = append(list, c)
	}
	ResponseJSON(w, 200, list)
}

func GetCategoriasTiendaByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var c models.CategoriasTienda
	err = config.DB.QueryRow(`
		SELECT id_categorias_tienda, nombre, icono, activa
		FROM categorias_tienda WHERE id_categorias_tienda = $1`, id).
		Scan(&c.ID_Categorias_Tienda, &c.Nombre, &c.Icono, &c.Activa)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, c)
}

func CreateCategoriasTienda(w http.ResponseWriter, r *http.Request) {
	var c models.CategoriasTienda
	json.NewDecoder(r.Body).Decode(&c)
	err := config.DB.QueryRow(
		"INSERT INTO categorias_tienda (nombre, icono, activa) VALUES ($1, $2, $3) RETURNING id_categorias_tienda",
		c.Nombre, c.Icono, c.Activa).Scan(&c.ID_Categorias_Tienda)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, c)
}

func UpdateCategoriasTienda(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var c models.CategoriasTienda
	json.NewDecoder(r.Body).Decode(&c)
	_, err := config.DB.Exec(
		"UPDATE categorias_tienda SET nombre = $1, icono = $2, activa = $3 WHERE id_categorias_tienda = $4",
		c.Nombre, c.Icono, c.Activa, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "CategoriasTienda updated successfully"})
}

func DeleteCategoriasTienda(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM categorias_tienda WHERE id_categorias_tienda = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "CategoriasTienda deleted successfully"})
}
