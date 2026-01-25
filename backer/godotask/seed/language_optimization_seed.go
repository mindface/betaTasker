package seed

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"encoding/json"
	// "encoding/json"

	"github.com/godotask/infrastructure/db/model"
	"gorm.io/gorm"
)

func SeedLanguageOptimizationFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/language_optimization.csv")
	if err != nil {
		return fmt.Errorf("could not open language_optimization.csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // allow variable columns
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read CSV: %v", err)
	}

	var count int

	for i, record := range records {
		if i == 0 { // skip header
			continue
		}

		id := record[0]
		taskID, _ := strconv.Atoi(record[1])
		precision, _ := strconv.ParseFloat(record[6], 64)
		clarity, _ := strconv.ParseFloat(record[7], 64)
		completeness, _ := strconv.ParseFloat(record[8], 64)
		evalScore, _ := strconv.ParseFloat(record[11], 64)

		createdAt, err := time.Parse(time.RFC3339, record[12])
		if err != nil {
			log.Printf("⚠️ Invalid created_at at line %d: %v", i+1, err)
			continue
		}
		updatedAt, err := time.Parse(time.RFC3339, record[13])
		if err != nil {
			log.Printf("⚠️ Invalid updated_at at line %d: %v", i+1, err)
			continue
		}

		// JSON デコード
		var context model.JSON
		if err := json.Unmarshal([]byte(record[9]), &context); err != nil {
				log.Printf("Invalid context JSON at line %d: %v", i+2, err)
				continue
		}

		var transformation model.JSON
		if err := json.Unmarshal([]byte(record[10]), &transformation); err != nil {
				log.Printf("Invalid transformation JSON at line %d: %v", i+2, err)
				continue
		}

		// build the record
		data := model.LanguageOptimization{
			ID:               id,
			TaskID:           taskID,
			OriginalText:     record[2],
			OptimizedText:    record[3],
			Domain:           record[4],
			AbstractionLevel: record[5],
			Precision:        precision,
			Clarity:          clarity,
			Completeness:     completeness,
			Context:          context,
			Transformation:   transformation,
			EvaluationScore:  evalScore,
			CreatedAt:        createdAt,
			UpdatedAt:        updatedAt,
		}

		if err := db.Create(&data).Error; err != nil {
			log.Printf("❌ Failed to insert record at line %d: %v", i+1, err)
		} else {
			log.Printf("✅ Inserted record #%s successfully", id)
		}
	}

	log.Printf("✓ Successfully seeded %d qualitative labels", count)
	return nil
}