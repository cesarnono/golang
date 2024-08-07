package config

import (
	"godev.com/todo-app/src/infrastructure/rest"
	"gofr.dev/pkg/gofr"
)


var Context *gofr.Context

type MyHandler struct{
    TaskController rest.TaskController
}


func(handler *MyHandler) HandlerCreateTask(ctx *gofr.Context) (interface{}, error) {
	Context=ctx
	return handler.TaskController.CreateTask(ctx)
}

func(handler *MyHandler) HandlerUpdateTask(ctx *gofr.Context) (interface{}, error) {
	Context=ctx
	return handler.TaskController.UpdateTask(ctx)
}






