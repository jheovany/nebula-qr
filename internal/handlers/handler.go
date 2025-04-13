package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	
	_ "nebula-qr/docs"
	"nebula-qr/internal/services"
	"nebula-qr/internal/dto"
)

type QRHandler struct {
	service *services.QRService
}

// @title Nebula QR API
// @version 0.1
// @description API para generar y recuperar códigos QR.
// @termsOfService http://example.com/terms/

// @contact.name Soporte Nebula QR
// @contact.email jheovany.menjivarg@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host https://nebula-qr.onrender.com/
// @BasePath /
func RegisterQRHandlers(router *gin.Engine, service *services.QRService) {
	handler := &QRHandler{
		service: service,
	}

	router.POST("/qr", handler.CreateQR)
	router.GET("/qr/:id", handler.GetQR)

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// CreateQR genera un código QR a partir de un texto enviado.
// @Summary Genera un código QR
// @Description Genera un código QR con el texto proporcionado y la la duracion por lo menos a un minuto en el futuro.
// @Tags QR
// @Accept json
// @Produce json
// @Param request body dto.QRRequest true "Datos para generar el QR"
// @Success 201 {object} dto.QRResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /qr [post]
func (h *QRHandler) CreateQR(c *gin.Context) {
	var req dto.QRRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qr, err := h.service.GenerateQR(req.Text, req.ExpiresIn.ToDuration())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	exp := qr.ExpiresAt.Time()
	ld := qr.LastDownloaded.Time()
	response := dto.QRResponse {
		ID:             qr.ID.Hex(),
		Text:           qr.Text,
		Format:         qr.Format,
		Size:           qr.Size,
		CreatedAt:      qr.CreatedAt.Time(),
		ExpiresAt:      &exp,
		LastDownloaded: &ld,
		DownloadCount:  qr.DownloadCount,
	}

	c.JSON(http.StatusCreated, response)
}

// GetQR obtiene un código QR por su ID.
// @Summary Obtiene un código QR
// @Description Retorna la imagen del código QR correspondiente al ID proporcionado.
// @Tags QR
// @Produce png
// @Param id path string true "ID del código QR"
// @Success 200 {file} image/png
// @Failure 404 {object} map[string]string
// @Router /qr/{id} [get]
func (h *QRHandler) GetQR(c *gin.Context) {
	id := c.Param("id")
	qr, err := h.service.GetQR(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "QR not found"})
		return
	}

	c.Data(http.StatusOK, "image/png", qr.ImageData)
}
