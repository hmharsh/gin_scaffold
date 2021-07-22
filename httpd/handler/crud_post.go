package handler

import (
	"ginapp/platrorm/crud"
	"net/http"

	"github.com/gin-gonic/gin"
)

type crudPostRequest struct {
	Name    string `json:name`
	Surname string `json:surname`
}

func PostCrud(data *crud.AllItems) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := crudPostRequest{}
		c.Bind(&requestBody)
		item := crud.Item{
			Name:    requestBody.Name,
			Surname: requestBody.Surname,
		}
		data.Add(item)
		c.Status(http.StatusNoContent)
	}
}
