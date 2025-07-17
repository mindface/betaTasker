package repository_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/godotask/model"
	"github.com/godotask/repository"
)

func setupTaskTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test DB: %v", err)
	}
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		t.Fatalf("failed to migrate Task model: %v", err)
	}
	return db
}

func TestTaskRepository_CRUD(t *testing.T) {
	db := setupTaskTestDB(t)
	repo := &repository.TaskRepositoryImpl{DB: db}

	// Create
	now := time.Now()
	task := &model.Task{
		UserID:      1,
		MemoryID:    nil,
		Title:       "Test Task",
		Description: "This is a test task.",
		Date:        &now,
		Status:      "todo",
		Priority:    1,
	}
	err := repo.Create(task)
	assert.NoError(t, err)
	assert.NotZero(t, task.ID)

	idStr := strconv.Itoa(task.ID)

	// FindByID
	found, err := repo.FindByID(idStr)
	assert.NoError(t, err)
	assert.Equal(t, task.Title, found.Title)
	assert.Equal(t, task.Status, found.Status)

	// Update
	task.Title = "Updated Task"
	err = repo.Update(idStr, task)
	assert.NoError(t, err)

	updated, err := repo.FindByID(idStr)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", updated.Title)

	// FindAll
	all, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, all, 1)

	// Delete
	err = repo.Delete(idStr)
	assert.NoError(t, err)

	_, err = repo.FindByID(idStr)
	assert.Error(t, err) // should return not found error
}
