package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	Log_FILE_PATH = "D:\\workspace\\go\\src\\go-rbac\\logs"
	LOG_FILE_NAME = "info.log"
)

func GetDb() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.11.120:3306)/db_rbac")
	if err != nil {
		panic(err.Error())
		fmt.Print("mysql 链接错误")
	}
	return db
}
