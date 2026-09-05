package main

import (
	"api/internal/data"
	"api/internal/data/repositories"
	"api/internal/dto"
	"api/internal/services"
	"context"
	"fmt"
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

	//mocked for testing

	resp := userService.Register(
		ctx,
		dto.Register{
			Username: "admin",
			Password: "admin123",
		},
	)

	fmt.Println(resp)
}
