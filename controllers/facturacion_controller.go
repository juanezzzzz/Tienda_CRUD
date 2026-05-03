package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllFacturacion(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_pedidos, id_usuarios, id_carritos, id_estado,
		       subtotal, descuento_total, cupon_descuento, envio, total,
		       factura_nombre, factura_doc, factura_correo, factura_telefono,
		       factura_direccion, factura_ciudad, factura_dpto, activo
		FROM facturacion`)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.Facturacion
	for rows.Next() {
		var f models.Facturacion
		rows.Scan(&f.ID_Pedidos, &f.ID_Usuarios, &f.ID_Carritos, &f.ID_Estado,
			&f.Subtotal, &f.Descuento_Total, &f.Cupon_Descuento, &f.Envio, &f.Total,
			&f.Factura_Nombre, &f.Factura_Doc, &f.Factura_Correo, &f.Factura_Telefono,
			&f.Factura_Direccion, &f.Factura_Ciudad, &f.Factura_Dpto, &f.Activo)
		list = append(list, f)
	}
	ResponseJSON(w, 200, list)
}

func GetFacturacionByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var f models.Facturacion
	err = config.DB.QueryRow(`
		SELECT id_pedidos, id_usuarios, id_carritos, id_estado,
		       subtotal, descuento_total, cupon_descuento, envio, total,
		       factura_nombre, factura_doc, factura_correo, factura_telefono,
		       factura_direccion, factura_ciudad, factura_dpto, activo
		FROM facturacion WHERE id_pedidos = $1`, id).
		Scan(&f.ID_Pedidos, &f.ID_Usuarios, &f.ID_Carritos, &f.ID_Estado,
			&f.Subtotal, &f.Descuento_Total, &f.Cupon_Descuento, &f.Envio, &f.Total,
			&f.Factura_Nombre, &f.Factura_Doc, &f.Factura_Correo, &f.Factura_Telefono,
			&f.Factura_Direccion, &f.Factura_Ciudad, &f.Factura_Dpto, &f.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, f)
}

func CreateFacturacion(w http.ResponseWriter, r *http.Request) {
	var f models.Facturacion
	json.NewDecoder(r.Body).Decode(&f)
	err := config.DB.QueryRow(`
		INSERT INTO facturacion (id_usuarios, id_carritos, id_estado,
		    subtotal, descuento_total, cupon_descuento, envio, total,
		    factura_nombre, factura_doc, factura_correo, factura_telefono,
		    factura_direccion, factura_ciudad, factura_dpto, activo)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)
		RETURNING id_pedidos`,
		f.ID_Usuarios, f.ID_Carritos, f.ID_Estado,
		f.Subtotal, f.Descuento_Total, f.Cupon_Descuento, f.Envio, f.Total,
		f.Factura_Nombre, f.Factura_Doc, f.Factura_Correo, f.Factura_Telefono,
		f.Factura_Direccion, f.Factura_Ciudad, f.Factura_Dpto, f.Activo).
		Scan(&f.ID_Pedidos)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, f)
}

func UpdateFacturacion(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var f models.Facturacion
	json.NewDecoder(r.Body).Decode(&f)
	_, err := config.DB.Exec(`
		UPDATE facturacion SET id_usuarios=$1, id_carritos=$2, id_estado=$3,
		    subtotal=$4, descuento_total=$5, cupon_descuento=$6, envio=$7, total=$8,
		    factura_nombre=$9, factura_doc=$10, factura_correo=$11, factura_telefono=$12,
		    factura_direccion=$13, factura_ciudad=$14, factura_dpto=$15, activo=$16
		WHERE id_pedidos=$17`,
		f.ID_Usuarios, f.ID_Carritos, f.ID_Estado,
		f.Subtotal, f.Descuento_Total, f.Cupon_Descuento, f.Envio, f.Total,
		f.Factura_Nombre, f.Factura_Doc, f.Factura_Correo, f.Factura_Telefono,
		f.Factura_Direccion, f.Factura_Ciudad, f.Factura_Dpto, f.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Facturacion updated successfully"})
}

func DeleteFacturacion(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM facturacion WHERE id_pedidos = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Facturacion deleted successfully"})
}
