package main

import (
	"fmt"

	"github.com/emmanuelshekinah/go-api-todo/pkg/adapters/secondary/etl"
	"github.com/emmanuelshekinah/go-api-todo/pkg/domain/ports/repository"
)

func main() {
	etlAdapter := etl.NewEtlAdapter()
	etlPOrt := repository.NewEtlPort(etlAdapter)

	rt, _ : = etlPort.Read.ExtractTodos()

	fmt.Println(rt)
}
