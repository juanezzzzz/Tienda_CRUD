package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllMetodosPagoGuardados(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_metodos_pago_guardados, id_usuarios, id_metodo_pago, alias,
		       numero_cuenta, id_tipo_cuenta, es_principal, activo
		FROM metodos_pago_guardados`)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.MetodosPagoGuardados
	for rows.Next() {
		var m models.MetodosPagoGuardados
		rows.Scan(&m.ID_Metodos_Pago_Guardados, &m.ID_Usuarios, &m.ID_Metodo_Pago, &m.Alias,
			&m.Numero_Cuenta, &m.ID_Tipo_Cuenta, &m.Es_Principal, &m.Activo)
		list = append(list, m)
	}
	ResponseJSON(w, 200, list)
}

func GetMetodosPagoGuardadosByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var m models.MetodosPagoGuardados
	err = config.DB.QueryRow(`
		SELECT id_metodos_pago_guardados, id_usuarios, id_metodo_pago, alias,
		       numero_cuenta, id_tipo_cuenta, es_principal, activo
		FROM metodos_pago_guardados WHERE id_metodos_pago_guardados = $1`, id).
		Scan(&m.ID_Metodos_Pago_Guardados, &m.ID_Usuarios, &m.ID_Metodo_Pago, &m.Alias,
			&m.Numero_Cuenta, &m.ID_Tipo_Cuenta, &m.Es_Principal, &m.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, m)
}

func CreateMetodosPagoGuardados(w http.ResponseWriter, r *http.Request) {
	var m models.MetodosPagoGuardados
	json.NewDecoder(r.Body).Decode(&m)
	err := config.DB.QueryRow(`
		INSERT INTO metodos_pago_guardados (id_usuarios, id_metodo_pago, alias, numero_cuenta, id_tipo_cuenta, es_principal, activo)
		VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id_metodos_pago_guardados`,
		m.ID_Usuarios, m.ID_Metodo_Pago, m.Alias, m.Numero_Cuenta, m.ID_Tipo_Cuenta, m.Es_Principal, m.Activo).
		Scan(&m.ID_Metodos_Pago_Guardados)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, m)
}

func UpdateMetodosPagoGuardados(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var m models.MetodosPagoGuardados
	json.NewDecoder(r.Body).Decode(&m)
	_, err := config.DB.Exec(`
		UPDATE metodos_pago_guardados SET id_usuarios=$1, id_metodo_pago=$2, alias=$3,
		    numero_cuenta=$4, id_tipo_cuenta=$5, es_principal=$6, activo=$7
		WHERE id_metodos_pago_guardados=$8`,
		m.ID_Usuarios, m.ID_Metodo_Pago, m.Alias, m.Numero_Cuenta, m.ID_Tipo_Cuenta, m.Es_Principal, m.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "MetodosPagoGuardados updated successfully"})
}

func DeleteMetodosPagoGuardados(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM metodos_pago_guardados WHERE id_metodos_pago_guardados = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "MetodosPagoGuardados deleted successfully"})
}
