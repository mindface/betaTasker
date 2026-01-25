package service

import (
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/infrastructure/db/repository"
)

type TaskService struct {
	Repo repository.TaskRepositoryInterface
}

func (s *TaskService) CreateTask(task *model.Task) error {
	return s.Repo.Create(task)
}
func (s *TaskService) GetTaskByID(id string) (*model.Task, error) {
	return s.Repo.FindByID(id)
}
func (s *TaskService) ListTasks(userID uint) ([]model.Task, error) {
	return s.Repo.FindAll(userID)
}
// ListTasksByUser: 特定ユーザーのタスク一覧を取得
func (s *TaskService) ListTasksByUser(userID uint) ([]model.Task, error) {
    return s.Repo.ListTasksByUser(userID)
}
// ListTasksByUserPager: 特定ユーザーのタスク一覧をページネーション取得
func (s *TaskService) ListTasksByUserPager(userID uint, page int, perPage int, offset int) ([]model.Task, int64, error) {
    return s.Repo.ListTasksByUserPager(userID, offset, perPage)
}
func (s *TaskService) UpdateTask(id string, task *model.Task) error {
	return s.Repo.Update(id, task)
}
func (s *TaskService) DeleteTask(id string) error {
	return s.Repo.Delete(id)
}
