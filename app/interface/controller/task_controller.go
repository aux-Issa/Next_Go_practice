package controller

import (
	"strconv"

	"github.com/aux-Issa/Next_Go_practice/app/domain"
	"github.com/aux-Issa/Next_Go_practice/app/interface/database"
	"github.com/aux-Issa/Next_Go_practice/app/usecase"
	"github.com/labstack/echo"
)

type TaskController struct {
	Interactor usecase.TaskInteractor
}

func NewTaskController(sqlHandler database.SqlHandler) *TaskController {
	return &TaskController{
		Interactor: usecase.TaskInteractor{
			TaskRepository: &database.TaskRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TaskController) Create(c echo.Context) (err error) {
	u := domain.Task{}
	c.Bind(&u)

	err = controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, "complete")
	return
}

func (controller *TaskController) Index(c echo.Context) (err error) {
	tasks, err := controller.Interactor.Tasks()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, tasks)
	return
}

func (controller *TaskController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := controller.Interactor.TaskById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, task)
	return
}
