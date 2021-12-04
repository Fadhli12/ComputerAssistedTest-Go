package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	router.GET("/:id", paramHandler)
	router.Run(":8888")
}

func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"test" : "test",
	})
}

func paramHandler(c *gin.Context){
	param := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"ini" : param,
	})
}