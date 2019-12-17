package fooController

import (
	"github.com/gin-gonic/gin"
)

func HandleGetFoo(c *gin.Context) {
	println("foo request")
	c.JSON(200, gin.H{
		"foo": "bar",
	})
}
