package models

type MetodoPago struct {
	ID_Metodo_Pago int    `json:"id_metodo_pago"`
	Nombre         string `json:"nombre"`
	Descripcion    string `json:"descripcion"`
	Activo         bool   `json:"activo"`
}
