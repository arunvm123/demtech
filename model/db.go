package model

type DB interface {
	CreateAPILog(CreateAPILogArgs) error
	GetAggregatedLogs(GetAggregatedLogsArgs) ([]AggregatedLog, error)
}
