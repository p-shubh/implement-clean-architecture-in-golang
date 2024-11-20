package database

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/config"
)

type postgresDatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func PostgresDatabase(config *config.Config) Database {
	once.Do(func() {
		// fmt.Println(config.Env.DB_HOST, config.Env.DB_USERNAME, config.Env.DB_PASSWORD, config.Env.DB_DATABASE, config.Env.DB_PORT, config.Env.SSL_MODE)

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			config.Env.DB_HOST, config.Env.DB_USERNAME, config.Env.DB_PASSWORD, config.Env.DB_DATABASE, config.Env.DB_PORT, config.Env.SSL_MODE,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logrus.Info(dsn)
			logrus.Fatalf("failed to connect postgres : %s", err)
		} else {
			logrus.Info("connected to postgres")
		}
		dbInstance = &postgresDatabase{Db: db}
	})
	return dbInstance
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return dbInstance.Db
}
