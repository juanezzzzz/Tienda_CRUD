package models

import "database/sql"

type MetodosPagoGuardados struct {
	ID_Metodos_Pago_Guardados int            `json:"id_metodos_pago_guardados"`
	ID_Usuarios               int            `json:"id_usuarios"`
	ID_Metodo_Pago            int            `json:"id_metodo_pago"`
	Alias                     sql.NullString `json:"alias"`
	Numero_Cuenta             string         `json:"numero_cuenta"`
	ID_Tipo_Cuenta            int            `json:"id_tipo_cuenta"`
	Es_Principal              bool           `json:"es_principal"`
	Activo                    bool           `json:"activo"`
}
