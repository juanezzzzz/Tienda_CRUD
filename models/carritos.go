package models

import "database/sql"

type Carritos struct {
	ID_Carritos int           `json:"id_carritos"`
	ID_Usuarios int           `json:"id_usuarios"`
	ID_Estado   int           `json:"id_estado"`
	ID_Cupones  sql.NullInt64 `json:"id_cupones"`
	Activo      bool          `json:"activo"`
}
