package main

import (
	"api_reservas/controller"
	"api_reservas/models"
	"api_reservas/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	godotenv.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("error conectando la base da datos")
	}

	controller.DB = db

	models.MigrarTablaUsuario(db)
	models.MigrarTablaReserva(db)

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":7000")

}
