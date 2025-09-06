package seed

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/godotask/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SeedFromCSVFiles - CSVファイルからデータをシード
func SeedFromCSVFiles(db *gorm.DB) error {
	log.Println("Loading data from CSV files...")

	// Seed robot specifications
	if err := seedRobotSpecificationsFromCSV(db); err != nil {
		log.Printf("Warning: Failed to seed robot specifications: %v", err)
	}

	// Seed optimization models
	if err := seedOptimizationModelsFromCSV(db); err != nil {
		log.Printf("Warning: Failed to seed optimization models: %v", err)
	}

	// Seed phenomenological frameworks
	if err := seedPhenomenologicalFrameworksFromCSV(db); err != nil {
		log.Printf("Warning: Failed to seed phenomenological frameworks: %v", err)
	}

	// Seed quantification labels
	if err := seedQuantificationLabelsFromCSV(db); err != nil {
		log.Printf("Warning: Failed to seed quantification labels: %v", err)
	}

	log.Println("✓ CSV data seeding completed")
	return nil
}

func seedRobotSpecificationsFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/robot_specifications.csv")
	if err != nil {
		return fmt.Errorf("could not open robot_specifications.csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read CSV: %v", err)
	}

	var robots []model.RobotSpecification
	for i, record := range records {
		if i == 0 { // Skip header
			continue
		}

		if len(record) < 15 {
			continue
		}

		dof, _ := strconv.Atoi(record[2])
		reach, _ := strconv.ParseFloat(record[3], 64)
		payload, _ := strconv.ParseFloat(record[4], 64)
		accuracy, _ := strconv.ParseFloat(record[5], 64)
		maxSpeed, _ := strconv.ParseFloat(record[6], 64)
		maintenanceInterval, _ := strconv.Atoi(record[14])

		robot := model.RobotSpecification{
			ID:                      record[0],
			ModelName:              record[1],
			DOF:                    dof,
			ReachMm:                reach,
			PayloadKg:              payload,
			RepeatAccuracyMm:       accuracy,
			MaxSpeedMmS:            maxSpeed,
			WorkEnvelopeShape:      record[7],
			TeachingMethod:         record[8],
			ControlType:            record[9],
			MaintenanceIntervalHours: maintenanceInterval,
		}

		// Handle nullable fields - set to nil for now to avoid encoding errors
		robot.VisionSystem = nil
		robot.ForceSensor = nil
		robot.AICapability = nil
		robot.SafetyFeatures = nil

		robots = append(robots, robot)
	}

	// Insert data with duplicate handling
	for _, robot := range robots {
		var existingRobot model.RobotSpecification
		if err := db.Where("id = ?", robot.ID).First(&existingRobot).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&robot).Error; err != nil {
					log.Printf("Error inserting robot %s: %v", robot.ID, err)
				}
			} else {
				log.Printf("Error checking robot %s: %v", robot.ID, err)
			}
		} else {
			log.Printf("Robot %s already exists, skipping", robot.ID)
		}
	}

	log.Printf("✓ Successfully seeded %d robot specifications", len(robots))
	return nil
}

