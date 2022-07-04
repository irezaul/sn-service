package routes

import (
	"ginorm/database"
	"ginorm/handlers"

	"github.com/gin-gonic/gin"
)

func Route() {
	r := gin.Default()
	database.Connection()
	r.GET("/api", handlers.Home)

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

	frontend := r.Group("/")
	{
		frontend.GET("/", handlers.FrontendService)
		// frontend.GET("/login", handlers.FrontendLogin)
		// frontend.GET("/logout", handlers.FrontendLogout)
		// frontend.GET("/add", handlers.FrontendAddService)
		// frontend.GET("/update", handlers.FrontendUpdateService)
	}

	r.Run(":8080")

}
