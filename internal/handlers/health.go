package handlers

import (
	models "begonia/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary Health check
// @Schemes
// @Description Basic healthcheck API
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} models.HealthCheckResponse
// @Router /health [get]
func HealthCheckHandler(c *gin.Context) {
	currentTime := time.Now().UTC()
	c.JSON(http.StatusOK, models.HealthCheckResponse{
		Status:    "healthy",
		Timestamp: currentTime.Format(time.RFC3339),
	})
}

func RegisterHealthHandlers(router *gin.RouterGroup) {
	router.GET("/health", HealthCheckHandler)
}
