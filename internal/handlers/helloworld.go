package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /helloworld [get]
func HelloworldHandler(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func RegisterHelloworldHandlers(router *gin.RouterGroup) {
	router.GET("/helloworld", HelloworldHandler)
}
