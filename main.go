package main

import (
	"github.com/DanielUhlik/go-docker-test/controllers/fooController"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/foo", fooController.HandleGetFoo)

	app.Run()
}
