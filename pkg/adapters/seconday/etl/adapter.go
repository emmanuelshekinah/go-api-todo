package etl

import "net/http"

type etladapter struct {
	client *http.client
}

func NewetlAdapter() *etladapter {
	return &etladapter{
		client: http.DefaultClient(),
	}
}
