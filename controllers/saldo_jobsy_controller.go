package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"tienda/config"
	"tienda/models"

	"github.com/gorilla/mux"
)

func GetAllSaldoJobsy(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_saldo_jobsy, id_usuarios, saldo_disponible, total_ganado, total_retirado, proximo_pago, activo
		FROM saldo_jobsy`)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.SaldoJobsy
	for rows.Next() {
		var s models.SaldoJobsy
		rows.Scan(&s.ID_Saldo_Jobsy, &s.ID_Usuarios, &s.Saldo_Disponible, &s.Total_Ganado, &s.Total_Retirado, &s.Proximo_Pago, &s.Activo)
		list = append(list, s)
	}
	ResponseJSON(w, 200, list)
}

func GetSaldoJobsyByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var s models.SaldoJobsy
	err = config.DB.QueryRow(`
		SELECT id_saldo_jobsy, id_usuarios, saldo_disponible, total_ganado, total_retirado, proximo_pago, activo
		FROM saldo_jobsy WHERE id_saldo_jobsy = $1`, id).
		Scan(&s.ID_Saldo_Jobsy, &s.ID_Usuarios, &s.Saldo_Disponible, &s.Total_Ganado, &s.Total_Retirado, &s.Proximo_Pago, &s.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, s)
}

func CreateSaldoJobsy(w http.ResponseWriter, r *http.Request) {
	var s models.SaldoJobsy
	json.NewDecoder(r.Body).Decode(&s)
	err := config.DB.QueryRow(`
		INSERT INTO saldo_jobsy (id_usuarios, saldo_disponible, total_ganado, total_retirado, proximo_pago, activo)
		VALUES ($1,$2,$3,$4,$5,$6) RETURNING id_saldo_jobsy`,
		s.ID_Usuarios, s.Saldo_Disponible, s.Total_Ganado, s.Total_Retirado, s.Proximo_Pago, s.Activo).
		Scan(&s.ID_Saldo_Jobsy)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, s)
}

func UpdateSaldoJobsy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var s models.SaldoJobsy
	json.NewDecoder(r.Body).Decode(&s)
	_, err := config.DB.Exec(`
		UPDATE saldo_jobsy SET id_usuarios=$1, saldo_disponible=$2, total_ganado=$3,
		    total_retirado=$4, proximo_pago=$5, activo=$6
		WHERE id_saldo_jobsy=$7`,
		s.ID_Usuarios, s.Saldo_Disponible, s.Total_Ganado, s.Total_Retirado, s.Proximo_Pago, s.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "SaldoJobsy updated successfully"})
}

func DeleteSaldoJobsy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM saldo_jobsy WHERE id_saldo_jobsy = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "SaldoJobsy deleted successfully"})
}
