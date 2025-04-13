package main

import (
	"log"
	"nebula-qr/docs"
	"nebula-qr/internal/handlers"
	"nebula-qr/internal/services"
	"nebula-qr/pkg/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error cargando archivo .env")
		}
	}

	// Configurar Swagger dinámicamente
    host := os.Getenv("SWAGGER_HOST")
	schemas := []string{"https"}
    if host == "" {
        host = "localhost:8080" // Valor por defecto para desarrollo
		schemas = []string{"http"} // Usar HTTP en desarrollo
    }
    docs.SwaggerInfo.Host = host
    docs.SwaggerInfo.Schemes = schemas // Usar HTTPS en producción

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
