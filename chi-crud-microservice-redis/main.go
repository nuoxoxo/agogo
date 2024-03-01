package main

import (
	"context"
	"fmt"

	"github.com/nuoxoxo/agogo/tree/main/chi-crud-microservice-redis/application"
)

func main() {
	app := application.New()
	err := app.Start(context.TODO()) // means that we will impl. it later on
	if err != nil {
		fmt.Errorf("/err: %w", err)
	}
}
