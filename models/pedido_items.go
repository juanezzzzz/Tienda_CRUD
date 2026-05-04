package models

type PedidoItems struct {
	ID_Pedido_Items int     `json:"id_pedido_items"`
	ID_Pedidos      int     `json:"id_pedidos"`
	ID_Productos    int     `json:"id_productos"`
	Cantidad        int     `json:"cantidad"`
	Precio_Unitario float64 `json:"precio_unitario"`
	Subtotal        float64 `json:"subtotal"`
	Activo          bool    `json:"activo"`
}
