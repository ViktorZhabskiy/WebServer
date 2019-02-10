package main

import (
	"WebServer/handlers"
	"WebServer/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main(){
	r := gin.New()

	r.Use(middleware.LogRequest)

	r.GET("/dashboard", handlers.NewDashboard())
	r.NoRoute(handlers.NotFound)

	err := r.Run()
	if err != nil {
		logrus.Fatalf("Failed start web server. Err: %s", err.Error())
	}
}