package handler

import (
	"ginapp/platrorm/goginrestapi"
	"ginapp/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type goGinRestApiPostRequest struct {
	Title string `json: "title"`
	Post  string `json:"post"`
}

func GoGinRestApiPost(restApi goginrestapi.Adder, repository *storage.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := goGinRestApiPostRequest{}
		c.Bind(&requestBody)

		item := goginrestapi.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}

		//config := configuration.New()

		restApi.Add(item)
		c.Status(http.StatusNoContent)

	}
}
