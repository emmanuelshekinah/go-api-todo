package repository

import "github.com/emmanuelshekinah/go-api-todo/pkg/domain/core/reading"

type ETL interface {
	reading.ETL
}

type ETLPort struct {
	Read read.ETLService
}

func NewETLPort(etl ETL) *ETLPort {
	return *ETLPort{
		Read: reading.NewETLService(etl),
	}
}
