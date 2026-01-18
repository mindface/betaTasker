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

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to in-memory DB: %v", err)
	}
	err = db.AutoMigrate(&model.Memory{})
	if err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func TestMemoryRepository_CRUD(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewMemoryRepository(db)

	// --- Create ---
	readDate := time.Now()
	memory := &model.Memory{
		UserID:            1,
		SourceType:        "book",
		Title:             "Test Title",
		Author:            "Author Name",
		Notes:             "Some notes",
		Tags:              "tag1,tag2",
		ReadStatus:        "read",
		ReadDate:          &readDate,
		Factor:            "Deep Insight",
		Process:           "Summarized",
		EvaluationAxis:    "Relevance",
		InformationAmount: "High",
	}
	err := repo.Create(memory)
	assert.NoError(t, err)
	assert.NotZero(t, memory.ID)

	intStr := strconv.Itoa(memory.ID)
	// --- FindByID ---
	found, err := repo.FindByID(intStr)
	assert.NoError(t, err)
	assert.Equal(t, memory.Title, found.Title)

	// --- Update ---
	memory.Title = "Updated Title"
	err = repo.Update(intStr, memory)
	assert.NoError(t, err)

	updated, err := repo.FindByID(intStr)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", updated.Title)

	// --- FindAll ---
	all, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, all, 1)

	// --- Delete ---
	err = repo.Delete(intStr)
	assert.NoError(t, err)

	deleted, err := repo.FindByID(intStr)
	assert.Error(t, err)
	assert.Nil(t, deleted)
}
