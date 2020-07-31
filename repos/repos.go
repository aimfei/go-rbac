package repos

import (
	"database/sql"
	"go-rbac/config"
	"log"
)

var db *sql.DB

func init() {
	log.Println("init database ")
	db = config.GetDb()
	db.Ping()
	db.SetMaxOpenConns(64)
	db.SetMaxIdleConns(5)
}
