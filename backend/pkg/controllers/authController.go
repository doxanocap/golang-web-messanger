package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	firstname := ctx.Query("firstname")
	lastname := ctx.Query("lastname") 
	ctx.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
