package seed

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"io"
	// "encoding/json"

	"github.com/godotask/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SeedFromCSVFiles - CSVファイルからデータをシード
func SeedFromCSVFiles(db *gorm.DB) error {
	log.Println("Loading data from CSV files...")

	// Seed robot specifications
	// if err := seedRobotSpecificationsFromCSV(db); err != nil {
	// 	log.Printf("Warning: Failed to seed robot specifications: %v", err)
	// }

	// Seed optimization models
	if err := seedOptimizationModelsFromCSV(db); err != nil {
		log.Printf("Warning: Failed to seed optimization models: %v", err)
		return fmt.Errorf("failed to seed optimization models: %w", err)
	}

	// Seed phenomenological frameworks
	if err := seedPhenomenologicalFrameworksFromCSV(db); err != nil {
		log.Printf("Warning: Failed to seed phenomenological frameworks: %v", err)
		return fmt.Errorf("failed to seed phenomenological frameworks: %w", err)
	}

	// Seed quantification labels
	if err := seedQuantificationLabelsFromCSV(db); err != nil {
		log.Printf("Warning: Failed to seed quantification labels: %v", err)
		return fmt.Errorf("failed to seed quantification labels: %w", err)
	}

	log.Println("✓ CSV data seeding completed")
	return nil
}

// func seedRobotSpecificationsFromCSV(db *gorm.DB) error {
// 	file, err := os.Open("seed/data/robot_specifications.csv")
// 	if err != nil {
// 		return fmt.Errorf("could not open robot_specifications.csv: %v", err)
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		return fmt.Errorf("could not read CSV: %v", err)
// 	}

// 	var robots []model.RobotSpecification
// 	for i, record := range records {
// 		if i == 0 { // Skip header
// 			continue
// 		}

// 		if len(record) < 15 {
// 			continue
// 		}

// 		dof, _ := strconv.Atoi(record[2])
// 		reach, _ := strconv.ParseFloat(record[3], 64)
// 		payload, _ := strconv.ParseFloat(record[4], 64)
// 		accuracy, _ := strconv.ParseFloat(record[5], 64)
// 		maxSpeed, _ := strconv.ParseFloat(record[6], 64)
// 		maintenanceInterval, _ := strconv.Atoi(record[14])

// 		robot := model.RobotSpecification{
// 			ID:                      record[0],
// 			ModelName:              record[1],
// 			DOF:                    dof,
// 			ReachMm:                reach,
// 			PayloadKg:              payload,
// 			RepeatAccuracyMm:       accuracy,
// 			MaxSpeedMmS:            maxSpeed,
// 			WorkEnvelopeShape:      record[7],
// 			TeachingMethod:         record[8],
// 			ControlType:            record[9],
// 			MaintenanceIntervalHours: maintenanceInterval,
// 		}

// 		// Handle nullable fields - set to nil for now to avoid encoding errors
// 		robot.VisionSystem = nil
// 		robot.ForceSensor = nil
// 		robot.AICapability = nil
// 		robot.SafetyFeatures = nil

// 		robots = append(robots, robot)
// 	}

// 	// Insert data with duplicate handling
// 	for _, robot := range robots {
// 		var existingRobot model.RobotSpecification
// 		if err := db.Where("id = ?", robot.ID).First(&existingRobot).Error; err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				if err := db.Create(&robot).Error; err != nil {
// 					log.Printf("Error inserting robot %s: %v", robot.ID, err)
// 				}
// 			} else {
// 				log.Printf("Error checking robot %s: %v", robot.ID, err)
// 			}
// 		} else {
// 			log.Printf("Robot %s already exists, skipping", robot.ID)
// 		}
// 	}

// 	log.Printf("✓ Successfully seeded %d robot specifications", len(robots))
// 	return nil
// }

func seedOptimizationModelsFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/optimization_models.csv")
	if err != nil {
		return fmt.Errorf("could not open optimization_models.csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// records, err := reader.ReadAll()
	// if err != nil {
	// 	return fmt.Errorf("could not read CSV: %v", err)
	// }

	var models []model.OptimizationModel
	for {
		record, err := reader.Read()
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
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
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

func seedPhenomenologicalFrameworksFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/phenomenological_frameworks.csv")
	if err != nil {
		return fmt.Errorf("could not open phenomenological_frameworks.csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	// ヘッダー行をスキップ
	if _, err := reader.Read(); err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var count int
	now := time.Now()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
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

func seedQuantificationLabelsFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/quantification_labels.csv")
	if err != nil {
		return fmt.Errorf("could not open quantification_labels.csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read CSV: %v", err)
	}

	tx := db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %v", tx.Error)
	}

	var count int

	for i, record := range records {
		if i == 0 { // skip header
			continue
		}

		if len(record) < 10 {
			log.Printf("Skipping incomplete record at line %d", i+1)
			continue
		}

		// Parse numerical fields
		value, _ := strconv.ParseFloat(record[5], 64)
		minRange, _ := strconv.ParseFloat(record[6], 64)
		maxRange, _ := strconv.ParseFloat(record[7], 64)

		// Related concepts and tags
		concepts := map[string]interface{}{
			"concepts": strings.Split(record[8], "|"),
		}
		tags := map[string]interface{}{
			"tags": strings.Split(record[9], "|"),
		}

		label := model.QuantificationLabel{
			ID:              uuid.New().String(),
			UserID:          1, // ✅ 固定
			TaskID:          1, // ✅ 固定
			OriginalText:    record[1],
			NormalizedText:  strings.ToLower(strings.TrimSpace(record[1])),
			Category:        record[2],
			Context:         record[3],
			Domain:          record[4],
			Value:           value,
			Unit:            record[4], // domainをunitとして使用
			MinRange:        minRange,
			MaxRange:        maxRange,
			TypicalValue:    value,
			Precision:       2,
			Confidence:      0.8,
			AbstractLevel:   "concrete",
			RelatedConcepts: model.JSON(concepts),
			SemanticTags:    model.JSON(tags),
			Accuracy:        0.85,
			Consistency:     0.80,
			Reproducibility: 0.75,
			Usability:       0.90,
			Source:          "csv_import",
			Validated:       true,
			Version:         1,
			CreatedBy:       "system",
			UpdatedBy:       "system",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		// 既存チェック
		var existing model.QuantificationLabel
		result := tx.Where("id = ?", label.ID).First(&existing)

		// 存在しない場合のみ作成
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

	// コミット
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Printf("✓ Successfully seeded %d quantification labels", count)
	return nil
}