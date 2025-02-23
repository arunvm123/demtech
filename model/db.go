package model

type DB interface {
	CreateAPILog(CreateAPILogArgs) error
}
