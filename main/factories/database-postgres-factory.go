package factories

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db_postgres_con *gorm.DB

func NewDatabasePostgresOpenConnection() error {
	host := os.Getenv("DB_PG_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PG_PORT"))
	user := os.Getenv("DB_PG_USER")
	dbname := os.Getenv("DB_PG_NAME")
	pass := os.Getenv("DB_PG_PASSWORD")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo", host, user, pass, dbname, port)
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db_postgres_con = con
	newMigratePostgresModels()
	return err
}

func NewCloseDatabasePostgresConnection() error {
	db, err := db_postgres_con.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
