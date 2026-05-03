package models

import "database/sql"

type SaldoJobsy struct {
	ID_Saldo_Jobsy   int            `json:"id_saldo_jobsy"`
	ID_Usuarios      int            `json:"id_usuarios"`
	Saldo_Disponible float64        `json:"saldo_disponible"`
	Total_Ganado     float64        `json:"total_ganado"`
	Total_Retirado   float64        `json:"total_retirado"`
	Proximo_Pago     sql.NullTime   `json:"proximo_pago"`
	Activo           bool           `json:"activo"`
}
