package routes

import (
	"ginorm/database"
	"ginorm/handlers"

	"github.com/gin-gonic/gin"
)

func Route() {
	r := gin.Default()
	database.Connection()
	r.GET("/", handlers.Home)

	user := r.Group("/user")
	{

		user.POST("/", handlers.AddUser)
		user.GET("/", handlers.GetUser)
		user.DELETE("/:id", handlers.DeleteUser)
		user.PUT("/:id", handlers.UpdateUser)

	}
	auth := r.Group("/auth/login")
	{
		auth.POST("/", handlers.LoginAuth)
	}

	service := r.Group("/auth/service")
	{
		service.POST("/", handlers.AddService)
		service.GET("/", handlers.GetService)
		service.DELETE("/:id", handlers.DeleteService)
		service.PUT("/:id", handlers.UpdateService)

	}

	fontend := r.Group("/")
	{
		fontend.GET("/", handlers.FontendService)
		// fontend.GET("/login", handlers.FontendLogin)
		// fontend.GET("/logout", handlers.FontendLogout)
		// fontend.GET("/add", handlers.FontendAddService)
		// fontend.GET("/update", handlers.FontendUpdateService)
	}

	r.Run(":8080")

}
