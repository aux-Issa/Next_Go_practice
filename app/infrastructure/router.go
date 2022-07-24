package infrastructure

import (
	"github.com/aux-Issa/Next_Go_practice/app/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	e := echo.New()
	taskController := controller.NewTaskController(NewSqlHandler())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/tasks", func(c echo.Context) error {
		return taskController.Index(c)
	})
	e.GET("/tasks/:id", func(c echo.Context) error {
		return taskController.Show(c)
	})
	e.POST("/create", func(c echo.Context) error {
		return taskController.Create(c)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
