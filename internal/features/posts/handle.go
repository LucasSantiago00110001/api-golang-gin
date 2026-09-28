package posts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	GetPosts(c *gin.Context)
}

type HandlePost struct {
	service Service
}

func NewHandler(service Service) *HandlePost {
	return &HandlePost{service: service}
}

// GetPosts godoc
// @Summary Lista posts
// @Description Retorna os posts da página solicitada. Se page não for informada, utiliza 1.
// @Tags posts
// @Produce json
// @Param page query int false "Número da página" minimum(1) default(1)
// @Success 200 {object} ListResponse
// @Failure 400 {object} map[string]string "Parâmetro page inválido"
// @Failure 500 {object} map[string]string "Erro ao listar posts"
// @Router /posts [get]
func (h *HandlePost) GetPosts(c *gin.Context) {
	page := 1
	if rawPage := c.Query("page"); rawPage != "" {
		parsedPage, err := strconv.Atoi(rawPage)
		if err != nil || parsedPage < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "page deve ser um inteiro maior que zero"})
			return
		}
		page = parsedPage
	}

	response, err := h.service.List(c.Request.Context(), page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "não foi possível listar os posts"})
		return
	}

	c.JSON(http.StatusOK, response)
}
