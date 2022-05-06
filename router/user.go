package router

import (
	"golang-crud/controllers"
	"golang-crud/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	rUser := r.Group("/users").Use(middleware.JwtTokenCheck)
	{
		rUser.GET("", controllers.GetListUser)
		rUser.POST("", controllers.CreateUser)
		rUser.GET("/:id", controllers.GetUser)
		rUser.PUT("/:id", controllers.UpdateUser)
		rUser.DELETE("/:id", controllers.DeleteUser)
	}
}
