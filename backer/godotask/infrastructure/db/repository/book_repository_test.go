package repository

import (
	"fmt"
	"testing"

	"github.com/godotask/infrastructure/db/model"
	"github.com/stretchr/testify/assert"
	"github.com/godotask/infrastructure/db/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to in-memory DB: %v", err)
	}
	// テーブル作成
	err = db.AutoMigrate(&model.Book{})
	if err != nil {
		t.Fatalf("Failed to migrate Book model: %v", err)
	}
	return db
}

func intToStr(id int) string {
	return fmt.Sprintf("%d", id)
}

func TestBookRepositoryCRUD(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewBookRepository(db)

	book := &model.Book{
		Title:   "Test Title",
		Name:    "Test Author",
		Text:    "Test content",
		Disc:    "Test description",
		ImgPath: "/path/to/image.jpg",
		Status:  "draft",
	}

	// --- Create ---
	err := repo.Create(book)
	assert.NoError(t, err)
	assert.NotZero(t, book.ID)

	// --- FindByID ---
	found, err := repo.FindByID(intToStr(book.ID))
	assert.NoError(t, err)
	assert.Equal(t, book.Title, found.Title)
	assert.Equal(t, book.Name, found.Name)

	// --- FindAll ---
	books, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, books, 1)

	// --- Update ---
	book.Status = "published"
	book.Title = "Updated Title"
	err = repo.Update(intToStr(book.ID), book)
	assert.NoError(t, err)

	updated, err := repo.FindByID(intToStr(book.ID))
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", updated.Title)
	assert.Equal(t, "published", updated.Status)

	// --- Delete ---
	err = repo.Delete(intToStr(book.ID))
	assert.NoError(t, err)

	deleted, err := repo.FindByID(intToStr(book.ID))
	assert.Error(t, err)
	assert.Nil(t, deleted)
}
