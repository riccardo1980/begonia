package main

import (
	docs "begonia/docs"
	"begonia/internal/config"
	handlers "begonia/internal/handlers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg, err := config.LoadConfig()
	fmt.Printf("cfg: %v\n", cfg.IP)
	fmt.Printf("cfg: %v\n", cfg.Port)

	if err != nil {
		fmt.Print("error")
		os.Exit(-1)
	}
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	r.SetTrustedProxies([]string{cfg.IP})

	handlers.RegisterHealthHandlers(v1)
	handlers.RegisterHelloworldHandlers(v1)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":" + cfg.Port)

}
