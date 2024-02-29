package main

import (
	"context"
	"fmt"
	"github.com/nuoxoxo/gg/tree/main/chi-crud-microservice-redis/app"
)

func main() {
	app := app.New()
	err := app.Start(context.TODO()) // means that we will impl. it later on
	if err != nil {
		fmt.Errorf("/err: %w", err)
	}
}
