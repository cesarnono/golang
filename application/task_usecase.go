package application

import (
	"godev.com/todo-app/domain/model"
)

type TaskUseCase struct {
	taskService model.TaskService
}

func (usecase *TaskUseCase) CreateTask(description string) (model.Task, error) {
	return usecase.taskService.Create(description)
}

func (usecase *TaskUseCase) UpdateTask(request model.TaskRequest) (string, error) {
	return usecase.taskService.Update(request)
}

func NewTaskUseCase(taskService model.TaskService) TaskUseCase{
	return TaskUseCase{taskService: taskService}
}