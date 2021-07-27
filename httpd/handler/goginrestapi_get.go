package handler

import (
	"ginapp/platrorm/goginrestapi"
	"ginapp/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoGinRestApiGet(restApi goginrestapi.Getter, repository *storage.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := restApi.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
