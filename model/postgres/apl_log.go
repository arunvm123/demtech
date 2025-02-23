package postgres

import (
	"time"

	"github.com/arunvm123/demtech/model"
)

type APILog struct {
	ID          string `gorm:"primary_key"`
	UserName    string `gorm:"primary_key"`
	Content     string
	FromAddress string
	ToAddress   string
	Scenario    string
	CreatedAt   int64
}

func (postgres *Postgres) CreateAPILog(args model.CreateAPILogArgs) error {
	apiLog := APILog{
		ID:          args.ID,
		UserName:    args.UserName,
		Content:     args.Content,
		FromAddress: args.FromAddress,
		ToAddress:   args.ToAddress,
		Scenario:    args.Scenario,
		CreatedAt:   time.Now().Unix(),
	}

	err := postgres.Client.Create(&apiLog).Error
	if err != nil {
		return err
	}

	return nil
}
