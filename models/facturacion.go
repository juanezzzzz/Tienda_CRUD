package models

import "database/sql"

type Facturacion struct {
	ID_Pedidos        int            `json:"id_pedidos"`
	ID_Usuarios       int            `json:"id_usuarios"`
	ID_Carritos       int            `json:"id_carritos"`
	ID_Estado         int            `json:"id_estado"`
	Subtotal          float64        `json:"subtotal"`
	Descuento_Total   float64        `json:"descuento_total"`
	Cupon_Descuento   float64        `json:"cupon_descuento"`
	Envio             float64        `json:"envio"`
	Total             float64        `json:"total"`
	Factura_Nombre    sql.NullString `json:"factura_nombre"`
	Factura_Doc       sql.NullString `json:"factura_doc"`
	Factura_Correo    sql.NullString `json:"factura_correo"`
	Factura_Telefono  sql.NullString `json:"factura_telefono"`
	Factura_Direccion sql.NullString `json:"factura_direccion"`
	Factura_Ciudad    sql.NullString `json:"factura_ciudad"`
	Factura_Dpto      sql.NullString `json:"factura_dpto"`
	Activo            bool           `json:"activo"`
}
