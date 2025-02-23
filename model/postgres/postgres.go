package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Client *gorm.DB
}

func New(connectionString string) (*Postgres, error) {
	client, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Postgres{
		Client: client,
	}, nil
}

func (postgres Postgres) MigrateDB() {
	postgres.Client.AutoMigrate(&APILog{})
}
