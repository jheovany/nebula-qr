package main

import (
	"log"
	"nebula-qr/internal/handlers"
	"nebula-qr/internal/services"
	"nebula-qr/pkg/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI no está configurada en el archivo .env")
	}

	db := database.ConnectMongoDB(mongoURI)

	// Crear el servicio con la base de datos inyectada
	qrService := services.NewQRService(db)

	// Inicializar router Gin
	router := gin.Default()

	// Registrar rutas
	// Registrar rutas pasando el servicio
	handlers.RegisterQRHandlers(router, qrService)

	// Iniciar servidor
	router.Run(":8080")
}
