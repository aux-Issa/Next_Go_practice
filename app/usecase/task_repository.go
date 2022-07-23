package usecase

import "github.com/aux-Issa/Next_Go_practice/app/domain"

type TaskRepository interface {
	Store(task domain.Task) error
	FindAll() (domain.Tasks, error)
	FindById(id int) (domain.Task, error)
}
