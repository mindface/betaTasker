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

func SeedPhenomenologicalFrameworksFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/phenomenological_frameworks.csv")
	if err != nil {
		return fmt.Errorf("could not open phenomenological_frameworks.csv: %v", err)
	}

	reader := csv.NewReader(file)
	// ヘッダー行をスキップ
	if _, err := reader.Read(); err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}
	defer file.Close()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read CSV: %v", err)
	}

	var count int
	now := time.Now()

	for i, record := range records {
		if i == 0 { // skip header
			continue
		}
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to read CSV record: %w", err)
		}

		// レコードの長さをチェック (12フィールド)
		if len(record) < 12 {
			tx.Rollback()
			return fmt.Errorf("invalid CSV record: expected 12 fields, got %d", len(record))
		}

		// limit_minをfloat64に変換
		limitMin, err := strconv.ParseFloat(record[7], 64)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to parse limit_min for %s: %w", record[0], err)
		}

		// limit_maxをfloat64に変換
		limitMax, err := strconv.ParseFloat(record[8], 64)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to parse limit_max for %s: %w", record[0], err)
		}

		// Process情報をJSON形式で構造化
		processData := map[string]interface{}{
			"type":        record[5], // process_type
			"description": record[2], // description
		}

		// Feedback情報をJSON形式で構造化
		feedbackData := map[string]interface{}{
			"type":      record[6], // feedback_type
			"limit_min": limitMin,
			"limit_max": limitMax,
		}

		// Result用の空のJSON（必要に応じて拡張）
		resultData := map[string]interface{}{}

		framework := model.PhenomenologicalFramework{
			ID:            record[0],
			TaskID:        1, // デフォルト値、必要に応じて設定
			Name:          record[1],
			Description:   record[2],
			Goal:          record[3],
			Scope:         record[4],
			Process:       model.JSON(processData),
			Result:        model.JSON(resultData),
			Feedback:      model.JSON(feedbackData),
			LimitMin:      limitMin,
			LimitMax:      limitMax,
			GoalFunction:  record[9],
			AbstractLevel: record[10],
			Domain:        record[11],
			CreatedAt:     now,
			UpdatedAt:     now,
		}

		// 既存データをチェック
		var existing model.PhenomenologicalFramework
		result := tx.Where("id = ?", record[0]).First(&existing)

		if result.Error == gorm.ErrRecordNotFound {
			// 新規作成
			if err := tx.Create(&framework).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to insert framework %s: %w", record[0], err)
			}
			count++
		} else if result.Error == nil {
			// 既存データがある場合はスキップ
			log.Printf("Framework %s already exists, skipping", record[0])
		} else {
			tx.Rollback()
			return fmt.Errorf("failed to query framework %s: %w", record[0], result.Error)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Printf("✓ Successfully seeded %d phenomenological frameworks", count)
	return nil
}