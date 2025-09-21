package seed

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/godotask/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DataMigrationManager - データ移行管理
type DataMigrationManager struct {
	db *gorm.DB
}

// NewDataMigrationManager - 新しい移行管理インスタンス
func NewDataMigrationManager(db *gorm.DB) *DataMigrationManager {
	return &DataMigrationManager{db: db}
}

// MigrateExistingToNewFramework - 既存データを新フレームワークに移行
func (dmm *DataMigrationManager) MigrateExistingToNewFramework() error {
	log.Println("Starting migration from existing data to phenomenological framework...")

	// 1. MemoryContext → PhenomenologicalFramework
	if err := dmm.migrateMemoryContextToFramework(); err != nil {
		return fmt.Errorf("failed to migrate memory context: %v", err)
	}

	// 2. TechnicalFactor → QuantificationLabel
	if err := dmm.migrateTechnicalFactorToLabel(); err != nil {
		return fmt.Errorf("failed to migrate technical factors: %v", err)
	}

	// 3. HeuristicsPattern → KnowledgePattern
	if err := dmm.migrateHeuristicsToKnowledge(); err != nil {
		return fmt.Errorf("failed to migrate heuristics: %v", err)
	}

	// 4. Assessment → ProcessOptimization
	if err := dmm.migrateAssessmentToOptimization(); err != nil {
		return fmt.Errorf("failed to migrate assessments: %v", err)
	}

	log.Println("✓ Migration completed successfully")
	return nil
}

