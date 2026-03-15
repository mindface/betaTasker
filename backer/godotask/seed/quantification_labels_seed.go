package seed

import (
	"encoding/csv"
	"fmt"
	"os"
	"io"
	"strconv"
	// "strings"
	"time"
	// "encoding/json"

	"github.com/godotask/seed/utils"
	"github.com/godotask/infrastructure/db/model"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SeedQuantificationLabelsFromCSV(db *gorm.DB) error {
	path := utils.GetSeedPath()
	filePath := fmt.Sprintf("seed/%s/qualitative_label.csv", path)

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open qualitative_label.csv: %v", err)
	}

	reader := csv.NewReader(file)
	defer file.Close()

	// header skip
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV header: %w", err)
	}

	var models []model.QualitativeLabel

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

		createdAt, _ := time.Parse(time.RFC3339, record[5])
		updatedAt, _ := time.Parse(time.RFC3339, record[6])

		models = append(models, model.QualitativeLabel{
			ID:        id,
			UserID:    userID,
			TaskID:    taskID,
			Content:   record[3],
			Category:  record[4],
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	// batch insert
	if len(models) > 0 {

		seedDB := db.Session(&gorm.Session{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		err := seedDB.Transaction(func(tx *gorm.DB) error {
			return tx.CreateInBatches(models, 1000).Error
		})

		if err != nil {
			return fmt.Errorf("failed to insert qualitative labels: %w", err)
		}

		fmt.Printf("Successfully seeded %d qualitative labels\n", len(models))
	}

	return nil
}