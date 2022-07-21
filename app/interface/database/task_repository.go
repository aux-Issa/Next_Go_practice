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

func (repo *TaskRepository) FindAll(tasks domain.Tasks, err error){
	rows, err := repo.Query("SELECT id, title, conetent FORM tasks")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next(){
		var id int
		var title string
		var content string
		if err = rows.Scan(&id, &title, &content); err != nil {
			return
		}
		task := domain.Task{
			ID:       id,
			Title:    title,
			Content:  content,
		}
		tasks = append(tasks, task)
	}
	return tasks
}