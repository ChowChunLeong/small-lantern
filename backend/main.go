package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ChowChunLeong/pineapple-language-api.git/database"
	"github.com/ChowChunLeong/pineapple-language-api.git/pkg/setting"
	"github.com/ChowChunLeong/pineapple-language-api.git/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Function to read secrets from Docker Secrets if they exist
func getSecret(secretName string) string {
	secretPath := "/run/secrets/" + secretName
	if content, err := ioutil.ReadFile(secretPath); err == nil {
		return string(content)
	}
	return os.Getenv(secretName) // Fallback to env variable if secret not found
}

func init() {
	// Load .env only in the local(debug) environment, do not load in production(release)
	if os.Getenv("ENV") != "release" {
		err := godotenv.Load("conf/.env")
		if err != nil {
			log.Println("No .env file found, using system environment variables.")
		}
		// Override environment variables with secrets in production
		// if os.Getenv("ENV") == "release" {
		// 	os.Setenv("PLE_DB_USER", getSecret("ple_db_user"))
		// 	os.Setenv("PLE_DB_PASS", getSecret("ple_db_pass"))
		// 	os.Setenv("PLE_DB_HOST", getSecret("ple_db_host"))
		// 	os.Setenv("PLE_DB_NAME", getSecret("ple_db_name"))
		// }
	}
}

func main() {
	setting.Setup()

	gin.SetMode(os.Getenv("ENV"))

	database.SetupDatabaseConnection()

	/* Custom HTTP configuration */
	routersInit := router.SetupRouter()
	routersInit.Run(setting.AppSetting.Port)

}
