package router

import (
	"golang-crud/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	rAuth := r.Group("/")
	{
		rAuth.POST("login", controllers.Login)
	}
}
