package router

import "github.com/gin-gonic/gin"

func Initalize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes GET/POST/DELET...
	initializeRoute(router)

	// Run the server
	router.Run(":8080")
}
