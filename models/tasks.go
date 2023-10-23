package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID               string `json:"-" gorm:"type:uuid;primaryKey"`
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

type TaskService struct {
	Service ITaskService
}

type GormTaskService struct {
	DB *gorm.DB
}

func (ts *GormTaskService) FetchAllTasks(userID string) (interface{}, error) {
	var tasks []TaskResponse

	if err := ts.DB.Table("tasks").Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, errors.New("no tasks found")
	}

	return tasks, nil
}

func (ts *GormTaskService) FetchTask(taskID string, userID string) (interface{}, error) {
	task := &TaskResponse{}

	if err := ts.DB.Table("tasks").Where("id = ? AND user_id = ?", taskID, userID).First(task).Error; err != nil {
		return nil, errors.New("failed to fetch task")
	}

	return task, nil
}

func (ts *GormTaskService) CreateTask(task *Task) (interface{}, error) {
	task.ID = uuid.New().String()

	if err := ts.DB.Table("tasks").Create(&task).Error; err != nil {
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

func (ts *GormTaskService) DeleteTask(taskID string, userID string) (interface{}, error) {
	existingTask := &TaskResponse{}

	if err := ts.DB.Table("tasks").Where("id = ? AND user_id = ?", taskID, userID).First(&existingTask).Error; err != nil {
		return nil, errors.New("task does not exist")
	}

	if err := ts.DB.Table("tasks").Where("id = ? AND user_id = ?", taskID, userID).Delete(existingTask).Error; err != nil {
		return nil, errors.New("failed to delete task")
	}

	return nil, nil
}

func (ts *GormTaskService) UpdateTask(task *Task) (interface{}, error) {

	existingTask := &Task{}
	if err := ts.DB.Table("tasks").Where("id = ? AND user_id = ?", task.ID, task.UserID).First(&existingTask).Error; err != nil {
		return nil, errors.New("task does not exist")
	}

	task.UpdatedAt = time.Now()

	if err := ts.DB.Table("tasks").Model(&Task{}).Where("id = ? AND user_id = ?", task.ID, task.UserID).Updates(&task).Error; err != nil {
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
