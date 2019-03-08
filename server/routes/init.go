package routes

import "github.com/gin-gonic/gin"

func InitRoutes(e *gin.Engine) {
	InitUserRoute(e.Group("/users"))
	InitAuthRoute(e.Group("/auth"))
}