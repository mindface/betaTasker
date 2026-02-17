package seed

import (
	"encoding/json"
	"fmt"
	"time"
	"io"
	"encoding/csv"
	"strconv"
	"os"

	"github.com/godotask/infrastructure/db/model"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SeedHeuristics - ヒューリスティクスデータのシード
func SeedHeuristics(db *gorm.DB) error {

	// 分析データのシード
	if err := seedHeuristicsAnalysis(db); err != nil {
		return fmt.Errorf("failed to seed heuristics analysis: %v", err)
	}
	// トラッキングデータのシード
	if err := seedHeuristicsTracking(db); err != nil {
		return fmt.Errorf("failed to seed heuristics tracking: %v", err)
	}
	// インサイトデータのシード
	if err := seedHeuristicsInsights(db); err != nil {
		return fmt.Errorf("failed to seed heuristics insights: %v", err)
	}
	// パターンデータのシード
	if err := seedHeuristicsPatterns(db); err != nil {
		return fmt.Errorf("failed to seed heuristics patterns: %v", err)
	}
	// モデルデータのシード
	if err := seedHeuristicsModelers(db); err != nil {
		return fmt.Errorf("failed to seed heuristics models: %v", err)
	}

	return nil
}

func seedHeuristicsAnalysis(db *gorm.DB) error {
	file, err := os.Open("seed/data/heuristics_analysis.csv")
	if err != nil {
		fmt.Errorf("could not open heuristics_analysis.csv: %v", err)
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

	var models []model.HeuristicsAnalysis
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read CSV record: %w", err)
		}

    id, _ := strconv.Atoi(record[0])
    userID, _ := strconv.Atoi(record[1])
    taskID, _ := strconv.Atoi(record[2])
		timeSpent, _ := strconv.Atoi(record[5])
		difficulty, _ := strconv.ParseFloat(record[6], 64)
		efficiency, _ := strconv.ParseFloat(record[7], 64)
		errorCount, _ := strconv.Atoi(record[8])
		confidence, _ := strconv.ParseFloat(record[9], 64)
		score, _ := strconv.ParseFloat(record[10], 64)

		createdAt, _ := time.Parse(time.RFC3339, record[12])
		updatedAt, _ := time.Parse(time.RFC3339, record[13])

		models = append(models, model.HeuristicsAnalysis{
			ID:                id,
			UserID:            userID,
			TaskID:            taskID,
			AnalysisType:      record[3],
			Result:            record[4],
			TimeSpentMinutes:  timeSpent,
			DifficultyScore:   difficulty,
			EfficiencyScore:   efficiency,
			ErrorCount:        errorCount,
			Confidence:        confidence,
			Score:             score,
			Status:            record[11],
			CreatedAt:         createdAt,
			UpdatedAt:         updatedAt,
		})
	}

	// バッチインサート
	if len(models) > 0 {
		seedDB := db.Session(&gorm.Session{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		err := seedDB.Transaction(func(tx *gorm.DB) error {
			return tx.CreateInBatches(models, 1000).Error
		})
		if err != nil {
			return fmt.Errorf("failed to insert heuristics analysis models: %w", err)
		}
		fmt.Printf("Successfully seeded %d heuristics analysis models\n", len(models))
	}

	return nil
}

func seedHeuristicsTracking(db *gorm.DB) error {
	file, err := os.Open("seed/data/heuristics_tracking.csv")
	if err != nil {
		return fmt.Errorf("could not open heuristics_tracking.csv: %v", err)
	}

	reader := csv.NewReader(file)
	defer file.Close()
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	var trackings []model.HeuristicsTracking
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read CSV record: %w", err)
		}

		id, _ := strconv.Atoi(record[0])
		userID, _ := strconv.Atoi(record[1])
		taskID, _ := strconv.Atoi(record[2])

		var focusLevel *float64
		if record[6] != "" {
			fl, _ := strconv.ParseFloat(record[6], 64)
			focusLevel = &fl
		}

		isDistraction, _ := strconv.ParseBool(record[7])
		duration, _ := strconv.Atoi(record[9])

		timestamp, _ := time.Parse(time.RFC3339, record[8])
		createdAt, _ := time.Parse(time.RFC3339, record[10])
		updatedAt, _ := time.Parse(time.RFC3339, record[11])

		trackings = append(trackings, model.HeuristicsTracking{
			ID:            id,
			UserID:        userID,
			TaskID:        taskID,
			Action:        record[3],
			Context:       record[4],
			SessionID:     record[5],
			FocusLevel:    focusLevel,
			IsDistraction: isDistraction,
			Timestamp:     timestamp,
			Duration:      duration,
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
		})
	}

	// バッチインサート
	if len(trackings) > 0 {
		seedDB := db.Session(&gorm.Session{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		err := seedDB.Transaction(func(tx *gorm.DB) error {
			return tx.CreateInBatches(trackings, 500).Error
		})
		if err != nil {
			return fmt.Errorf("failed to insert heuristics tracking data: %w", err)
		}
		fmt.Printf("Successfully seeded %d heuristics tracking records\n", len(trackings))
	}

	return nil
}

func seedHeuristicsInsights(db *gorm.DB) error {
	file, err := os.Open("seed/data/heuristics_insights.csv")
	if err != nil {
		return fmt.Errorf("could not open heuristics_insights.csv: %v", err)
	}

	reader := csv.NewReader(file)
	defer file.Close()
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	var insights []model.HeuristicsInsight
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read CSV record: %w", err)
		}

		id, _ := strconv.Atoi(record[0])
		userID, _ := strconv.Atoi(record[1])
		taskID, _ := strconv.Atoi(record[2])

		confidence, _ := strconv.ParseFloat(record[6], 64)
		expectedImpact, _ := strconv.ParseFloat(record[10], 64)
		isActive, _ := strconv.ParseBool(record[11])

		createdAt, _ := time.Parse(time.RFC3339, record[12])
		updatedAt, _ := time.Parse(time.RFC3339, record[13])

		var sourceAnalysisID *int
		if record[8] != "" {
			sid, _ := strconv.Atoi(record[8])
			sourceAnalysisID = &sid
		}

		insights = append(insights, model.HeuristicsInsight{
			ID:               id,
			UserID:           userID,
			TaskID:           taskID,
			Type:             record[3],
			Title:            record[4],
			Description:      record[5],
			Confidence:       confidence,
			Data:             record[7],
			SourceAnalysisID: sourceAnalysisID,
			Recommendation:   record[9],
			ExpectedImpact:   expectedImpact,
			IsActive:         isActive,
			CreatedAt:        createdAt,
			UpdatedAt:        updatedAt,
		})
	}

	// バッチインサート
	if len(insights) > 0 {
		seedDB := db.Session(&gorm.Session{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		err := seedDB.Transaction(func(tx *gorm.DB) error {
			return tx.CreateInBatches(insights, 500).Error
		})
		if err != nil {
			return fmt.Errorf("failed to insert heuristics insights data: %w", err)
		}
		fmt.Printf("Successfully seeded %d heuristics insights records\n", len(insights))
	}

	return nil
}

func seedHeuristicsPatterns(db *gorm.DB) error {
	file, err := os.Open("seed/data/heuristics_patterns.csv")
	if err != nil {
		return fmt.Errorf("could not open heuristics_patterns.csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// ヘッダー読み飛ばし
	if _, err := reader.Read(); err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	var patterns []model.HeuristicsPattern

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read CSV record: %w", err)
		}

		id, _ := strconv.Atoi(record[0])
		userID, _ := strconv.Atoi(record[2])
		taskID, _ := strconv.Atoi(record[3])
		frequency, _ := strconv.Atoi(record[7])
		accuracy, _ := strconv.ParseFloat(record[8], 64)
		impactScore, _ := strconv.ParseFloat(record[9], 64)

		lastSeen, err := time.Parse(time.RFC3339, record[10])
		if err != nil {
			return fmt.Errorf("invalid last_seen: %w", err)
		}
		createdAt, err := time.Parse(time.RFC3339, record[11])
		if err != nil {
			return fmt.Errorf("invalid created_at: %w", err)
		}
		updatedAt, err := time.Parse(time.RFC3339, record[12])
		if err != nil {
			return fmt.Errorf("invalid updated_at: %w", err)
		}

		patterns = append(patterns, model.HeuristicsPattern{
			ID:          id,
			Name:        record[1],
			UserID:      userID,
			TaskID:      taskID,
			TaskType:    record[4],
			Category:    record[5],
			Pattern:     record[6], // jsonb string
			Frequency:   frequency,
			Accuracy:    accuracy,
			ImpactScore: impactScore,
			LastSeen:    lastSeen,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}

	if len(patterns) == 0 {
		return nil
	}

	// seed 高速化 + ログ抑制
	seedDB := db.Session(&gorm.Session{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	return seedDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(patterns, 500).Error; err != nil {
			return fmt.Errorf("failed to insert heuristics patterns: %w", err)
		}
		fmt.Printf("Successfully seeded %d heuristics patterns records\n", len(patterns))
		return nil
	})
}

func seedHeuristicsModelers(db *gorm.DB) error {
	file, err := os.Open("seed/data/heuristics_models.csv")
	if err != nil {
		return fmt.Errorf("could not open heuristics_models.csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// ヘッダー読み飛ばし
	if _, err := reader.Read(); err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	var models []model.HeuristicsModeler

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read CSV record: %w", err)
		}

		id, _ := strconv.Atoi(record[0])
		userID, _ := strconv.Atoi(record[1])
		taskID, _ := strconv.Atoi(record[2])

		trainedAt, err := time.Parse(time.RFC3339, record[8])
		if err != nil {
			return fmt.Errorf("invalid trained_at: %w", err)
		}
		createdAt, err := time.Parse(time.RFC3339, record[9])
		if err != nil {
			return fmt.Errorf("invalid created_at: %w", err)
		}
		updatedAt, err := time.Parse(time.RFC3339, record[10])
		if err != nil {
			return fmt.Errorf("invalid updated_at: %w", err)
		}

		models = append(models, model.HeuristicsModeler{
			ID:          id,
			UserID:      userID,
			TaskID:      taskID,
			ModelType:   record[3],
			Version:     record[4],
			Parameters:  record[5], // jsonb
			Performance: record[6], // jsonb
			Status:      record[7],
			TrainedAt:   trainedAt,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		})
	}

	if len(models) == 0 {
		return nil
	}

	// seed 時はログ抑制 + トランザクション + バッチ
	seedDB := db.Session(&gorm.Session{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	return seedDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(models, 500).Error; err != nil {
			return fmt.Errorf("failed to insert heuristics models: %w", err)
		}
		fmt.Printf("Successfully seeded %d heuristics models records\n", len(models))
		return nil
	})
}

// Helper function to convert map to JSON string
func toJSON(data map[string]interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}