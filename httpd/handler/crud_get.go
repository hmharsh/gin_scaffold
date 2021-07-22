package handler

import (
	"ginapp/platrorm/crud"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCrud(data *crud.AllItems) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, data.List())
	}
}
