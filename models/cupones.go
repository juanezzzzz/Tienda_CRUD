package models

import (
	"database/sql"
	"time"
)

type Cupones struct {
	ID_Cupones     int             `json:"id_cupones"`
	Codigo         string          `json:"codigo"`
	Descuento_Pct  sql.NullInt64   `json:"descuento_pct"`
	Descuento_Fijo sql.NullFloat64 `json:"descuento_fijo"`
	Usos_Maximos   sql.NullInt64   `json:"usos_maximos"`
	Usos_Actuales  int             `json:"usos_actuales"`
	Activo         bool            `json:"activo"`
	Expira_En      sql.NullTime    `json:"expira_en"`
}

// Ensure time is imported
var _ = time.Now
