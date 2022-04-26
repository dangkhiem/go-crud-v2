package router

import (
	"golang-crud/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	rUser := r.Group("/users")
	{
		rUser.GET("", controllers.GetListUser)
		rUser.POST("", controllers.CreateUser)
		rUser.GET("/:id", controllers.GetUser)
		rUser.PUT("/:id", controllers.UpdateUser)
		rUser.DELETE("/:id", controllers.DeleteUser)
	}
}
