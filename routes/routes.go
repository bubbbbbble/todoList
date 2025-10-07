package routes

import (
	"todolist/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	r.POST("/user/register", controllers.Register)
	r.POST("/user/login", controllers.Login)
	auth := r.Group("/")
	auth.Use(controllers.JWTAuthMiddleware())
	{
		auth.POST("/todo/add", controllers.AddTodo)
		auth.GET("/todo/list", controllers.GetTodoList)
		auth.PUT("/todo/update", controllers.UpdateTodo)
		auth.DELETE("/todo/delete", controllers.DeleteTodo)
	}

}
