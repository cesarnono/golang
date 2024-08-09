package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type statusTask string


const(
	CREATED statusTask ="CREATED"
	WIP statusTask="WIP"
	COMPLETED statusTask="COMPLETED"
)

type TaskRepository interface {
   Get(id string) (*Task)
   Save (task Task)
   Update (task Task)
}

type TaskRequest struct {
	TaskId string  `json:"id"`
	TaskDescription string `json:"description"`
	Status string `json:"status"`
}

type TaskService struct {
 taskRepository TaskRepository 
}

type Task struct {
	Id string 
	Description string
	Status statusTask
	DateCreated time.Time
	DateUpdated time.Time
}

func (task Task) GetId() string {
	return task.Id
}

func (task Task) GetDescription() string {
	return task.Description
}

func (task Task) GetStatus() statusTask {
	return task.Status
}

func (task Task) GetDateCreated() time.Time {
	return task.DateCreated
}

func (task Task) GetDateUpdated() time.Time {
	return task.DateUpdated
}

func(taskService *TaskService) Create(description string) (Task, error){
	if strings.Trim(description, " ") == ""{
		return Task{},fmt.Errorf("invalid description task")
    }
	task := newTask(description)
	taskService.taskRepository.Save(task)
	return task, nil
}

func (taskService * TaskService)GetById(id string) Task {
	return *taskService.taskRepository.Get(id)
}

func validateChangeStatus(currentStatus, nextStatus string) error{
	if currentStatus == "CREATED" && nextStatus == "WIP" {
		return nil
	}
	if currentStatus == "WIP" && nextStatus == "COMPLETED" {
		return nil
	}

	if (currentStatus == "WIP" || currentStatus == "COMPLETED") && currentStatus == nextStatus {
		return nil
	}

	return fmt.Errorf("invalid change status")
}

func(taskService *TaskService) Update(request TaskRequest) (string, error) {
	if err := validateRequest(request); err != nil {
		return "", ValidationError{error: err.Error()}
	}
	task := taskService.GetById(request.TaskId)
	if(task == Task{}) {
		return "", NotFoundError{error: fmt.Sprintf("task: %s not found", request.TaskId)}
	}
	if err := validateChangeStatus(string(task.Status), request.Status); err != nil{
		return "", ValidationError{error: err.Error()}
	}
	task.Description=request.TaskDescription
	task.Status= statusTask(request.Status)
	task.DateUpdated= time.Now()
	taskService.taskRepository.Update(task)
	return fmt.Sprintf("task %s have been updated", task.GetId()),nil
}

func NewTaskService(taskRepository TaskRepository) TaskService{
	return TaskService{
		taskRepository: taskRepository,
	}
}

func newTask(description string) Task {
	task := Task{Id: uuid.NewString(), Description: description,
	Status: CREATED, DateCreated: time.Now(),DateUpdated: time.Now(),}
	return task
}

func validateRequest(request TaskRequest) error {
	if strings.Trim(request.TaskId, " ") == ""{ 
         return fmt.Errorf("invalid taskId")
	}

	if strings.Trim(request.TaskDescription, " ") == ""{
		return fmt.Errorf("invalid description task")
    }

	if strings.Trim(request.Status, " ") == "" {
		return fmt.Errorf("invalid status task")
	}

	return nil
}

