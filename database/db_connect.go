package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var sqlerr error

func ConnectDB() {
	db, sqlerr = sqlConnect()
		if sqlerr != nil {
			panic(sqlerr.Error())
		} else {
			fmt.Println("DB接続成功")
		}
}

func GetDB() *gorm.DB {
	return db
}


// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "Naka3Naka4"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "go_practice"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}