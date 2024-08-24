package mysql

import (
	"fmt"

	"godev.com/todo-app/domain/model"
	"godev.com/todo-app/infrastructure/rest"
)

type TaskMySqlRepository struct {
}

func (repository TaskMySqlRepository) Get(id string) *model.Task {
    task := &model.Task{};
	rest.Context.SQL.Select(rest.Context, task, "SELECT * FROM task where id=?",id)
	return task
   }

func (repository TaskMySqlRepository) Save(task model.Task) {
	res, err := rest.Context.SQL.ExecContext(rest.Context, "INSERT INTO task (id, description, status,date_created, date_updated) VALUES ($1, $2, $3, $4, $5)",task.GetId(),
	 task.GetDescription(), task.GetStatus(), task.GetDateCreated(), task.GetDateUpdated())
	 if err != nil {
		fmt.Printf("error: %s", err.Error());
		return
	 }

	 rows, err := res.RowsAffected()
	 if err != nil {
		fmt.Printf("error: %s", err.Error());
		return
	 }

	 fmt.Print(rows)
}

func (repository TaskMySqlRepository) Update(task model.Task) {
	rest.Context.SQL.ExecContext(rest.Context, "UPDATE task set description =?, status =?, date_updated=? where id=?",
	task.GetDescription(), task.GetStatus(),task.GetDateUpdated(), task.Id)
}