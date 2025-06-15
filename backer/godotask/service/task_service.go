package service

import (
	"github.com/godotask/model"
	"github.com/godotask/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func (s *TaskService) CreateTask(task *model.Task) error {
	return s.Repo.Create(task)
}
func (s *TaskService) GetTaskByID(id string) (*model.Task, error) {
	return s.Repo.FindByID(id)
}
func (s *TaskService) ListTasks() ([]model.Task, error) {
	return s.Repo.FindAll()
}
func (s *TaskService) UpdateTask(id string, task *model.Task) error {
	return s.Repo.Update(id, task)
}
func (s *TaskService) DeleteTask(id string) error {
	return s.Repo.Delete(id)
}
