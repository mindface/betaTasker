package migration

import (
	"github.com/godotask/model"
	"gorm.io/gorm"
)

func CreateStateEvaluationTables(db *gorm.DB) error {
	// Create state_evaluations table
	err := db.AutoMigrate(&model.StateEvaluation{})
	if err != nil {
		return err
	}

	// Create tool_matching_results table
	err = db.AutoMigrate(&model.ToolMatchingResult{})
	if err != nil {
		return err
	}

	// Create process_monitoring table
	err = db.AutoMigrate(&model.ProcessMonitoring{})
	if err != nil {
		return err
	}

	// Create learning_patterns table
	err = db.AutoMigrate(&model.LearningPattern{})
	if err != nil {
		return err
	}

	// Add indexes for better performance
	err = addStateEvaluationIndexes(db)
	if err != nil {
		return err
	}

	return nil
}

func addStateEvaluationIndexes(db *gorm.DB) error {
	// Add indexes for state_evaluations table
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_state_evaluations_user_id ON state_evaluations(user_id);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_state_evaluations_task_id ON state_evaluations(task_id);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_state_evaluations_level ON state_evaluations(level);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_state_evaluations_status ON state_evaluations(status);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_state_evaluations_created_at ON state_evaluations(created_at DESC);").Error; err != nil {
		return err
	}

	// Add indexes for tool_matching_results table
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_tool_matching_state_evaluation_id ON tool_matching_results(state_evaluation_id);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_tool_matching_robot_id ON tool_matching_results(robot_id);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_tool_matching_model_id ON tool_matching_results(optimization_model_id);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_tool_matching_score ON tool_matching_results(matching_score DESC);").Error; err != nil {
		return err
	}

	// Add indexes for process_monitoring table
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_process_monitoring_state_evaluation_id ON process_monitoring(state_evaluation_id);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_process_monitoring_process_type ON process_monitoring(process_type);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_process_monitoring_status ON process_monitoring(status);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_process_monitoring_start_time ON process_monitoring(start_time DESC);").Error; err != nil {
		return err
	}

	// Add indexes for learning_patterns table
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_learning_patterns_user_id ON learning_patterns(user_id);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_learning_patterns_pattern_type ON learning_patterns(pattern_type);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_learning_patterns_domain ON learning_patterns(domain);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_learning_patterns_seci_stage ON learning_patterns(seci_stage);").Error; err != nil {
		return err
	}
	
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_learning_patterns_validated ON learning_patterns(validated);").Error; err != nil {
		return err
	}

	return nil
}