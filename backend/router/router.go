package router

import (
	"fmt"
	"net/http"

	"github.com/ChowChunLeong/pineapple-language-api.git/controller"
	"github.com/ChowChunLeong/pineapple-language-api.git/pkg/setting"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	fmt.Println("testing")
	c.JSON(http.StatusOK, "hello")
}

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			allowedOrigins := setting.AppSetting.Origin

			for _, allowed := range allowedOrigins {
				if origin == allowed {
					return true
				}
			}
			return false
		},
		AllowMethods:     []string{"POST", "GET"}, // Allow multiple methods
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/", handler)
	auth := r.Group("/api/auth")
	{
		auth.POST("/oauth", controller.OAuth)
	}
	return r
}
