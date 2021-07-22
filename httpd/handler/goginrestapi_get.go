package handler

import (
	"ginapp/platrorm/goginrestapi"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoGinRestApiGet(restApi goginrestapi.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := restApi.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
