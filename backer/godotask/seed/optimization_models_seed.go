package seed

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"io"
	// "encoding/json"

	"github.com/godotask/infrastructure/db/model"
	"gorm.io/gorm"
)

func SeedOptimizationModelsFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/optimization_models.csv")
	if err != nil {
		return fmt.Errorf("could not open optimization_models.csv: %v", err)
	}

	reader := csv.NewReader(file)
	// records, err := reader.ReadAll()
	// if err != nil {
	// 	return fmt.Errorf("could not read CSV: %v", err)
	// }
	defer file.Close()
	_, err = reader.Read()
	if err != nil {
			return fmt.Errorf("failed to read CSV header: %w", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read CSV: %v", err)
	}

	var models []model.OptimizationModel
	for i, record := range records {
		if i == 0 { // skip header
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read CSV record: %w", err)
		}

		// 数値フィールドの変換
		iterationCount, err := strconv.ParseFloat(record[7], 64)
		if err != nil {
			return fmt.Errorf("failed to parse iteration_count: %w", err)
		}

		convergenceRate, err := strconv.ParseFloat(record[8], 64)
		if err != nil {
			return fmt.Errorf("failed to parse convergence_rate: %w", err)
		}

		createdAt, err := time.Parse(time.RFC3339, record[11])
		if err != nil {
			log.Printf("⚠️ Invalid created_at at line %d: %v", i+1, err)
			continue
		}
		updatedAt, err := time.Parse(time.RFC3339, record[12])
		if err != nil {
			log.Printf("⚠️ Invalid updated_at at line %d: %v", i+1, err)
			continue
		}

		models = append(models, model.OptimizationModel{
			ID:                record[0],
			Name:              record[1],
			Type:              record[2],
			ObjectiveFunction: record[3],
			Constraints:       record[4],
			Parameters:        record[5],
			PerformanceMetric: record[6],
			IterationCount:    iterationCount,
			ConvergenceRate:   convergenceRate,
			Domain:            record[9],
			Application:       record[10],
			CreatedAt:         createdAt,
			UpdatedAt:         updatedAt,
		})
	}

	// バッチインサート
	if len(models) > 0 {
		if err := db.Create(&models).Error; err != nil {
			return fmt.Errorf("failed to insert optimization models: %w", err)
		}
		fmt.Printf("Successfully seeded %d optimization models\n", len(models))
	}

	log.Printf("✓ Successfully seeded %d optimization models", len(models))
	return nil
}