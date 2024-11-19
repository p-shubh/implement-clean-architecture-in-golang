package database

import (
	"sync"

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

func PostgresDatabase(config *config.Config)  {
	once.Do(func ()  {
		
	})
}