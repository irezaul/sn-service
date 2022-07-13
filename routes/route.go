package routes

import (
	"ginorm/database"
	"ginorm/handlers"

	// "github.com/gin-gonic/contrib/static"

	"github.com/gin-gonic/gin"
)

func App() {
	r := gin.Default()
	database.Connection()
	r.GET("/api", handlers.Home)

	// r.Use(static.Serve("/asset", static.LocalFile("views/frontend/assets", true)))
	r.LoadHTMLGlob("templates/*.gohtml")

	user := r.Group("/api/user")
	{

		user.POST("/", handlers.AddUser)
		user.GET("/", handlers.GetUser)
		user.DELETE("/:id", handlers.DeleteUser)
		user.PUT("/:id", handlers.UpdateUser)

	}
	auth := r.Group("/api/auth/login")
	{
		auth.POST("/", handlers.LoginAuth)
	}

	service := r.Group("/api/service")
	{
		service.POST("/", handlers.AddService)
		service.GET("/", handlers.GetService)
		service.GET("/:id", handlers.GetServiceById)
		service.DELETE("/:id", handlers.DeleteService)
		service.PUT("/:id", handlers.UpdateService)

	}

	frontend := r.Group("/")
	{
		frontend.GET("/", handlers.FrontendService)
		// frontend.GET("/login", handlers.FrontendLogin)
		// frontend.GET("/logout", handlers.FrontendLogout)
		frontend.GET("/add", handlers.FrontendAddService)
		frontend.POST("/add", handlers.FrontendAddProcess)
		frontend.GET("/edit/:id", handlers.FrontendEditService)
		// frontend.GET("/update", handlers.FrontendUpdateService)
	}

	r.Run(":8080")

}
