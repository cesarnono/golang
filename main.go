package main

import (
	"godev.com/todo-app/application"
	"godev.com/todo-app/domain/model"
	"godev.com/todo-app/infrastructure/repository/mysql"
	"godev.com/todo-app/infrastructure/rest"
	"godev.com/todo-app/migration"
	"gofr.dev/pkg/gofr"
)

var taskController rest.TaskController

func main() {
    app := gofr.New()

    app.Migrate(migration.All())
    app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

        return "Hello World!", nil
    })
    app.POST("/task", func(ctx *gofr.Context) (interface{}, error) {
      return taskController.CreateTask(ctx);
    })

    app.PUT("/task/{id}/update", func(ctx *gofr.Context) (interface{}, error) {
        return taskController.UpdateTask(ctx);
    })

   app.Run()
}


func init() {
    
	taskMySqlRepository := mysql.TaskMySqlRepository{}

    taskService := model.NewTaskService(taskMySqlRepository)

    taskUseCase := application.NewTaskUseCase(taskService)

    taskController= rest.TaskController{TaskUseCase: taskUseCase}
}



