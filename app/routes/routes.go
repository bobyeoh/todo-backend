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
	prefix := "/api"
	user := controllers.InitUser(server)
	column := controllers.InitColumn(server)
	task := controllers.InitTask(server)

	server.Echo.Use(middleware.Logger())

	server.Echo.GET(prefix+"/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST(prefix+"/user/login", user.Login)

	server.Echo.GET(prefix+"/user/logout", user.Logout)

	server.Echo.GET(prefix+"/column", column.GetColumns, permission.Auth(server))

	server.Echo.GET(prefix+"/task/:column_id", task.GetTasks, permission.Auth(server))
	server.Echo.POST(prefix+"/task", task.CreateTask, permission.Auth(server))
	server.Echo.PUT(prefix+"/task/:id", task.UpdateTask, permission.Auth(server))
	server.Echo.DELETE(prefix+"/task/:id", task.DeleteTask, permission.Auth(server))
}
