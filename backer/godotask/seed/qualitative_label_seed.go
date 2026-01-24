package seed

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	// "encoding/json"

	"github.com/godotask/infrastructure/db/model"
	"gorm.io/gorm"
)

func SeedQualitativeLabelsFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/qualitative_label.csv")
	if err != nil {
		return fmt.Errorf("could not open qualitative_labels.csv: %v", err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read CSV: %v", err)
	}
	defer file.Close()

	tx := db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %v", tx.Error)
	}

	var count int

	for i, record := range records {
		if i == 0 { // skip header
			continue
		}

    id, _ := strconv.Atoi(record[0])
    userID, _ := strconv.Atoi(record[1])
    taskID, _ := strconv.Atoi(record[2])

    createdAt, _ := time.Parse("2006-01-02", record[5])
    updatedAt, _ := time.Parse("2006-01-02", record[6])

		label := model.QualitativeLabel{
			ID:              id,
			UserID:          userID,
			TaskID:          taskID,
			Content:    		 record[3],
			Category:        record[4],
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}

		// 既存チェック
		var existing model.QualitativeLabel
		result := tx.Where("id = ?", label.ID).First(&existing)

		// 存在しない場合のみ作成
		if result.RowsAffected == 0 {
			if err := tx.Create(&label).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to insert qualitative label %s: %w", label.ID, err)
			}
			count++
		} else if result.Error != nil {
			tx.Rollback()
			return fmt.Errorf("failed to query qualitative label %s: %w", label.ID, result.Error)
		} else {
			log.Printf("qualitative label '%s' already exists, skipping", label.ID)
		}
	}

	// コミット
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Printf("✓ Successfully seeded %d qualitative labels", count)
	return nil
}