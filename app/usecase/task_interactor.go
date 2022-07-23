package usecase

import "github.com/aux-Issa/Next_Go_practice/app/domain"

// 依存性逆転を防ぐために同じディレクトリにrepositoryを作成する
type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *TaskInteractor) Add(u domain.Task) (err error) {
	err = interactor.TaskRepository.Store(u)
	return
}

func (interactor *TaskInteractor) Tasks() (tasks domain.Tasks, err error) {
	tasks, err = interactor.TaskRepository.FindAll()
	return tasks, err
}

func (interactor *TaskInteractor) TaskById(id int) (task domain.Task, err error) {
	task, err = interactor.TaskRepository.FindById(id)
	return task, err
}
