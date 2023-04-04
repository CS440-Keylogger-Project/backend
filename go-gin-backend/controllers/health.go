package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

// @Summary Health Check
// @Description Health Check Endpoint that doesn't require authentication
// @Tags health
// @Success 200 {object} string	"Success"
// @Failure 404 {object} models.HTTPError
// @Router /health [get]
func (h HealthController) GetStatus(c *gin.Context) {
	log.Println("Testing Health")
	c.String(http.StatusOK, "Success")
}
