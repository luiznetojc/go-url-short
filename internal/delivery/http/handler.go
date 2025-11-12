package http

import (
	"net/http"
	"github.com/luizn/go-url-short/internal/usercase"
	"github.com/gin-gonic/gin"
)


type UrlHandler struct{
	CreateUseCase *usercase.CreateUrlUseCase
}

func NewUrlHandler(createUseCase *usercase.CreateUrlUseCase) *UrlHandler{
	return &UrlHandler{CreateUseCase: createUseCase}
}

func (h *UrlHandler) CreateUrl(c *gin.Context){
		var input struct {
		OriginalURL string `json:"original_url" binding:"required,url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	url, err := h.CreateUseCase.Execute(c.Request.Context(), input.OriginalURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, url)

}