package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllPagos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_pagos, id_pedidos, id_contrataciones, id_usuarios,
		       id_metodo_pago, id_tipo_pago, id_estado, id_tipo_cuenta,
		       monto, referencia_ext, banco, titular_tarjeta, ultimos4, activo
		FROM pagos`)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.Pagos
	for rows.Next() {
		var p models.Pagos
		rows.Scan(&p.ID_Pagos, &p.ID_Pedidos, &p.ID_Contrataciones, &p.ID_Usuarios,
			&p.ID_Metodo_Pago, &p.ID_Tipo_Pago, &p.ID_Estado, &p.ID_Tipo_Cuenta,
			&p.Monto, &p.Referencia_Ext, &p.Banco, &p.Titular_Tarjeta, &p.Ultimos4, &p.Activo)
		list = append(list, p)
	}
	ResponseJSON(w, 200, list)
}

func GetPagosByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var p models.Pagos
	err = config.DB.QueryRow(`
		SELECT id_pagos, id_pedidos, id_contrataciones, id_usuarios,
		       id_metodo_pago, id_tipo_pago, id_estado, id_tipo_cuenta,
		       monto, referencia_ext, banco, titular_tarjeta, ultimos4, activo
		FROM pagos WHERE id_pagos = $1`, id).
		Scan(&p.ID_Pagos, &p.ID_Pedidos, &p.ID_Contrataciones, &p.ID_Usuarios,
			&p.ID_Metodo_Pago, &p.ID_Tipo_Pago, &p.ID_Estado, &p.ID_Tipo_Cuenta,
			&p.Monto, &p.Referencia_Ext, &p.Banco, &p.Titular_Tarjeta, &p.Ultimos4, &p.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, p)
}

func CreatePagos(w http.ResponseWriter, r *http.Request) {
	var p models.Pagos
	json.NewDecoder(r.Body).Decode(&p)
	err := config.DB.QueryRow(`
		INSERT INTO pagos (id_pedidos, id_contrataciones, id_usuarios,
		    id_metodo_pago, id_tipo_pago, id_estado, id_tipo_cuenta,
		    monto, referencia_ext, banco, titular_tarjeta, ultimos4, activo)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING id_pagos`,
		p.ID_Pedidos, p.ID_Contrataciones, p.ID_Usuarios,
		p.ID_Metodo_Pago, p.ID_Tipo_Pago, p.ID_Estado, p.ID_Tipo_Cuenta,
		p.Monto, p.Referencia_Ext, p.Banco, p.Titular_Tarjeta, p.Ultimos4, p.Activo).
		Scan(&p.ID_Pagos)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, p)
}

func UpdatePagos(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Pagos
	json.NewDecoder(r.Body).Decode(&p)
	_, err := config.DB.Exec(`
		UPDATE pagos SET id_pedidos=$1, id_contrataciones=$2, id_usuarios=$3,
		    id_metodo_pago=$4, id_tipo_pago=$5, id_estado=$6, id_tipo_cuenta=$7,
		    monto=$8, referencia_ext=$9, banco=$10, titular_tarjeta=$11, ultimos4=$12, activo=$13
		WHERE id_pagos=$14`,
		p.ID_Pedidos, p.ID_Contrataciones, p.ID_Usuarios,
		p.ID_Metodo_Pago, p.ID_Tipo_Pago, p.ID_Estado, p.ID_Tipo_Cuenta,
		p.Monto, p.Referencia_Ext, p.Banco, p.Titular_Tarjeta, p.Ultimos4, p.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Pagos updated successfully"})
}

func DeletePagos(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM pagos WHERE id_pagos = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Pagos deleted successfully"})
}
