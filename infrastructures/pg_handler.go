package infrastructures

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgHandler struct {
	Conn *gorm.DB
}

func NewPgHandler() *PgHandler {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return &PgHandler{Conn: db}
}

func (handler *PgHandler) FindByID(id int) (interface{}, error) {
	return nil, nil
}
