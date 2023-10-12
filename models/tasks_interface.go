package models

type ITaskService interface {
	FetchAllTasks(userID string) (interface{}, error)
	FetchTask(taskID, userID string) (interface{}, error)
	CreateTask(task *Task) (interface{}, error)
	DeleteTask(taskID, userID string) (interface{}, error)
	UpdateTask(task *Task) (interface{}, error)
}
