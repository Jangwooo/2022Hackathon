package mysql

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("mysql_user"),
		os.Getenv("mysql_pwd"),
		os.Getenv("mysql_host"),
		os.Getenv("mysql_port"),
		os.Getenv("mysql_db"))

	sqlDB, err := sql.Open("mysql", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:              sqlDB,
		DefaultStringSize: 191,
	}))

	if err != nil {
		panic(err.Error())
	}
	return db
}
