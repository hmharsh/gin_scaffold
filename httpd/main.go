package main

import (
	"database/sql"
	"ginapp/configuration"
	"ginapp/connection"
	"ginapp/httpd/handler"
	"ginapp/storage"
	"os/signal"
	"syscall"

	"ginapp/platrorm/goginrestapi"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	log.Infof("Loading config...")
	config := configuration.New()
	log.Infof("Starting ginapp...")
	r := gin.Default()
	log.Infof("Initiating the database connection...")
	db, err := connection.NewConnection(config.Database)
	if err != nil {
		log.Fatalf("Cannot connect to the database, err: %v", err)
	}
	defer db.Close()
	repository := storage.New(db)
	go handleShutdown(db)
	log.Infof("Connected to database successfully")

	log.Infof("Adding the endpoints and ApiGroup")
	r.GET("/ping", handler.PingGet())
	restApi := goginrestapi.New(repository)
	v1 := r.Group("/v1")
	{
		v1.GET("/goginrestapi", handler.GoGinRestApiGet(restApi))
		v1.POST("/goginrestapi", handler.GoGinRestApiPost(restApi))
	}

	log.Infof("Starting gin webserver")
	//Listen and serve on 0.0.0.0:8080
	if err := r.Run(); err != nil {
		log.Fatalf("Cannot start server, err: %v", err)
	}
}

func handleShutdown(db *sql.DB) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	// handle ctrl+c event here
	// for example, close database
	log.Warn("Closing DB connection before complete shutdown")
	if err := db.Close(); err != nil {
		log.Errorf("error while closing the connection to the database: %v", err)
	}
	os.Exit(0)
}
