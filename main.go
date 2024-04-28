package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// initicializa o Router usando as configs padrão do Gin
	r := gin.Default()
	// Definindo uma rota get
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
