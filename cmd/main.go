package main

import (
	"api/internal/data"
	"api/internal/data/repositories"
	"api/internal/services"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	database, err := data.InitConnection(ctx)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	defer database.Close(ctx)

	userRepository := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepository)

}
