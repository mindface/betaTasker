package service

import (
	"github.com/godotask/model"
	"github.com/godotask/infrastructure/repository"
)

type BookService struct {
  Repo repository.BookRepositoryInterface
}

func (s *BookService) CreateBook(book *model.Book) error {
	return s.Repo.Create(book)
}
func (s *BookService) GetBookByID(id string) (*model.Book, error) {
	return s.Repo.FindByID(id)
}
func (s *BookService) ListBooks(userID uint) ([]model.Book, error) {
	return s.Repo.FindAll(userID)
}
func (s *BookService) UpdateBook(id string, book *model.Book) error {
	return s.Repo.Update(id, book)
}
func (s *BookService) DeleteBook(id string) error {
	return s.Repo.Delete(id)
}
