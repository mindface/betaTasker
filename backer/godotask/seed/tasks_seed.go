package seed

import (
	"log"
	"fmt"
	"strconv"
	"time"
	"encoding/csv"
	"io"
	"os"

	"github.com/godotask/model"
	"gorm.io/gorm"
)


func SeedTaskModelsFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/task.csv")
	if err != nil {
		return fmt.Errorf("could not open task_models.csv: %v", err)
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

	var models []model.Task
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
    priority, _ := strconv.Atoi(record[5])

    // Date（NULL もありうるのでポインタ）
    var datePtr *time.Time
    if record[3] != "" {
        t, _ := time.Parse("2006-01-02", record[3])
        datePtr = &t
    }

    createdAt, _ := time.Parse("2006-01-02", record[6])
    updatedAt, _ := time.Parse("2006-01-02", record[7])

		models = append(models, model.Task{
			ID:        id,
			UserID:    userID,
			Title:     record[3],
			Description:   record[4],
			Date:      datePtr,
			Status:    record[6],
			Priority:  priority,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
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
