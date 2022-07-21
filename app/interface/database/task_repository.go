package database

import (
	"github.com/aux-Issa/Next_Go_practice/app/domain"
)

import 

type TaskRepository interface {
	SqlHandler
}

func (repo *TaskRepository) Store(task domain.Task) (err error) {
	_, err := repo.Execute(
		"INSERT INTO tasks (title, content) VALUES(?,?)",
		task.Title,
		task.Content
	)
	if err != nil {
		return
	}
	return
}

func (repo *TaskRepository) FindById(id int)(task domain.Task, err error){
	row, err := repo.Query("SELECT id, title, content FROM tasks WHERE id=?", id)
	defer row.Close()
	if err != nil {
		return		
	}
	var id int
	var title string
	var content string
	row.Next()
	// idで検索しているのでrowは単数系
	err = row.Scan(&id, &title, &content)
	if err != nil {
		return
	}

	task.ID = id
	task.Title = title
	task.Content = content
	return
 }
