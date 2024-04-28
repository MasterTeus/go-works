package router

import "github.com/gin-gonic/gin"

func Initalize() {
	// initicializa o Router usando as configs padrão do Gin
	router := gin.Default()
	// Definindo uma rota get
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
