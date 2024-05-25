package router

import "github.com/gin-gonic/gin"

// Quando iniciamos um metodo em maiusculo ele automaticamente esta sendo exportado
func Initialize() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":3333")
}
