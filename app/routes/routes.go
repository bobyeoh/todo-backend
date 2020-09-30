package routes

import (
	"todo/app"
	"todo/app/controllers"
	"todo/app/permission"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Routes godoc
func Routes(server *app.Server) {
	user := controllers.InitUser(server)
	column := controllers.InitColumn(server)
	task := controllers.InitTask(server)

	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST("/user/login", user.Login)

	server.Echo.GET("/column", column.GetColumns, permission.Auth(server))

	server.Echo.GET("/task", task.GetTasks, permission.Auth(server))
	server.Echo.POST("/task", task.CreateTask, permission.Auth(server))
	server.Echo.PUT("/task", task.UpdateTask, permission.Auth(server))
	server.Echo.DELETE("/task", task.DeleteTask, permission.Auth(server))
}
