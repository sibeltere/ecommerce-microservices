package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgreStorage struct{}

var DB *sql.DB

func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Eror opening database connection: %v", err)
	}

	//defer DB.Close() //defer kaynakları serbest bırakır

	//bağlantıyı test eder
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Erro connecting to database %v", err)
	}

	log.Println("Connected to PostgreSql database")

}
