package models

import "database/sql"

type Transacciones struct {
	ID_Transacciones    int            `json:"id_transacciones"`
	ID_Usuarios         int            `json:"id_usuarios"`
	ID_Pagos            int            `json:"id_pagos"`
	ID_Tipo_Transaccion int            `json:"id_tipo_transaccion"`
	Monto               float64        `json:"monto"`
	Descripcion         sql.NullString `json:"descripcion"`
	Referencia          sql.NullString `json:"referencia"`
	Activo              bool           `json:"activo"`
}
