package models

import "database/sql"

type CategoriasTienda struct {
	ID_Categorias_Tienda int            `json:"id_categorias_tienda"`
	Nombre               string         `json:"nombre"`
	Icono                sql.NullString `json:"icono"`
	Activa               bool           `json:"activa"`
}
