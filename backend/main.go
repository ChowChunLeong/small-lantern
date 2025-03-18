package main

import (
	"fmt"
	"os"

	"github.com/ChowChunLeong/pineapple-language-api.git/database"
	"github.com/ChowChunLeong/pineapple-language-api.git/pkg/setting"
	"github.com/ChowChunLeong/pineapple-language-api.git/router"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Initializing application...")
}

func main() {
	setting.Setup()

	gin.SetMode(os.Getenv("ENV"))

	database.SetupDatabaseConnection()

	/* Custom HTTP configuration */
	routersInit := router.SetupRouter()
	routersInit.Run(setting.AppSetting.Port)
}
