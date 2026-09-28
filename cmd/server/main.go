package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	_ "api-golang-gin/docs"
	"api-golang-gin/internal/database"
	"api-golang-gin/internal/features/posts"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sannockit API
// @version 1.0
// @description API do Sannockit.
// @BasePath /
//
// Health godoc
// @Summary Verifica a saúde da API
// @Description Retorna o estado atual da API.
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func main() {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatalf("não foi possível carregar .env: %v", err)
	}

	db, err := database.OpenPostgres()
	if err != nil {
		log.Fatalf("não foi possível conectar ao PostgreSQL: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/health", healthHandler)
	posts.RegisterRoutes(router, db)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("API rodando em http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
