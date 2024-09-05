package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
}
