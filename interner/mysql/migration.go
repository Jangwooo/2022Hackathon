package mysql

import (
	"fmt"

	"github.com/Jangwooo/2022Hackathon/interner/mysql/model"
)

var ErrMigration = fmt.Errorf("can not migrate database")

func Migration() error {
	db := Connection()

	if err := db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Post{},
	); err != nil {
		return ErrMigration
	}

	return nil
}
