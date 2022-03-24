package reading

import "github.com/emmanuelshekinah/go-api-todo/pkg/domain/models"

type ETL interface {
	ExtractTodos() (rawTodos []byte, err error)
}

type ETLService interface {
	ExtractTodos() (todos []models.RawTodo, err error)
}

type etlservice struct {
	etl ETL
}

func NewETLService(etl ETL) ETLService {
	return &etlservice{etl}
}

//a.etl.Read.ExtractTools
