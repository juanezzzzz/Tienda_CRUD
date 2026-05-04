package models

import "database/sql"

type Pagos struct {
	ID_Pagos           int            `json:"id_pagos"`
	ID_Pedidos         sql.NullInt64  `json:"id_pedidos"`
	ID_Contrataciones  sql.NullInt64  `json:"id_contrataciones"`
	ID_Usuarios        int            `json:"id_usuarios"`
	ID_Metodo_Pago     sql.NullInt64  `json:"id_metodo_pago"`
	ID_Tipo_Pago       int            `json:"id_tipo_pago"`
	ID_Estado          int            `json:"id_estado"`
	ID_Tipo_Cuenta     sql.NullInt64  `json:"id_tipo_cuenta"`
	Monto              float64        `json:"monto"`
	Referencia_Ext     sql.NullString `json:"referencia_ext"`
	Banco              sql.NullString `json:"banco"`
	Titular_Tarjeta    sql.NullString `json:"titular_tarjeta"`
	Ultimos4           sql.NullString `json:"ultimos4"`
	Activo             bool           `json:"activo"`
}
