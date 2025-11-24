package seed

import (
	"log"
	"time"
	"encoding/json"

	"github.com/godotask/model"
	"gorm.io/gorm"
)

// SeedPhenomenologicalData - 現象学的フレームワークのシードデータ
func SeedPhenomenologicalData(db *gorm.DB) error {
	log.Println("Starting phenomenological data seeding...")

	// ロボットアーム仕様をシード
	if err := seedRobotArmSpecifications(db); err != nil {
		return err
	}

	// 最適化モデルをシード
	if err := seedOptimizationModelsFixed(db); err != nil {
		return err
	}

	log.Println("✓ Phenomenological data seeding completed")
	return nil
}

func seedOptimizationModelsFixed(db *gorm.DB) error {
		constraints1, _ := json.Marshal(map[string]interface{}{
				"cycle_time_max": 3.0,
				"accuracy_min":   0.02,
				"payload":        5.0,
		})

		constraints2, _ := json.Marshal(map[string]interface{}{
				"collision_free":        true,
				"joint_limits":          true,
				"singularity_avoidance": true,
		})

		constraints3, _ := json.Marshal(map[string]interface{}{
				"max_force":        "100N",
				"stability_margin": 0.1,
		})
		params, _ := json.Marshal(map[string]interface{}{})
		metrics, _ := json.Marshal(map[string]interface{}{})

		models := []model.OptimizationModel{
		{
        ID:                "energy_optimization",
        Name:              "エネルギー最適化",
        Type:              "ml_based",
        ObjectiveFunction: "minimize(energy_consumption) subject to performance_constraints",
        Constraints:       string(constraints1),
        Parameters:        string(params),
        PerformanceMetric: string(metrics),
        IterationCount:    0,
        ConvergenceRate:   0,
        Domain:            "robot_efficiency",
        Application:       "エネルギー効率化アプリケーション",
        CreatedAt:         time.Now(),
        UpdatedAt:         time.Now(),
    },
    {
        ID:                "trajectory_optimization",
        Name:              "軌道最適化",
        Type:              "control_theory",
        ObjectiveFunction: "minimize(time) + minimize(energy)",
        Constraints:       string(constraints2),
        Parameters:        string(params),
        PerformanceMetric: string(metrics),
        IterationCount:    0,
        ConvergenceRate:   0,
        Domain:            "robot_motion",
        Application:       "ロボット動作の最適化",
        CreatedAt:         time.Now(),
        UpdatedAt:         time.Now(),
    },
    {
        ID:                "force_optimization",
        Name:              "力制御最適化",
        Type:              "hybrid",
        ObjectiveFunction: "minimize(force_error) + minimize(position_error)",
        Constraints:       string(constraints3),
        Parameters:        string(params),
        PerformanceMetric: string(metrics),
        IterationCount:    0,
        ConvergenceRate:   0,
        Domain:            "robot_control",
        Application:       "精密力制御アプリケーション",
        CreatedAt:         time.Now(),
        UpdatedAt:         time.Now(),
    },
	}

	for _, optModel := range models {
		var existingModel model.OptimizationModel
		if err := db.Where("id = ?", optModel.ID).First(&existingModel).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&optModel).Error; err != nil {
					log.Printf("Error creating optimization model %s: %v", optModel.ID, err)
					return err
				}
			} else {
				return err
			}
		} else {
			log.Printf("Optimization model %s already exists, skipping", optModel.ID)
		}
	}

	log.Printf("✓ Successfully seeded %d optimization models", len(models))
	return nil
}

func seedRobotArmSpecifications(db *gorm.DB) error {
	// specs := []model.RobotSpecification{
	// 	{
	// 		ID:                      "teaching_free_arm_v1",
	// 		ModelName:              "教示フリーアーム V1",
	// 		DOF:                    6,
	// 		ReachMm:                850.0,
	// 		PayloadKg:              5.0,
	// 		RepeatAccuracyMm:       0.02,
	// 		MaxSpeedMmS:            1000.0,
	// 		WorkEnvelopeShape:      "spherical",
	// 		TeachingMethod:         "vision",
	// 		ControlType:            "hybrid",
	// 		VisionSystem:           nil,
	// 		ForceSensor:            nil,
	// 		SafetyFeatures:         nil,
	// 		MaintenanceIntervalHours: 2000,
	// 		CreatedAt:              time.Now(),
	// 		UpdatedAt:              time.Now(),
	// 	},
	// 	{
	// 		ID:                      "collaborative_robot_v2",
	// 		ModelName:              "協働ロボット V2",
	// 		DOF:                    7,
	// 		ReachMm:                1200.0,
	// 		PayloadKg:              10.0,
	// 		RepeatAccuracyMm:       0.01,
	// 		MaxSpeedMmS:            2000.0,
	// 		WorkEnvelopeShape:      "spherical",
	// 		TeachingMethod:         "ai",
	// 		ControlType:            "adaptive",
	// 		VisionSystem:           nil,
	// 		ForceSensor:            nil,
	// 		AICapability:           nil,
	// 		SafetyFeatures:         nil,
	// 		MaintenanceIntervalHours: 1500,
	// 		CreatedAt:              time.Now(),
	// 		UpdatedAt:              time.Now(),
	// 	},
	// }

	// for _, spec := range specs {
	// 	var existingSpec model.RobotSpecification
	// 	if err := db.Where("id = ?", spec.ID).First(&existingSpec).Error; err != nil {
	// 		if err == gorm.ErrRecordNotFound {
	// 			if err := db.Create(&spec).Error; err != nil {
	// 				log.Printf("Error creating robot specification %s: %v", spec.ID, err)
	// 				return err
	// 			}
	// 		} else {
	// 			return err
	// 		}
	// 	} else {
	// 		log.Printf("Robot specification %s already exists, skipping", spec.ID)
	// 	}
	// }

	log.Printf("✓ Successfully seeded %d robot specifications")
	return nil
}