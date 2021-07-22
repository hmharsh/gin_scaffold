package main

import (
	"ginapp/httpd/handler"
	"ginapp/platrorm/crud"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handler.PingGet())

	data := crud.New()
	r.GET("/crud", handler.GetCrud(data))
	r.POST("/crud", handler.PostCrud(data))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// testData := crud.New()
	// fmt.Println(testData)
	// testData.Add(crud.Item{Name: "Harshit", Surname: "harshit"})
	// fmt.Println(testData)

	// Example POST request
	// {
	// 	"name": "harshit",
	// 	"surname": "mahajan"
	// }
}
