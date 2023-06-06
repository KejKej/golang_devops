package main

import (
	"log"
	"net/http"
	"oauth2-okta/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading our custom .env file")
	}

	log.Printf("Loaded our env variables")

	router := initRouter()
	router.Run(":8080")
}

func ProtectedHello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello from protected endpoint"})
}

func PublicHello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello from public endpoint"})
}

func initRouter() *gin.Engine {
	router := gin.Default()

	public := router.Group("/public")
	public.GET("", func(c *gin.Context) {
		PublicHello(c)
	})

	protected := router.Group("/protected")
	protected.Use(middleware.OAuth2Middleware())
	protected.GET("", func(c *gin.Context) {
		ProtectedHello(c)
	})

	return router
}