// 1. MemoryContext → PhenomenologicalFramework の移行
func (dmm *DataMigrationManager) migrateMemoryContextToFramework() error {
	var memoryContexts []struct {
		ID               int    `json:"id"`
		WorkTarget       string `json:"work_target"`
		ChangeFactor     string `json:"change_factor"`
		Goal             string `json:"goal"`
		ToolSpec         string `json:"tool_spec"`
		Concern          string `json:"concern"`
		Countermeasure   string `json:"countermeasure"`
		LearnedKnowledge string `json:"learned_knowledge"`
	}

	// 既存のMemoryContextデータを取得
	if err := dmm.db.Table("memory_contexts").Find(&memoryContexts).Error; err != nil {
		return err
	}

	log.Printf("Migrating %d memory contexts to phenomenological frameworks", len(memoryContexts))

	for _, mc := range memoryContexts {
		// 職務カテゴリとレベルの抽出
		domain := dmm.extractDomainFromWorkTarget(mc.WorkTarget)
		level := dmm.extractLevelFromWorkTarget(mc.WorkTarget)

		framework := model.PhenomenologicalFramework{
			ID:          fmt.Sprintf("migrated_framework_%d", mc.ID),
			Name:        dmm.generateFrameworkName(mc.WorkTarget),
			Description: fmt.Sprintf("Migrated from MemoryContext ID: %d", mc.ID),
			Goal:        fmt.Sprintf("G: %s", mc.Goal),
			Scope:       fmt.Sprintf("A: %s", dmm.extractScopeFromWorkTarget(mc.WorkTarget)),
			Process: model.JSON(map[string]interface{}{
				"Pa": mc.Countermeasure,
				"original_concern": mc.Concern,
				"tool_specification": mc.ToolSpec,
				"learned_knowledge": mc.LearnedKnowledge,
				"migration_source": "memory_context",
			}),
			Result: model.JSON(map[string]interface{}{
				"change_factor": mc.ChangeFactor,
				"status": "migrated",
			}),
			Feedback: model.JSON(map[string]interface{}{
				"type": "post_migration",
				"learned_knowledge": mc.LearnedKnowledge,
			}),
			LimitMin:      0.0,
			LimitMax:      1.0,
			GoalFunction:  dmm.generateGoalFunction(mc.Goal),
			AbstractLevel: level,
			Domain:        domain,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		// 重複チェック
		var existing model.PhenomenologicalFramework
		if err := dmm.db.Where("id = ?", framework.ID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := dmm.db.Create(&framework).Error; err != nil {
				log.Printf("Failed to create framework %s: %v", framework.ID, err)
			}
		}
	}

	return nil
}

// 2. TechnicalFactor → QuantificationLabel の移行
func (dmm *DataMigrationManager) migrateTechnicalFactorToLabel() error {
	var technicalFactors []struct {
		ID         int    `json:"id"`
		ContextID  int    `json:"context_id"`
		Factor     string `json:"factor"`
		Value      string `json:"value"`
		Unit       string `json:"unit"`
		Evaluation string `json:"evaluation"`
		Remark     string `json:"remark"`
	}

	if err := dmm.db.Table("technical_factors").Find(&technicalFactors).Error; err != nil {
		return err
	}

	log.Printf("Migrating %d technical factors to quantification labels", len(technicalFactors))

	for _, tf := range technicalFactors {
		// 数値変換の試行
		numericValue := dmm.parseNumericValue(tf.Value)
		confidence := dmm.evaluationToConfidence(tf.Evaluation)

		label := model.QuantificationLabel{
			ID:              fmt.Sprintf("migrated_label_%d", tf.ID),
			OriginalText:    tf.Factor,
			NormalizedText:  strings.ToLower(strings.TrimSpace(tf.Factor)),
			Category:        dmm.factorToCategory(tf.Factor),
			Domain:          "manufacturing",
			Value:           numericValue,
			Unit:            tf.Unit,
			MinRange:        numericValue * 0.9,
			MaxRange:        numericValue * 1.1,
			TypicalValue:    numericValue,
			Precision:       2,
			Confidence:      confidence,
			AbstractLevel:   "concrete",
			RelatedConcepts: model.JSON(map[string]interface{}{
				"original_evaluation": tf.Evaluation,
				"remark": tf.Remark,
				"migration_source": "technical_factor",
				"context_id": tf.ContextID,
			}),
			Notes:     fmt.Sprintf("Migrated from TechnicalFactor ID: %d. %s", tf.ID, tf.Remark),
			Source:    "manual",
			Validated: tf.Evaluation == "適正" || tf.Evaluation == "良好",
			Version:   1,
			CreatedBy: "migration_system",
			UpdatedBy: "migration_system",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		var existing model.QuantificationLabel
		if err := dmm.db.Where("id = ?", label.ID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := dmm.db.Create(&label).Error; err != nil {
				log.Printf("Failed to create label %s: %v", label.ID, err)
			}
		}
	}

	return nil
}

// 3. HeuristicsPattern → KnowledgePattern の移行
func (dmm *DataMigrationManager) migrateHeuristicsToKnowledge() error {
	var heuristicsPatterns []model.HeuristicsPattern

	if err := dmm.db.Find(&heuristicsPatterns).Error; err != nil {
		return err
	}

	log.Printf("Migrating %d heuristics patterns to knowledge patterns", len(heuristicsPatterns))

	for _, hp := range heuristicsPatterns {
		// HeuristicsPatternの構造を解析してKnowledgePatternに変換
		pattern := model.KnowledgePattern{
			ID:             fmt.Sprintf("migrated_pattern_%d", hp.ID),
			Type:           "hybrid", // ヒューリスティクスは通常hybrid
			Domain:         dmm.patternTypeToDomain(hp.Category),
			TacitKnowledge: dmm.extractTacitKnowledge(hp),
			ExplicitForm:   dmm.extractExplicitForm(hp),
			ConversionPath: model.JSON(map[string]interface{}{
				"original_pattern": hp.Name,
				"migration_method": "heuristics_to_knowledge",
				"accuracy": hp.Accuracy,
				"pattern_data": hp.Pattern,
			}),
			Accuracy:      hp.Accuracy,
			Coverage:      0.8, // デフォルト値
			Consistency:   hp.Accuracy * 0.9,
			AbstractLevel: "L1",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		var existing model.KnowledgePattern
		if err := dmm.db.Where("id = ?", pattern.ID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := dmm.db.Create(&pattern).Error; err != nil {
				log.Printf("Failed to create pattern %s: %v", pattern.ID, err)
			}
		}
	}

	return nil
}

// 4. Assessment → ProcessOptimization の移行
func (dmm *DataMigrationManager) migrateAssessmentToOptimization() error {
	var assessments []model.Assessment

	if err := dmm.db.Find(&assessments).Error; err != nil {
		return err
	}

	log.Printf("Migrating %d assessments to process optimizations", len(assessments))

	for _, assessment := range assessments {
		optimization := model.ProcessOptimization{
			ID:              uuid.New().String(),
			ProcessID:       fmt.Sprintf("task_%d", assessment.TaskID),
			OptimizationType: "quality",
			InitialState: model.JSON(map[string]interface{}{
				"effectiveness_score": 50, // 初期値想定
				"effort_score": 50,
				"impact_score": 50,
			}),
			OptimizedState: model.JSON(map[string]interface{}{
				"effectiveness_score": assessment.EffectivenessScore,
				"effort_score": assessment.EffortScore,
				"impact_score": assessment.ImpactScore,
			}),
			Improvement:     dmm.calculateImprovement(assessment),
			Method:          "manual_assessment_migration",
			Iterations:      1,
			ConvergenceTime: 0.0,
			ValidatedBy:     fmt.Sprintf("user_%d", assessment.UserID),
			ValidationDate:  assessment.CreatedAt,
			CreatedAt:       assessment.CreatedAt,
			UpdatedAt:       assessment.UpdatedAt,
		}

		if err := dmm.db.Create(&optimization).Error; err != nil {
			log.Printf("Failed to create optimization %s: %v", optimization.ID, err)
		}
	}

	return nil
}

// Helper Functions

func (dmm *DataMigrationManager) extractDomainFromWorkTarget(workTarget string) string {
	if strings.Contains(workTarget, "切削") {
		return "machining"
	} else if strings.Contains(workTarget, "組立") {
		return "assembly"
	} else if strings.Contains(workTarget, "検査") {
		return "inspection"
	}
	return "manufacturing"
}

func (dmm *DataMigrationManager) extractLevelFromWorkTarget(workTarget string) string {
	if strings.Contains(workTarget, "L1") {
		return "L1"
	} else if strings.Contains(workTarget, "L2") {
		return "L2"
	} else if strings.Contains(workTarget, "L3") {
		return "L3"
	} else if strings.Contains(workTarget, "L4") {
		return "L4"
	} else if strings.Contains(workTarget, "L5") {
		return "L5"
	}
	return "L1"
}

func (dmm *DataMigrationManager) extractScopeFromWorkTarget(workTarget string) string {
	// 角括弧内の内容を抽出
	start := strings.Index(workTarget, "[")
	end := strings.Index(workTarget, "]")
	if start != -1 && end != -1 && end > start {
		return workTarget[start+1 : end]
	}
	return workTarget
}

func (dmm *DataMigrationManager) generateFrameworkName(workTarget string) string {
	parts := strings.Split(workTarget, ":")
	if len(parts) > 1 {
		return strings.TrimSpace(parts[len(parts)-1]) + " Framework"
	}
	return "Manufacturing Process Framework"
}

func (dmm *DataMigrationManager) generateGoalFunction(goal string) string {
	if strings.Contains(goal, "精度") {
		return "minimize(deviation_from_target)"
	} else if strings.Contains(goal, "不良率") {
		return "minimize(defect_rate)"
	} else if strings.Contains(goal, "時間") {
		return "minimize(cycle_time)"
	}
	return "optimize(quality_metrics)"
}

func (dmm *DataMigrationManager) parseNumericValue(value string) float64 {
	// 数値部分のみを抽出（簡易実装）
	var result float64 = 0.0
	_, err := fmt.Sscanf(value, "%f", &result)
	if err != nil {
		// 数値変換失敗時のデフォルト値
		return 1.0
	}
	return result
}

func (dmm *DataMigrationManager) evaluationToConfidence(evaluation string) float64 {
	switch evaluation {
	case "適正", "良好":
		return 0.9
	case "やや良好":
		return 0.75
	case "普通":
		return 0.6
	case "やや不良":
		return 0.4
	case "不良":
		return 0.2
	default:
		return 0.5
	}
}

func (dmm *DataMigrationManager) factorToCategory(factor string) string {
	factor = strings.ToLower(factor)
	if strings.Contains(factor, "速度") {
		return "speed"
	} else if strings.Contains(factor, "送り") {
		return "feed"
	} else if strings.Contains(factor, "温度") {
		return "temperature"
	} else if strings.Contains(factor, "圧力") {
		return "pressure"
	} else if strings.Contains(factor, "寸法") {
		return "dimension"
	}
	return "parameter"
}

func (dmm *DataMigrationManager) patternTypeToDomain(category string) string {
	switch category {
	case "performance", "efficiency":
		return "performance_optimization"
	case "behavior", "user_interaction":
		return "user_behavior"
	case "learning", "adaptation":
		return "learning_system"
	default:
		return "general"
	}
}

func (dmm *DataMigrationManager) extractTacitKnowledge(hp model.HeuristicsPattern) string {
	// パターンデータから暗黙知を推定
	return fmt.Sprintf("ユーザーの%sパターンから得られる直感的判断", hp.Category)
}

func (dmm *DataMigrationManager) extractExplicitForm(hp model.HeuristicsPattern) string {
	// パターンデータを形式知として表現
	return fmt.Sprintf("パターン分析結果: %s, 精度: %.2f", hp.Pattern, hp.Accuracy)
}

func (dmm *DataMigrationManager) calculateImprovement(assessment model.Assessment) float64 {
	// スコアの平均値から改善率を計算（簡易実装）
	avgScore := float64(assessment.EffectivenessScore+assessment.EffortScore+assessment.ImpactScore) / 3.0
	baseScore := 50.0 // 基準値
	
	if avgScore > baseScore {
		return ((avgScore - baseScore) / baseScore) * 100.0
	}
	return 0.0
}

// ValidateMigration - 移行データの検証
func (dmm *DataMigrationManager) ValidateMigration() error {
	log.Println("Validating migration results...")

	// 各テーブルのレコード数チェック
	var counts struct {
		OriginalMemoryContexts   int64
		MigratedFrameworks       int64
		OriginalTechnicalFactors int64
		MigratedLabels          int64
		OriginalHeuristics      int64
		MigratedPatterns        int64
	}

	dmm.db.Table("memory_contexts").Count(&counts.OriginalMemoryContexts)
	dmm.db.Model(&model.PhenomenologicalFramework{}).Where("description LIKE ?", "%Migrated from MemoryContext%").Count(&counts.MigratedFrameworks)
	
	dmm.db.Table("technical_factors").Count(&counts.OriginalTechnicalFactors)
	dmm.db.Model(&model.QuantificationLabel{}).Where("notes LIKE ?", "%Migrated from TechnicalFactor%").Count(&counts.MigratedLabels)
	
	dmm.db.Model(&model.HeuristicsPattern{}).Count(&counts.OriginalHeuristics)
	dmm.db.Model(&model.KnowledgePattern{}).Where("id LIKE ?", "migrated_pattern_%").Count(&counts.MigratedPatterns)

	log.Printf("Migration validation results:")
	log.Printf("  MemoryContexts: %d → Frameworks: %d", counts.OriginalMemoryContexts, counts.MigratedFrameworks)
	log.Printf("  TechnicalFactors: %d → Labels: %d", counts.OriginalTechnicalFactors, counts.MigratedLabels)
	log.Printf("  HeuristicsPatterns: %d → KnowledgePatterns: %d", counts.OriginalHeuristics, counts.MigratedPatterns)

	return nil
}