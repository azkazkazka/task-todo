package models

import (
	"time"

	"github.com/azkazkazka/task-todo/db"
)

type Task struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	UserID           uint   `json:"user_id"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	DueDateString    string `json:"due_date" gorm:"-"`
	DueDate          time.Time
	CompletionStatus bool      `json:"completion_status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func FetchTasks() (Response, error) {
	var res Response
	var tasks []Task
	con := db.CreateCon()

	if err := con.Find(&tasks).Error; err != nil {
		return res, err
	}

	res.Message = "Successfully fetched tasks"
	res.Data = tasks

	return res, nil
}

func CreateTask(task *Task) (Response, error) {
	var res Response
	con := db.CreateCon()

	task.CompletionStatus = false
	task.UserID = 1 // TODO: sambung sama user

	if err := con.Create(task).Error; err != nil {
		return res, err
	}

	res.Message = "Successfully created tasks"
	return res, nil
}

func DeleteTask(taskID uint) (Response, error) {
	var res Response
	con := db.CreateCon()
	existingTask := &Task{}

	if err := con.First(existingTask, taskID).Error; err != nil {
		return res, err
	}

	if err := con.Delete(existingTask).Error; err != nil {
		return res, err
	}

	res.Message = "Successfully deleted task"
	return res, nil
}

func UpdateTask(task *Task) (Response, error) {
	var res Response
	con := db.CreateCon()

	existingTask := &Task{}
	if err := con.First(existingTask, task.ID).Error; err != nil {
		return res, err
	}

	if err := con.Model(&Task{}).Where("id = ?", task.ID).Updates(task).Error; err != nil {
		return res, err
	}

	res.Message = "Successfully updated task"
	return res, nil
}
