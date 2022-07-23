package usecase

import "github.com/aux-Issa/Next_Go_practice/app/domain"

type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *TaskInteractor) Add(u domain.Task) (err error) {
	err = interactor.TaskRepository.Store(u)
	return
}

func (interactor *TaskInteractor) Tasks(tasks domain.Tasks) (err error) {
	tasks, err = interactor.TaskRepository.FindAll()
	return
}
func (interactor *TaskInteractor) TaskById(id int) (task domain.Task, err error) {
	task, err = interactor.TaskRepository.FindById(id)
	return task, err
}
