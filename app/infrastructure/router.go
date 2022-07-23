package infrastructure

import (
	"database/sql"

	"github.com/aux-Issa/Next_Go_practice/app/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	e := echo.New()
	taskController := controller.NewTaskController(NewSqlHandler(&sql.DB{}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/tasks", func(c echo.Context) error {
		return taskController.Index(c)
	})
	e.GET("/task/:id", func(c echo.Context) error {
		return taskController.Show(c)
	})
	e.POST("/create", func(c echo.Context) error {
		return taskController.Show(c)
	})

	e.Logger.Fatal(e.Start(":12345"))
}
