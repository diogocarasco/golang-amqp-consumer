package handler_negocio1

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Handle(headers amqp.Table, body string) {
	fmt.Println(body)
}
