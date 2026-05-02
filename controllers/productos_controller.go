package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllProductos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_productos, id_categorias_tienda, nombre, descripcion, id_tipo,
		       precio, precio_original, descuento_pct, stock, imagen_url,
		       es_nuevo, destacado, envio_gratis, rating_promedio, total_resenas, activo
		FROM productos`)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.Productos
	for rows.Next() {
		var p models.Productos
		rows.Scan(&p.ID_Productos, &p.ID_Categorias_Tienda, &p.Nombre, &p.Descripcion, &p.ID_Tipo,
			&p.Precio, &p.Precio_Original, &p.Descuento_Pct, &p.Stock, &p.Imagen_Url,
			&p.Es_Nuevo, &p.Destacado, &p.Envio_Gratis, &p.Rating_Promedio, &p.Total_Resenas, &p.Activo)
		list = append(list, p)
	}
	ResponseJSON(w, 200, list)
}

func GetProductosByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var p models.Productos
	err = config.DB.QueryRow(`
		SELECT id_productos, id_categorias_tienda, nombre, descripcion, id_tipo,
		       precio, precio_original, descuento_pct, stock, imagen_url,
		       es_nuevo, destacado, envio_gratis, rating_promedio, total_resenas, activo
		FROM productos WHERE id_productos = $1`, id).
		Scan(&p.ID_Productos, &p.ID_Categorias_Tienda, &p.Nombre, &p.Descripcion, &p.ID_Tipo,
			&p.Precio, &p.Precio_Original, &p.Descuento_Pct, &p.Stock, &p.Imagen_Url,
			&p.Es_Nuevo, &p.Destacado, &p.Envio_Gratis, &p.Rating_Promedio, &p.Total_Resenas, &p.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, p)
}

func CreateProductos(w http.ResponseWriter, r *http.Request) {
	var p models.Productos
	json.NewDecoder(r.Body).Decode(&p)
	err := config.DB.QueryRow(`
		INSERT INTO productos (id_categorias_tienda, nombre, descripcion, id_tipo,
		    precio, precio_original, descuento_pct, stock, imagen_url,
		    es_nuevo, destacado, envio_gratis, rating_promedio, total_resenas, activo)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)
		RETURNING id_productos`,
		p.ID_Categorias_Tienda, p.Nombre, p.Descripcion, p.ID_Tipo,
		p.Precio, p.Precio_Original, p.Descuento_Pct, p.Stock, p.Imagen_Url,
		p.Es_Nuevo, p.Destacado, p.Envio_Gratis, p.Rating_Promedio, p.Total_Resenas, p.Activo).
		Scan(&p.ID_Productos)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, p)
}

func UpdateProductos(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Productos
	json.NewDecoder(r.Body).Decode(&p)
	_, err := config.DB.Exec(`
		UPDATE productos SET id_categorias_tienda=$1, nombre=$2, descripcion=$3, id_tipo=$4,
		    precio=$5, precio_original=$6, descuento_pct=$7, stock=$8, imagen_url=$9,
		    es_nuevo=$10, destacado=$11, envio_gratis=$12, rating_promedio=$13, total_resenas=$14, activo=$15
		WHERE id_productos=$16`,
		p.ID_Categorias_Tienda, p.Nombre, p.Descripcion, p.ID_Tipo,
		p.Precio, p.Precio_Original, p.Descuento_Pct, p.Stock, p.Imagen_Url,
		p.Es_Nuevo, p.Destacado, p.Envio_Gratis, p.Rating_Promedio, p.Total_Resenas, p.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Productos updated successfully"})
}

func DeleteProductos(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM productos WHERE id_productos = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Productos deleted successfully"})
}
