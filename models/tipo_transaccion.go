package models

type TipoTransaccion struct {
	ID_Tipo_Transaccion int    `json:"id_tipo_transaccion"`
	Nombre              string `json:"nombre"`
	Descripcion         string `json:"descripcion"`
	Activo              bool   `json:"activo"`
}
