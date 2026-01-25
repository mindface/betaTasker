package seed

import (
	"log"
	"fmt"
	"strconv"
	"time"
	"encoding/csv"
	"io"
	"os"

	"github.com/godotask/infrastructure/db/model"
	"gorm.io/gorm"
)


func SeedBookModelsFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/book.csv")
	if err != nil {
		return fmt.Errorf("could not open memories_models.csv: %v", err)
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

	var models []model.Book
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read CSV record: %w", err)
		}

    id, _ := strconv.Atoi(record[0])
    taskID, _ := strconv.Atoi(record[1])

    createdAt, _ := time.Parse("2006-01-02", record[8])
    updatedAt, _ := time.Parse("2006-01-02", record[9])

		models = append(models, model.Book{
			ID:        id,
			TaskID:    taskID,
			Title:     record[2],
			Name:      record[3],
			Text:      record[4],
			Disc:      record[5],
			ImgPath:   record[6],
			Status:    record[7],
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