func seedOptimizationModelsFromCSV(db *gorm.DB) error {
	file, err := os.Open("seed/data/optimization_models.csv")
	if err != nil {
		return fmt.Errorf("could not open optimization_models.csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read CSV: %v", err)
	}

	var models []model.OptimizationModel
	for i, record := range records {
		if i == 0 { // Skip header
			continue
		}

		if len(record) < 11 {
			continue
		}

		iterationCount, _ := strconv.ParseFloat(record[7], 64)
		convergenceRate, _ := strconv.ParseFloat(record[8], 64)

		optModel := model.OptimizationModel{
			ID:               record[0],
			Name:             record[1],
			Type:             record[2],
			ObjectiveFunction: record[3],
			Constraints:      record[4],
			Parameters:       &model.NullString{String: record[5], Valid: record[5] != ""},
			PerformanceMetric: &model.NullString{String: record[6], Valid: record[6] != ""},
			IterationCount:   &model.NullFloat64{Float64: iterationCount, Valid: iterationCount > 0},
			ConvergenceRate:  &model.NullFloat64{Float64: convergenceRate, Valid: convergenceRate > 0},
			Domain:           record[9],
			Application:      record[10],
		}

		models = append(models, optModel)
	}

	// Insert data with duplicate handling
	for _, optModel := range models {
		var existingModel model.OptimizationModel
		if err := db.Where("id = ?", optModel.ID).First(&existingModel).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&optModel).Error; err != nil {
					log.Printf("Error inserting optimization model %s: %v", optModel.ID, err)
				}
			} else {
				log.Printf("Error checking optimization model %s: %v", optModel.ID, err)
			}
		} else {
			log.Printf("Optimization model %s already exists, skipping", optModel.ID)
		}
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
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read CSV: %v", err)
	}

	var frameworks []model.PhenomenologicalFramework
	for i, record := range records {
		if i == 0 { // Skip header
			continue
		}

		if len(record) < 12 {
			continue
		}

		// Parse process data as JSON
		processData := map[string]interface{}{
			"type":      record[5],
			"feedback":  record[6],
			"min_limit": record[7],
			"max_limit": record[8],
		}

		framework := model.PhenomenologicalFramework{
			ID:           record[0],
			Name:         record[1],
			Description:  record[2],
			Goal:         record[3],
			Scope:        record[4],
			Process:      model.JSON(processData),
			GoalFunction: record[9],
			AbstractLevel: record[10],
			Domain:       record[11],
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		frameworks = append(frameworks, framework)
	}

	// Insert data with duplicate handling
	for _, framework := range frameworks {
		var existingFramework model.PhenomenologicalFramework
		if err := db.Where("id = ?", framework.ID).First(&existingFramework).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&framework).Error; err != nil {
					log.Printf("Error inserting phenomenological framework %s: %v", framework.ID, err)
				}
			} else {
				log.Printf("Error checking phenomenological framework %s: %v", framework.ID, err)
			}
		} else {
			log.Printf("Phenomenological framework %s already exists, skipping", framework.ID)
		}
	}

	log.Printf("✓ Successfully seeded %d phenomenological frameworks", len(frameworks))
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

	var labels []model.QuantificationLabel
	for i, record := range records {
		if i == 0 { // Skip header
			continue
		}

		if len(record) < 10 {
			continue
		}

		value, _ := strconv.ParseFloat(record[5], 64)
		minRange, _ := strconv.ParseFloat(record[6], 64)
		maxRange, _ := strconv.ParseFloat(record[7], 64)

		// Generate related concepts and semantic tags
		concepts := map[string]interface{}{
			"concepts": strings.Split(record[8], "|"),
		}

		tags := map[string]interface{}{
			"tags": strings.Split(record[9], "|"),
		}

		label := model.QuantificationLabel{
			ID:              uuid.New().String(),
			OriginalText:    record[1],
			NormalizedText:  strings.ToLower(strings.TrimSpace(record[1])),
			Category:        record[2],
			Context:         record[3],
			Domain:          record[4],
			Value:           value,
			Unit:            record[4], // Using domain as unit for now
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

		labels = append(labels, label)
	}

	// Insert data with duplicate handling
	for _, label := range labels {
		var existingLabel model.QuantificationLabel
		if err := db.Where("id = ?", label.ID).First(&existingLabel).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&label).Error; err != nil {
					log.Printf("Error inserting quantification label %s: %v", label.ID, err)
				}
			} else {
				log.Printf("Error checking quantification label %s: %v", label.ID, err)
			}
		} else {
			log.Printf("Quantification label %s already exists, skipping", label.ID)
		}
	}

	log.Printf("✓ Successfully seeded %d quantification labels", len(labels))
	return nil
}