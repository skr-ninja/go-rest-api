package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rest-api/controllers"
	"github.com/rest-api/middlewares"
	"github.com/rest-api/models"
)

func main() {

	models.ConnectDataBase()
	// cfg, err := config.GetConfig()
	// if err != nil {
	// 	fmt.Println("Configuration Error")
	// }

	// config.GetDb(cfg)

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protectedUser := r.Group("/api/user")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	protectedUser.POST("/update", controllers.UpdateUser)
	protectedUser.DELETE("/delete", controllers.DeleteUser)

	r.Run(":8080")

}
