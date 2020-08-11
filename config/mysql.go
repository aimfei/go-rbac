package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
)

var db *gorm.DB

func init() {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "password", "localhost", 3306, "test")
	orb, err := gorm.Open("mysql", connArgs)
	if err != nil {
		logrus.Error("链接数据库异常")
	}
	orb.DB().SetMaxOpenConns(10)
	orb.DB().SetMaxIdleConns(200)
	orb.DB().SetConnMaxLifetime(time.Hour)
	db = orb
}
func GetDb1() *gorm.DB {
	return db
}

type DbConfig struct {
	Dbname   string
	Port     int
	Username string
	Password string
	Charset  string
	Host     string
}
