package infrastructures

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/khuongnguyenBlue/slacky/transport"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type PgHandler struct {
	Conn *gorm.DB
}

func NewPgHandler() *PgHandler {
	// build dsn string by environment variables
	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetUint("DB_PORT")
	dbName := viper.GetString("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		dbHost, username, password, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "slacky.",
		},
	})
	if err != nil {
		panic(err.Error())
	}
	return &PgHandler{Conn: db}
}

func (handler *PgHandler) HandleError(res *gorm.DB) error {
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		err := res.Error.(*pgconn.PgError)
		return transport.NewErrorBody(err.Message, err.Code)
	}

	return nil
}

func (handler *PgHandler) Create(model interface{}) error {
	result := handler.Conn.Create(model)
	return handler.HandleError(result)
}

func (handler *PgHandler) FindByField(model interface{}, field string, value string) (error) {
	result := handler.Conn.Where(fmt.Sprintf("%s = ?", field), value).First(model)
	return handler.HandleError(result)
}
