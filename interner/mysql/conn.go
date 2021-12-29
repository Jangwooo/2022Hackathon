package mysql

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ErrConnection = fmt.Errorf("can not connect database")

func Connection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("mysql_user"),
		os.Getenv("mysql_pwd"),
		os.Getenv("mysql_host"),
		os.Getenv("mysql_port"),
		os.Getenv("mysql_db"))

	sqlDB, err := sql.Open("mysql", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:              sqlDB,
		DefaultStringSize: 256,
	}),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	return db, err
}
