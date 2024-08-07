package main

import (
	"godev.com/todo-app/src/application"
	"godev.com/todo-app/src/domain/model"
	"godev.com/todo-app/src/infrastructure/config"
	"godev.com/todo-app/src/infrastructure/repository/mysql"
	"godev.com/todo-app/src/infrastructure/rest"
	"gofr.dev/pkg/gofr"
)

var myCustomHandler config.MyHandler

func main() {
    app := gofr.New()
    app.POST("/task", func(ctx *gofr.Context) (interface{}, error) {
      return myCustomHandler.HandlerCreateTask(ctx);
    })

    app.PUT("/task/{id}/update", func(ctx *gofr.Context) (interface{}, error) {
        return myCustomHandler.HandlerUpdateTask(ctx);
    })

   app.Run()
}


func init() {
	taskMySqlRepository := mysql.TaskMySqlRepository{}

    taskService := model.NewTaskService(taskMySqlRepository)

    taskUseCase := application.NewTaskUseCase(taskService)

    taskController:= rest.TaskController{TaskUseCase: taskUseCase}

    myCustomHandler = config.MyHandler{TaskController: taskController}
}



