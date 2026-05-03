package models

type CarritoItems struct {
	ID_TR_Carrito_Items int     `json:"id_tr_carrito_items"`
	ID_Carritos         int     `json:"id_carritos"`
	ID_Productos        int     `json:"id_productos"`
	Cantidad            int     `json:"cantidad"`
	Precio_Unitario     float64 `json:"precio_unitario"`
	Seleccionado        bool    `json:"seleccionado"`
	Activo              bool    `json:"activo"`
}
