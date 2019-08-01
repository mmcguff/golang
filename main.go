package main

import "github.com/gin-gonic/gin"

//I should create a full API using this or clone one.

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/dog", func( c *gin.Context){
		c.JSON(200, gin.H{
			"message": "shit",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}