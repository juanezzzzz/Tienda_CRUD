package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllCupones(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_cupones, codigo, descuento_pct, descuento_fijo,
		       usos_maximos, usos_actuales, activo, expira_en
		FROM cupones`)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.Cupones
	for rows.Next() {
		var c models.Cupones
		rows.Scan(&c.ID_Cupones, &c.Codigo, &c.Descuento_Pct, &c.Descuento_Fijo,
			&c.Usos_Maximos, &c.Usos_Actuales, &c.Activo, &c.Expira_En)
		list = append(list, c)
	}
	ResponseJSON(w, 200, list)
}

func GetCuponesByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var c models.Cupones
	err = config.DB.QueryRow(`
		SELECT id_cupones, codigo, descuento_pct, descuento_fijo,
		       usos_maximos, usos_actuales, activo, expira_en
		FROM cupones WHERE id_cupones = $1`, id).
		Scan(&c.ID_Cupones, &c.Codigo, &c.Descuento_Pct, &c.Descuento_Fijo,
			&c.Usos_Maximos, &c.Usos_Actuales, &c.Activo, &c.Expira_En)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, c)
}

func CreateCupones(w http.ResponseWriter, r *http.Request) {
	var c models.Cupones
	json.NewDecoder(r.Body).Decode(&c)
	err := config.DB.QueryRow(`
		INSERT INTO cupones (codigo, descuento_pct, descuento_fijo, usos_maximos, usos_actuales, activo, expira_en)
		VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id_cupones`,
		c.Codigo, c.Descuento_Pct, c.Descuento_Fijo, c.Usos_Maximos, c.Usos_Actuales, c.Activo, c.Expira_En).
		Scan(&c.ID_Cupones)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, c)
}

func UpdateCupones(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var c models.Cupones
	json.NewDecoder(r.Body).Decode(&c)
	_, err := config.DB.Exec(`
		UPDATE cupones SET codigo=$1, descuento_pct=$2, descuento_fijo=$3,
		    usos_maximos=$4, usos_actuales=$5, activo=$6, expira_en=$7
		WHERE id_cupones=$8`,
		c.Codigo, c.Descuento_Pct, c.Descuento_Fijo,
		c.Usos_Maximos, c.Usos_Actuales, c.Activo, c.Expira_En, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Cupones updated successfully"})
}

func DeleteCupones(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM cupones WHERE id_cupones = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Cupones deleted successfully"})
}
