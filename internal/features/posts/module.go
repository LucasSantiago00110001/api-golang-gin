package posts

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router gin.IRoutes, db *sql.DB) {
	handler := NewHandler(NewService(NewRepository(db)))
	router.GET("/posts", handler.GetPosts)
}
