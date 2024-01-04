package db

import (
	"fmt"

	"github.com/keshu12345/notes/config"
	"github.com/keshu12345/notes/model"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Module = fx.Options(
	fx.Provide(NewPostgresDBInstance),
)

type PostgresDB struct {
	fx.Out
	DB *gorm.DB
}

func NewPostgresDBInstance(config *config.Configuration) (PostgresDB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		config.Postgres.Host, config.Postgres.UserName, config.Postgres.Password, config.Postgres.DatabaseName, config.Postgres.Port, "UTC")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return PostgresDB{}, err
	}
	db.AutoMigrate(&model.User{}, &model.Note{})
	return PostgresDB{DB: db}, nil
}

