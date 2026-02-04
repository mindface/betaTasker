package repository_test

import (
	"testing"
	"time"
	"strconv"

	"github.com/godotask/infrastructure/db/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&model.Assessment{})
	return db
}

func TestAssessmentRepository(t *testing.T) {
	db := setupTestDB()
	repo := &repository.AssessmentRepositoryImpl{DB: db}

	// 作成
	assessment := &model.Assessment{
		TaskID:              1,
		UserID:              1,
		EffectivenessScore:  8,
		EffortScore:         5,
		ImpactScore:         7,
		QualitativeFeedback: "Good result",
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	err := repo.Create(assessment)
	assert.NoError(t, err)
	assert.NotZero(t, assessment.ID)

	// 取得
	found, err := repo.FindByID(strconv.Itoa(assessment.ID))
	assert.NoError(t, err)
	assert.Equal(t, assessment.TaskID, found.TaskID)
	assert.Equal(t, assessment.QualitativeFeedback, found.QualitativeFeedback)

	// 更新
	updated := &model.Assessment{
		EffectivenessScore: 9,
		EffortScore:        6,
		ImpactScore:        8,
		QualitativeFeedback: "Improved result",
	}
	err = repo.Update(strconv.Itoa(assessment.ID), updated)
	assert.NoError(t, err)

	// 再取得して確認
	found, _ = repo.FindByID(strconv.Itoa(assessment.ID))
	assert.Equal(t, 9, found.EffectivenessScore)
	assert.Equal(t, "Improved result", found.QualitativeFeedback)

	// 全取得
	all, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, all, 1)

	// 削除
	err = repo.Delete(strconv.Itoa(assessment.ID))
	assert.NoError(t, err)

	// 削除後確認
	deleted, err := repo.FindByID(strconv.Itoa(assessment.ID))
	assert.Nil(t, deleted)
	assert.Error(t, err)
}

