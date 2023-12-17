package routes

import (
	controller "github.com/jamyMarkos/gin-auth/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jamyMarkos/gin-auth/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine)()  {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	
}

