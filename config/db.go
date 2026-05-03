package config

import (
	"database/sql"
	"fmt"
	"log"

	_"github.com/lib/pq"
)

var DB *sql.DB

// connectDB

func ConnectDB() {
	host := "localhost"
	port := 5432
	user := "postgres"
	dbname := "Jobsy"
	password := "3105437922"
	schema := "tienda"

	psqlinfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, user, password, dbname, schema,
	)

	var err error
	DB, err = sql.Open("postgres", psqlinfo)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}

	fmt.Println("Conectado a la base de datos")
	fmt.Println("Schema: ", schema)
}
