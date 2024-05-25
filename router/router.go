package router

import "github.com/gin-gonic/gin"

// Quando iniciamos um metodo em maiusculo ele automaticamente esta sendo exportado
func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	router.Run(":3333")
}
