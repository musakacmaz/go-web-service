package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary check health
// @ID check-health
// @Produce json
// @Success 200 {string} string "OK"
// @Router /health [get]
func GetHealth(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "OK")
}
