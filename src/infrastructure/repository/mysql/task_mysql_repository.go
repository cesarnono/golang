package mysql

import (
	"godev.com/todo-app/src/domain/model"
	"godev.com/todo-app/src/infrastructure/rest"
)

type TaskMySqlRepository struct {
}

func (repository TaskMySqlRepository) Get(id string) *model.Task {
    task := &model.Task{};
	rest.Context.SQL.Select(rest.Context, task, "SELECT * FROM task where id=?",id)
	return task
   }

func (repository TaskMySqlRepository) Save(task model.Task) {
	rest.Context.SQL.ExecContext(rest.Context, "INSERT INTO task (id, description, status,date_created, date_updated) VALUES (?,?,?,?,?)",
	 task.GetId(), task.GetDescription(), task.GetStatus(), task.GetDateCreated(), task.GetDateUpdated())
}

func (repository TaskMySqlRepository) Update(task model.Task) {
	rest.Context.SQL.ExecContext(rest.Context, "UPDATE task set description =?, status =?, date_updated=? where id=?",
	task.GetDescription(), task.GetStatus(),task.GetDateUpdated(), task.Id)
}