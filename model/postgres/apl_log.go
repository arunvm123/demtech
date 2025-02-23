package postgres

import (
	"fmt"
	"time"

	"github.com/arunvm123/demtech/model"
	"gorm.io/gorm"
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

func (postgres *Postgres) GetAggregatedLogs(args model.GetAggregatedLogsArgs) ([]model.AggregatedLog, error) {
	var query *gorm.DB

	if args.UserName != nil {
		query = postgres.Client.Table("api_logs").
			Select("user_name, scenario, count(*) as count").
			Where("user_name = ?", *args.UserName).
			Group("user_name, scenario")
	} else {
		query = postgres.Client.Table("api_logs").
			Select("'' as user_name, scenario, count(*) as count").
			Group("scenario")
	}

	var results []model.AggregatedLog
	if err := query.Find(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to get aggregated logs: %v", err)
	}

	return results, nil
}
