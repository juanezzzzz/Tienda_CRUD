package models

import "database/sql"

type Productos struct {
	ID_Productos         int             `json:"id_productos"`
	ID_Categorias_Tienda int             `json:"id_categorias_tienda"`
	Nombre               string          `json:"nombre"`
	Descripcion          sql.NullString  `json:"descripcion"`
	ID_Tipo              int             `json:"id_tipo"`
	Precio               float64         `json:"precio"`
	Precio_Original      sql.NullFloat64 `json:"precio_original"`
	Descuento_Pct        int             `json:"descuento_pct"`
	Stock                int             `json:"stock"`
	Imagen_Url           sql.NullString  `json:"imagen_url"`
	Es_Nuevo             bool            `json:"es_nuevo"`
	Destacado            bool            `json:"destacado"`
	Envio_Gratis         bool            `json:"envio_gratis"`
	Rating_Promedio      float64         `json:"rating_promedio"`
	Total_Resenas        int             `json:"total_resenas"`
	Activo               bool            `json:"activo"`
}
