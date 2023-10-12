package models

import (
	"errors"
	"time"

	"github.com/azkazkazka/task-todo/db"
)

type Task struct {
	ID               string `json:"id" gorm:"type:uuid;primaryKey"`
	UserID           string `json:"user_id"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	DueDateString    string `json:"due_date" gorm:"-"`
	DueDate          time.Time
	CompletionStatus bool      `json:"completion_status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type TaskRequest struct {
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

type TaskResponse struct {
	ID               string    `json:"id"`
	UserID           string    `json:"user_id"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	DueDate          time.Time `json:"due_date"`
	CompletionStatus bool      `json:"completion_status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func FetchAllTasks(userID string) (interface{}, error) {
	var tasks []TaskResponse
	con := db.CreateCon()

	if err := con.Table("tasks").Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, errors.New("no tasks found")
	}

	return tasks, nil
}

func FetchTask(taskID string, userID string) (interface{}, error) {
	con := db.CreateCon()
	task := &TaskResponse{}

	if err := con.Table("tasks").Where("id = ? AND user_id = ?", taskID, userID).First(task).Error; err != nil {
		return nil, errors.New("failed to fetch task")
	}

	return task, nil
}

func CreateTask(task *Task) (interface{}, error) {
	con := db.CreateCon()

	if err := con.Table("tasks").Create(task).Error; err != nil {
		return nil, errors.New("failed to create task")
	}

	data := TaskResponse{
		ID:               task.ID,
		UserID:           task.UserID,
		Title:            task.Title,
		Description:      task.Description,
		DueDate:          task.DueDate,
		CompletionStatus: task.CompletionStatus,
		CreatedAt:        task.CreatedAt,
		UpdatedAt:        task.UpdatedAt,
	}

	return data, nil
}

func DeleteTask(taskID string, userID string) (interface{}, error) {
	con := db.CreateCon()
	existingTask := &TaskResponse{}

	if err := con.Table("tasks").Where("id = ? AND user_id = ?", taskID, userID).First(existingTask).Error; err != nil {
		return nil, errors.New("task does not exist")
	}

	if err := con.Table("tasks").Where("id = ? AND user_id = ?", taskID, userID).Delete(existingTask).Error; err != nil {
		return nil, errors.New("failed to delete task")
	}

	return nil, nil
}

func UpdateTask(task *Task) (interface{}, error) {
	con := db.CreateCon()

	existingTask := &Task{}
	if err := con.Table("tasks").Where("id = ? AND user_id = ?", task.ID, task.UserID).First(existingTask).Error; err != nil {
		return nil, errors.New("task does not exist")
	}

	task.UpdatedAt = time.Now()

	if err := con.Table("tasks").Model(&Task{}).Where("id = ? AND user_id = ?", task.ID, task.UserID).Updates(task).Error; err != nil {
		return nil, errors.New("failed to delete task")
	}

	data := TaskResponse{
		ID:               task.ID,
		UserID:           task.UserID,
		Title:            task.Title,
		Description:      task.Description,
		DueDate:          task.DueDate,
		CompletionStatus: task.CompletionStatus,
		CreatedAt:        task.CreatedAt,
		UpdatedAt:        task.UpdatedAt,
	}

	return data, nil
}
