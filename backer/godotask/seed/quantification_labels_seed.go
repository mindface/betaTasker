package seed

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	// "strings"
	"time"
	// "encoding/json"

	"github.com/godotask/infrastructure/db/model"
	"gorm.io/gorm"
)

func SeedQuantificationLabelsFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/quantification_labels.csv")
	if err != nil {
		return fmt.Errorf("could not open quantification_labels.csv: %v", err)
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
		if i == 0 { // ヘッダー行スキップ
			continue
		}

		if len(record) < 22 {
			log.Printf("Skipping incomplete record at line %d", i+1)
			continue
		}

		// ---- フィールド変換 ----
		userID, _ := strconv.Atoi(record[1])
		taskID, _ := strconv.Atoi(record[2])

		// value, _ := strconv.ParseFloat(record[7], 64)
		// minRange, _ := strconv.ParseFloat(record[9], 64)
		// maxRange, _ := strconv.ParseFloat(record[10], 64)
		// typicalValue, _ := strconv.ParseFloat(record[11], 64)
		// precision, _ := strconv.Atoi(record[12])
		// confidence, _ := strconv.ParseFloat(record[13], 64)

		// validated := (strings.ToLower(record[16]) == "true" || record[16] == "1")

		createdAt, _ := time.Parse("2006-01-02 15:04:05", record[20])
		updatedAt, _ := time.Parse("2006-01-02 15:04:05", record[21])

		// 空 JSON（CSV に存在しないため）
		// emptyJSON := model.JSON(map[string]interface{}{})

		label := model.QuantificationLabel{
			ID:              record[0],
			UserID:          userID,
			TaskID:          taskID,

			OriginalText:    record[3],
			NormalizedText:  record[4],
			Category:        record[5],
			Domain:          record[6],

			// Value:           value,
			// Unit:            record[8],
			// MinRange:        minRange,
			// MaxRange:        maxRange,
			// TypicalValue:    typicalValue,
			// Precision:       precision,
			// Confidence:      confidence,
			// AbstractLevel:   record[14],

			// Source:          record[15],
			// Validated:       validated,

			// RelatedConcepts: emptyJSON,
			// SemanticTags:    emptyJSON,
			// Tags:            emptyJSON,
			// Notes:           record[17],

			// CreatedBy:       record[18],
			// UpdatedBy:       record[19],
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}

		// ---- 既存チェック ----
		var existing model.QuantificationLabel
		result := tx.Where("id = ?", label.ID).First(&existing)

		if result.RowsAffected == 0 {
			if err := tx.Create(&label).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to insert quantification label %s: %w", label.ID, err)
			}
			count++
		} else if result.Error != nil {
			tx.Rollback()
			return fmt.Errorf("failed to query quantification label %s: %w", label.ID, result.Error)
		} else {
			log.Printf("Quantification label '%s' already exists, skipping", label.ID)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Printf("✓ Successfully seeded %d quantification labels", count)
	return nil
}