package model

type APILog struct {
	ID          string
	UserName    string
	Content     string
	FromAddress string
	ToAddress   string
	Scenario    string
	CreatedAt   int64
}

type CreateAPILogArgs struct {
	ID          string
	UserName    string
	Content     string
	FromAddress string
	Scenario    string
	ToAddress   string
}

type GetAggregatedLogsArgs struct {
	UserName *string
}

type AggregatedLog struct {
	UserName string
	Scenario string
	Count    int
}
