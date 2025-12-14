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

func SeedAssessmentsModelsFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/assessments.csv")
	if err != nil {
		return fmt.Errorf("could not open memories_models.csv: %v", err)
	}

	reader := csv.NewReader(file)
	// records, err := reader.ReadAll()
	// if err != nil {
	// 	 return fmt.Errorf("could not read CSV: %v", err)
	// }

	defer file.Close()
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	var models []model.Assessment
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
    effectiveness, _ := strconv.Atoi(record[3])
    effort, _ := strconv.Atoi(record[4])
    impact, _ := strconv.Atoi(record[5])
    noteText := record[6]

    createdAt, _ := time.Parse("2006-01-02", record[3])
    updatedAt, _ := time.Parse("2006-01-02", record[4])

		models = append(models, model.Assessment{
			ID:        id,
			UserID:    userID,
			TaskID:    taskID,
			EffectivenessScore:  effectiveness,
			EffortScore:         effort,
			ImpactScore:         impact,
			QualitativeFeedback: noteText,

			CreatedAt:  createdAt,
			UpdatedAt:  updatedAt,
		})
	}

	// バッチインサート
	if len(models) > 0 {
		if err := db.Create(&models).Error; err != nil {
			return fmt.Errorf("failed to insert assessment models: %w", err)
		}
		fmt.Printf("Successfully seeded %d assessment models\n", len(models))
	}

	log.Printf("✓ Successfully seeded %d assessment models", len(models))
	return nil
}
