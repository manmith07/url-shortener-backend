package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manmithsm/url-shortener/internal/services"
)

type URLHandler struct {
	Service *services.URLService
}

func NewURLHandler(s *services.URLService) *URLHandler {
	return &URLHandler{Service: s}
}

func (h *URLHandler) Shorten(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	code, err := h.Service.CreateShortURL(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/" + code,
	})
}

func (h *URLHandler) Redirect(c *gin.Context) {
	code := c.Param("shortCode")

	url, err := h.Service.GetOriginalURL(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.Redirect(http.StatusFound, url)
}
