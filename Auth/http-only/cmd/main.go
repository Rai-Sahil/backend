package main

import (
	"log"

	"github.com/Rai-Sahil/backend/internal/config"
	"github.com/Rai-Sahil/backend/internal/handler"
	"github.com/Rai-Sahil/backend/internal/repository"
	"github.com/Rai-Sahil/backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := config.ConnectDB()

	userRepo := repository.NewUserRepo(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	router := gin.Default()
	authHandler.RegisterRoutes(router)

	log.Println("Server started on :8000")
	log.Fatal(router.Run(":8000"))
}
