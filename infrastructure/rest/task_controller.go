package rest

import (
	"godev.com/todo-app/application"
	"godev.com/todo-app/domain/model"
	"gofr.dev/pkg/gofr"
)

var Context *gofr.Context

type TaskController struct {
	TaskUseCase application.TaskUseCase
}


type TaskResponse struct {
	ID string `json:"id"`
	Description string `json:"description"`
}

func (controller TaskController) CreateTask(ctx *gofr.Context) (interface{}, error) {
	Context = ctx
    var request model.TaskRequest
	ctx.Bind(&request)
    task, err := controller.TaskUseCase.CreateTask(request.TaskDescription)

	if err != nil {
		return nil,err
	}

	return TaskResponse{
		ID: task.GetId(),
		Description: task.GetDescription(),
	}, nil
}

func (controller TaskController) UpdateTask(ctx *gofr.Context) (interface{}, error) {
	Context = ctx
	var request model.TaskRequest
    taskId:= ctx.Request.PathParam("id")
    ctx.Bind(&request)

	return controller.TaskUseCase.UpdateTask(model.TaskRequest{TaskId: taskId, 
		TaskDescription: request.TaskDescription,
		Status: request.Status,
	})
}


