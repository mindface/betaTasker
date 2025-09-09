package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/model"
)

type QuantificationLabelRepository struct {
	db *gorm.DB
}

func NewQuantificationLabelRepository(db *gorm.DB) *QuantificationLabelRepository {
	return &QuantificationLabelRepository{
		db: db,
	}
}

// AutoMigrate - テーブル作成・マイグレーション
func (r *QuantificationLabelRepository) AutoMigrate() error {
	return r.db.AutoMigrate(
		&model.QuantificationLabel{},
		&model.ImageAnnotation{},
		&model.LabelRevision{},
		&model.LabelDataset{},
		&model.LabelRelation{},
		&model.VisualMetaphor{},
		&model.UserCalibration{},
		&model.MultimodalData{},
		// 現象的フレームワーク関連
		&model.PhenomenologicalFramework{},
		&model.KnowledgePattern{},
		&model.OptimizationModel{},
		&model.RobotArmSpecification{},
		&model.ProcessOptimization{},
		&model.TeachingFreeControl{},
		&model.LanguageOptimization{},
	)
}

// CreateIndexes - インデックス作成
func (r *QuantificationLabelRepository) CreateIndexes() error {
	// QuantificationLabel のインデックス
	queries := []string{
		`CREATE INDEX IF NOT EXISTS idx_quantification_labels_domain ON quantification_labels(domain);`,
		`CREATE INDEX IF NOT EXISTS idx_quantification_labels_category ON quantification_labels(category);`,
		`CREATE INDEX IF NOT EXISTS idx_quantification_labels_unit ON quantification_labels(unit);`,
		`CREATE INDEX IF NOT EXISTS idx_quantification_labels_validated ON quantification_labels(validated);`,
		`CREATE INDEX IF NOT EXISTS idx_quantification_labels_source ON quantification_labels(source);`,
		`CREATE INDEX IF NOT EXISTS idx_quantification_labels_abstract_level ON quantification_labels(abstract_level);`,
		`CREATE INDEX IF NOT EXISTS idx_quantification_labels_confidence ON quantification_labels(confidence);`,
		`CREATE INDEX IF NOT EXISTS idx_quantification_labels_value_unit ON quantification_labels(value, unit);`,
		`CREATE INDEX IF NOT EXISTS idx_quantification_labels_created_at ON quantification_labels(created_at);`,
		
		// 全文検索用インデックス（PostgreSQLの場合）
		`CREATE INDEX IF NOT EXISTS idx_quantification_labels_text_search ON quantification_labels USING gin(to_tsvector('english', original_text || ' ' || normalized_text));`,
		
		// ImageAnnotation のインデックス
		`CREATE INDEX IF NOT EXISTS idx_image_annotations_label_id ON image_annotations(label_id);`,
		`CREATE INDEX IF NOT EXISTS idx_image_annotations_type ON image_annotations(type);`,
		
		// LabelRevision のインデックス
		`CREATE INDEX IF NOT EXISTS idx_label_revisions_label_id ON label_revisions(label_id);`,
		`CREATE INDEX IF NOT EXISTS idx_label_revisions_timestamp ON label_revisions(timestamp);`,
		`CREATE INDEX IF NOT EXISTS idx_label_revisions_user_id ON label_revisions(user_id);`,
		
		// LabelRelation のインデックス
		`CREATE INDEX IF NOT EXISTS idx_label_relations_source_id ON label_relations(source_id);`,
		`CREATE INDEX IF NOT EXISTS idx_label_relations_target_id ON label_relations(target_id);`,
		`CREATE INDEX IF NOT EXISTS idx_label_relations_type ON label_relations(relation_type);`,
		
		// UserCalibration のインデックス
		`CREATE INDEX IF NOT EXISTS idx_user_calibrations_user_id ON user_calibrations(user_id);`,
		`CREATE INDEX IF NOT EXISTS idx_user_calibrations_reference_object ON user_calibrations(reference_object);`,
		
		// MultimodalData のインデックス
		`CREATE INDEX IF NOT EXISTS idx_multimodal_data_user_id ON multimodal_data(user_id);`,
		`CREATE INDEX IF NOT EXISTS idx_multimodal_data_task_id ON multimodal_data(task_id);`,
		`CREATE INDEX IF NOT EXISTS idx_multimodal_data_mapping_type ON multimodal_data(mapping_type);`,
		`CREATE INDEX IF NOT EXISTS idx_multimodal_data_verified ON multimodal_data(verified);`,
	}

	for _, query := range queries {
		if err := r.db.Exec(query).Error; err != nil {
			return err
		}
	}

	return nil
}

// CreateConstraints - 制約作成
func (r *QuantificationLabelRepository) CreateConstraints() error {
	constraints := []string{
		// QuantificationLabel の制約
		`ALTER TABLE quantification_labels ADD CONSTRAINT IF NOT EXISTS chk_confidence_range CHECK (confidence >= 0 AND confidence <= 1);`,
		`ALTER TABLE quantification_labels ADD CONSTRAINT IF NOT EXISTS chk_accuracy_range CHECK (accuracy >= 0 AND accuracy <= 1);`,
		`ALTER TABLE quantification_labels ADD CONSTRAINT IF NOT EXISTS chk_consistency_range CHECK (consistency >= 0 AND consistency <= 1);`,
		`ALTER TABLE quantification_labels ADD CONSTRAINT IF NOT EXISTS chk_reproducibility_range CHECK (reproducibility >= 0 AND reproducibility <= 1);`,
		`ALTER TABLE quantification_labels ADD CONSTRAINT IF NOT EXISTS chk_usability_range CHECK (usability >= 0 AND usability <= 1);`,
		`ALTER TABLE quantification_labels ADD CONSTRAINT IF NOT EXISTS chk_abstract_level CHECK (abstract_level IN ('concrete', 'semi-abstract', 'abstract'));`,
		`ALTER TABLE quantification_labels ADD CONSTRAINT IF NOT EXISTS chk_source CHECK (source IN ('manual', 'automatic', 'hybrid'));`,
		
		// ImageAnnotation の制約
		`ALTER TABLE image_annotations ADD CONSTRAINT IF NOT EXISTS chk_annotation_confidence CHECK (confidence >= 0 AND confidence <= 1);`,
		`ALTER TABLE image_annotations ADD CONSTRAINT IF NOT EXISTS chk_annotation_type CHECK (type IN ('region', 'point', 'measurement', 'text'));`,
		
		// LabelRelation の制約
		`ALTER TABLE label_relations ADD CONSTRAINT IF NOT EXISTS chk_relation_strength CHECK (strength >= 0 AND strength <= 1);`,
		`ALTER TABLE label_relations ADD CONSTRAINT IF NOT EXISTS chk_relation_confidence CHECK (confidence >= 0 AND confidence <= 1);`,
		`ALTER TABLE label_relations ADD CONSTRAINT IF NOT EXISTS chk_relation_type CHECK (relation_type IN ('synonym', 'hypernym', 'hyponym', 'meronym', 'holonym', 'similar', 'opposite'));`,
		
		// UserCalibration の制約
		`ALTER TABLE user_calibrations ADD CONSTRAINT IF NOT EXISTS chk_calibration_confidence CHECK (confidence >= 0 AND confidence <= 1);`,
		
		// MultimodalData の制約
		`ALTER TABLE multimodal_data ADD CONSTRAINT IF NOT EXISTS chk_multimodal_confidence CHECK (confidence >= 0 AND confidence <= 1);`,
		`ALTER TABLE multimodal_data ADD CONSTRAINT IF NOT EXISTS chk_multimodal_mapping_type CHECK (mapping_type IN ('direct', 'inferred', 'learned'));`,
		`ALTER TABLE multimodal_data ADD CONSTRAINT IF NOT EXISTS chk_multimodal_feedback CHECK (user_feedback IN ('correct', 'too_high', 'too_low', 'incorrect') OR user_feedback IS NULL);`,
	}

	for _, constraint := range constraints {
		if err := r.db.Exec(constraint).Error; err != nil {
			// 制約が既に存在する場合はエラーを無視
			continue
		}
	}

	return nil
}

// CreateTriggers - トリガー作成
func (r *QuantificationLabelRepository) CreateTriggers() error {
	triggers := []string{
		// QuantificationLabel の updated_at 自動更新
		`CREATE OR REPLACE FUNCTION update_quantification_label_updated_at()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.updated_at = CURRENT_TIMESTAMP;
			RETURN NEW;
		END;
		$$ language 'plpgsql';`,
		
		`DROP TRIGGER IF EXISTS update_quantification_label_updated_at_trigger ON quantification_labels;
		CREATE TRIGGER update_quantification_label_updated_at_trigger
		BEFORE UPDATE ON quantification_labels
		FOR EACH ROW EXECUTE FUNCTION update_quantification_label_updated_at();`,
		
		// バージョン自動インクリメント
		`CREATE OR REPLACE FUNCTION increment_label_version()
		RETURNS TRIGGER AS $$
		BEGIN
			IF OLD.* IS DISTINCT FROM NEW.* THEN
				NEW.version = OLD.version + 1;
			END IF;
			RETURN NEW;
		END;
		$$ language 'plpgsql';`,
		
		`DROP TRIGGER IF EXISTS increment_label_version_trigger ON quantification_labels;
		CREATE TRIGGER increment_label_version_trigger
		BEFORE UPDATE ON quantification_labels
		FOR EACH ROW EXECUTE FUNCTION increment_label_version();`,
		
		// 統計更新トリガー（ラベル作成・更新・削除時）
		`CREATE OR REPLACE FUNCTION update_label_statistics()
		RETURNS TRIGGER AS $$
		BEGIN
			-- 統計テーブルがあれば更新（実装は省略）
			RETURN COALESCE(NEW, OLD);
		END;
		$$ language 'plpgsql';`,
		
		`DROP TRIGGER IF EXISTS update_label_statistics_trigger ON quantification_labels;
		CREATE TRIGGER update_label_statistics_trigger
		AFTER INSERT OR UPDATE OR DELETE ON quantification_labels
		FOR EACH ROW EXECUTE FUNCTION update_label_statistics();`,
	}

	for _, trigger := range triggers {
		if err := r.db.Exec(trigger).Error; err != nil {
			return err
		}
	}

	return nil
}

// CreateViews - ビュー作成
func (r *QuantificationLabelRepository) CreateViews() error {
	views := []string{
		// ラベル統計ビュー
		`CREATE OR REPLACE VIEW label_statistics_view AS
		SELECT 
			COUNT(*) as total_labels,
			COUNT(CASE WHEN validated = true THEN 1 END) as verified_labels,
			AVG(confidence) as avg_confidence,
			AVG(accuracy) as avg_accuracy,
			AVG(consistency) as avg_consistency,
			AVG(reproducibility) as avg_reproducibility,
			AVG(usability) as avg_usability,
			COUNT(CASE WHEN confidence >= 0.8 THEN 1 END) as high_confidence_labels,
			COUNT(CASE WHEN confidence >= 0.5 AND confidence < 0.8 THEN 1 END) as medium_confidence_labels,
			COUNT(CASE WHEN confidence < 0.5 THEN 1 END) as low_confidence_labels
		FROM quantification_labels
		WHERE deleted_at IS NULL;`,
		
		// ドメイン別統計ビュー
		`CREATE OR REPLACE VIEW domain_statistics_view AS
		SELECT 
			domain,
			COUNT(*) as total_labels,
			COUNT(CASE WHEN validated = true THEN 1 END) as verified_labels,
			AVG(confidence) as avg_confidence,
			AVG(accuracy) as avg_accuracy,
			MIN(created_at) as first_created,
			MAX(created_at) as last_created
		FROM quantification_labels
		WHERE deleted_at IS NULL
		GROUP BY domain;`,
		
		// 最近のアクティビティビュー
		`CREATE OR REPLACE VIEW recent_activity_view AS
		SELECT 
			'label' as type,
			id as entity_id,
			original_text as title,
			'created' as action,
			created_by as user_id,
			created_at as timestamp
		FROM quantification_labels
		WHERE deleted_at IS NULL
		UNION ALL
		SELECT 
			'revision' as type,
			label_id as entity_id,
			comment as title,
			'updated' as action,
			user_id,
			timestamp
		FROM label_revisions
		ORDER BY timestamp DESC
		LIMIT 100;`,
		
		// 品質メトリクスビュー
		`CREATE OR REPLACE VIEW quality_metrics_view AS
		SELECT 
			id,
			original_text,
			domain,
			category,
			confidence,
			accuracy,
			consistency,
			reproducibility,
			usability,
			(confidence + accuracy + consistency + reproducibility + usability) / 5 as overall_quality,
			verification_count,
			validated,
			CASE 
				WHEN (confidence + accuracy + consistency + reproducibility + usability) / 5 >= 0.8 THEN 'excellent'
				WHEN (confidence + accuracy + consistency + reproducibility + usability) / 5 >= 0.6 THEN 'good'
				WHEN (confidence + accuracy + consistency + reproducibility + usability) / 5 >= 0.4 THEN 'fair'
				ELSE 'poor'
			END as quality_rating
		FROM quantification_labels
		WHERE deleted_at IS NULL;`,
	}

	for _, view := range views {
		if err := r.db.Exec(view).Error; err != nil {
			return err
		}
	}

	return nil
}

// SeedData - 初期データ投入
func (r *QuantificationLabelRepository) SeedData() error {
	// ロボットアーム開発用の視覚的メタファー
	metaphors := []model.VisualMetaphor{
		{
			ID:              "robot_gripper",
			Metaphor:        "グリッパー開口幅",
			ReferenceObject: "standard_gripper",
			Width:           12.0,
			Height:          8.0,
			MinVariability:  0.5,
			MaxVariability:  2.0,
		},
		{
			ID:              "work_envelope",
			Metaphor:        "作業領域",
			ReferenceObject: "robot_reach",
			Width:           100.0,
			Height:          100.0,
			MinVariability:  0.8,
			MaxVariability:  1.2,
		},
		{
			ID:              "joint_range",
			Metaphor:        "関節可動域",
			ReferenceObject: "joint_rotation",
			Width:           360.0,
			Height:          360.0,
			MinVariability:  0.9,
			MaxVariability:  1.0,
		},
		{
			ID:              "payload_size",
			Metaphor:        "搬送物サイズ",
			ReferenceObject: "standard_payload",
			Width:           30.0,
			Height:          30.0,
			MinVariability:  0.3,
			MaxVariability:  3.0,
		},
		{
			ID:              "palm_size",
			Metaphor:        "手のひらサイズ",
			ReferenceObject: "adult_palm",
			Width:           10.0,
			Height:          18.0,
			MinVariability:  0.8,
			MaxVariability:  1.2,
		},
	}

	for _, metaphor := range metaphors {
		var existing model.VisualMetaphor
		if err := r.db.Where("id = ?", metaphor.ID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := r.db.Create(&metaphor).Error; err != nil {
				return err
			}
		}
	}

	// ロボットアーム開発の現象学的アプローチに基づく定量化ラベル
	defaultLabels := []model.QuantificationLabel{
		// 設計・開発フェーズ
		{
			ID:              "design_precision",
			OriginalText:    "設計精度要求",
			NormalizedText:  "設計精度要求",
			Category:        "precision",
			Domain:          "robot_design",
			Value:           0.01,
			Unit:            "mm",
			MinRange:        0.005,
			MaxRange:        0.02,
			TypicalValue:    0.01,
			Precision:       3,
			Confidence:      0.95,
			AbstractLevel:   "concrete",
			Source:          "automatic",
			Notes:           "goalFn: 位置決め精度最適化",
			RelatedConcepts: model.JSON(map[string]interface{}{
				"process": "Pa_design",
				"goal": "G_precision",
				"feedback": "result_measurement",
			}),
			SemanticTags: model.JSON(map[string]interface{}{
				"tags": []string{"設計", "精度", "仕様", "CAD"},
			}),
			CreatedBy:       "system",
			UpdatedBy:       "system",
		},
		{
			ID:              "gripper_force",
			OriginalText:    "グリッパー把持力",
			NormalizedText:  "グリッパー把持力",
			Category:        "force",
			Domain:          "robot_gripper",
			Value:           50.0,
			Unit:            "N",
			MinRange:        10.0,
			MaxRange:        100.0,
			TypicalValue:    50.0,
			Precision:       0,
			Confidence:      0.85,
			AbstractLevel:   "concrete",
			Source:          "hybrid",
			Notes:           "limit: 最小把持力, max: 破損限界力",
			RelatedConcepts: model.JSON(map[string]interface{}{
				"tacit_knowledge": "熟練者の力加減",
				"formalized": "力センサーフィードバック制御",
			}),
			CreatedBy:       "system",
			UpdatedBy:       "system",
		},
		// 部品製造フェーズ
		{
			ID:              "joint_backlash",
			OriginalText:    "関節バックラッシュ",
			NormalizedText:  "関節バックラッシュ",
			Category:        "mechanical",
			Domain:          "robot_joint",
			Value:           0.05,
			Unit:            "degree",
			MinRange:        0.01,
			MaxRange:        0.1,
			TypicalValue:    0.05,
			Precision:       2,
			Confidence:      0.9,
			AbstractLevel:   "concrete",
			Source:          "automatic",
			Notes:           "SSM: 加工精度と動作精度の相関",
			RelatedConcepts: model.JSON(map[string]interface{}{
				"manufacturing": "歯車加工精度",
				"control": "バックラッシュ補正アルゴリズム",
			}),
			CreatedBy:       "system",
			UpdatedBy:       "system",
		},
		{
			ID:              "motor_torque",
			OriginalText:    "モーター定格トルク",
			NormalizedText:  "モーター定格トルク",
			Category:        "power",
			Domain:          "robot_actuator",
			Value:           10.0,
			Unit:            "Nm",
			MinRange:        5.0,
			MaxRange:        15.0,
			TypicalValue:    10.0,
			Precision:       1,
			Confidence:      0.95,
			AbstractLevel:   "concrete",
			Source:          "manual",
			Notes:           "goalFn: 負荷トルク最適化",
			CreatedBy:       "system",
			UpdatedBy:       "system",
		},
		// 組立フェーズ
		{
			ID:              "assembly_tolerance",
			OriginalText:    "組立公差累積",
			NormalizedText:  "組立公差累積",
			Category:        "tolerance",
			Domain:          "robot_assembly",
			Value:           0.2,
			Unit:            "mm",
			MinRange:        0.1,
			MaxRange:        0.3,
			TypicalValue:    0.2,
			Precision:       1,
			Confidence:      0.8,
			AbstractLevel:   "semi-abstract",
			Source:          "hybrid",
			Notes:           "SECI: 組立ノウハウの形式知化",
			RelatedConcepts: model.JSON(map[string]interface{}{
				"tacit": "熟練工の調整感覚",
				"explicit": "公差積み上げ計算",
				"feedback": "3次元測定結果",
			}),
			CreatedBy:       "system",
			UpdatedBy:       "system",
		},
		// 調整・検査フェーズ
		{
			ID:              "repeatability",
			OriginalText:    "繰り返し精度",
			NormalizedText:  "繰り返し精度",
			Category:        "performance",
			Domain:          "robot_calibration",
			Value:           0.02,
			Unit:            "mm",
			MinRange:        0.01,
			MaxRange:        0.05,
			TypicalValue:    0.02,
			Precision:       2,
			Confidence:      0.92,
			AbstractLevel:   "concrete",
			Source:          "automatic",
			Notes:           "nc(): 繰り返し測定による精度向上",
			RelatedConcepts: model.JSON(map[string]interface{}{
				"measurement": "レーザートラッカー測定",
				"optimization": "キャリブレーション最適化",
			}),
			CreatedBy:       "system",
			UpdatedBy:       "system",
		},
		{
			ID:              "cycle_time",
			OriginalText:    "サイクルタイム",
			NormalizedText:  "サイクルタイム",
			Category:        "time",
			Domain:          "robot_performance",
			Value:           2.5,
			Unit:            "second",
			MinRange:        1.5,
			MaxRange:        4.0,
			TypicalValue:    2.5,
			Precision:       1,
			Confidence:      0.88,
			AbstractLevel:   "concrete",
			Source:          "automatic",
			Notes:           "goalFn: 生産性最大化",
			CreatedBy:       "system",
			UpdatedBy:       "system",
		},
		// ティーチング不要の自律制御
		{
			ID:              "vision_accuracy",
			OriginalText:    "ビジョン認識精度",
			NormalizedText:  "ビジョン認識精度",
			Category:        "perception",
			Domain:          "robot_vision",
			Value:           1.0,
			Unit:            "mm",
			MinRange:        0.5,
			MaxRange:        2.0,
			TypicalValue:    1.0,
			Precision:       1,
			Confidence:      0.85,
			AbstractLevel:   "semi-abstract",
			Source:          "automatic",
			Notes:           "Teaching-free: 自律的物体認識",
			RelatedConcepts: model.JSON(map[string]interface{}{
				"AI": "深層学習モデル",
				"calibration": "カメラキャリブレーション",
				"lighting": "照明条件最適化",
			}),
			CreatedBy:       "system",
			UpdatedBy:       "system",
		},
		{
			ID:              "force_control",
			OriginalText:    "力制御感度",
			NormalizedText:  "力制御感度",
			Category:        "control",
			Domain:          "robot_control",
			Value:           0.1,
			Unit:            "N",
			MinRange:        0.05,
			MaxRange:        0.5,
			TypicalValue:    0.1,
			Precision:       2,
			Confidence:      0.9,
			AbstractLevel:   "semi-abstract",
			Source:          "hybrid",
			Notes:           "暗黙知→形式知: 人の力加減を数値化",
			RelatedConcepts: model.JSON(map[string]interface{}{
				"human_skill": "職人の手の感覚",
				"sensor": "6軸力覚センサー",
				"control": "インピーダンス制御",
			}),
			CreatedBy:       "system",
			UpdatedBy:       "system",
		},
		// 保守・運用フェーズ
		{
			ID:              "mtbf",
			OriginalText:    "平均故障間隔",
			NormalizedText:  "平均故障間隔",
			Category:        "reliability",
			Domain:          "robot_maintenance",
			Value:           5000.0,
			Unit:            "hour",
			MinRange:        3000.0,
			MaxRange:        8000.0,
			TypicalValue:    5000.0,
			Precision:       0,
			Confidence:      0.8,
			AbstractLevel:   "abstract",
			Source:          "automatic",
			Notes:           "予知保全: IoTセンサーによる状態監視",
			CreatedBy:       "system",
			UpdatedBy:       "system",
		},
	}

	for _, label := range defaultLabels {
		var existing model.QuantificationLabel
		if err := r.db.Where("id = ?", label.ID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := r.db.Create(&label).Error; err != nil {
				return err
			}
		}
	}

	// 現象学的フレームワークのseedデータ
	if err := r.SeedPhenomenologicalData(); err != nil {
		return err
	}

	return nil
}

// SeedPhenomenologicalData - 現象学的フレームワークのデータ投入
func (r *QuantificationLabelRepository) SeedPhenomenologicalData() error {
	// 現象学的フレームワーク
	frameworks := []model.PhenomenologicalFramework{
		{
			ID:          "robot_precision_framework",
			Name:        "ロボット精度最適化フレームワーク",
			Description: "位置決め精度を目的とした現象学的アプローチ",
			Goal:        "G: 位置決め精度±0.01mm達成",
			Scope:       "A: 6軸ロボットアームの動作範囲全体",
			Process: model.JSON(map[string]interface{}{
				"Pa": "キャリブレーション→測定→補正→検証の反復プロセス",
				"steps": []string{"初期測定", "誤差解析", "補正値計算", "適用", "再測定"},
			}),
			Result: model.JSON(map[string]interface{}{
				"achieved_accuracy": 0.012,
				"iterations": 5,
			}),
			Feedback: model.JSON(map[string]interface{}{
				"type": "continuous",
				"method": "レーザートラッカーによる実時間測定",
			}),
			LimitMin:      0.005,
			LimitMax:      0.02,
			GoalFunction:  "minimize(abs(measured_position - target_position))",
			AbstractLevel: "L1",
			Domain:        "robot_calibration",
		},
		{
			ID:          "force_control_framework",
			Name:        "力制御最適化フレームワーク",
			Description: "人の感覚を数値化した力制御",
			Goal:        "G: 破損なく確実な把持",
			Scope:       "A: 多様な形状・材質の対象物",
			Process: model.JSON(map[string]interface{}{
				"Pa": "力センサフィードバック制御",
				"tacit_to_explicit": "熟練者の把持力→センサ値マッピング",
			}),
			LimitMin:      10.0,
			LimitMax:      100.0,
			GoalFunction:  "optimize(grip_force, constraints=[no_damage, secure_grip])",
			AbstractLevel: "L2",
			Domain:        "robot_gripper",
		},
	}

	for _, framework := range frameworks {
		var existing model.PhenomenologicalFramework
		if err := r.db.Where("id = ?", framework.ID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := r.db.Create(&framework).Error; err != nil {
				return err
			}
		}
	}

	// 知識パターン（暗黙知→形式知）
	patterns := []model.KnowledgePattern{
		{
			ID:             "assembly_skill_pattern",
			Type:           "tacit",
			Domain:         "robot_assembly",
			TacitKnowledge: "熟練工の『しっくりくる』感覚",
			ExplicitForm:   "力覚センサ値: Fx<0.5N, Fy<0.5N, Tz<0.1Nm",
			ConversionPath: model.JSON(map[string]interface{}{
				"SECI": []string{"共同化", "表出化", "連結化", "内面化"},
				"method": "力覚データ記録→パターン分析→閾値設定",
			}),
			Accuracy:      0.85,
			Coverage:      0.75,
			Consistency:   0.90,
			AbstractLevel: "L1",
		},
		{
			ID:             "vision_recognition_pattern",
			Type:           "explicit",
			Domain:         "robot_vision",
			TacitKnowledge: "対象物の見分け方",
			ExplicitForm:   "深層学習モデル: ResNet-50, mAP=0.92",
			ConversionPath: model.JSON(map[string]interface{}{
				"process": "画像収集→アノテーション→学習→検証",
			}),
			Accuracy:      0.92,
			Coverage:      0.88,
			Consistency:   0.95,
			AbstractLevel: "L0",
		},
	}

	for _, pattern := range patterns {
		var existing model.KnowledgePattern
		if err := r.db.Where("id = ?", pattern.ID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := r.db.Create(&pattern).Error; err != nil {
				return err
			}
		}
	}

	// 最適化モデル
	optimizations := []model.OptimizationModel{
		{
			ID:            "trajectory_optimization",
			Name:          "軌道最適化",
			Type:          "control_theory",
			ObjectiveFunction: "minimize(time) + minimize(energy) subject to collision_free",
			Constraints: map[string]interface{}{
				"joint_limits": true,
				"collision_avoidance": true,
				"singularity_avoidance": true,
			},
			Parameters: map[string]interface{}{
				"max_velocity": 1000,
				"max_acceleration": 5000,
				"sampling_time": 0.001,
			},
			PerformanceMetric: map[string]interface{}{
				"avg_improvement": 25.5,
				"computation_time": 0.15,
			},
			IterationCount:  1000,
			ConvergenceRate: 0.95,
			Domain:          "robot_motion",
		},
	}

	for _, opt := range optimizations {
		var existing model.OptimizationModel
		if err := r.db.Where("id = ?", opt.ID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := r.db.Create(&opt).Error; err != nil {
				return err
			}
		}
	}

	// ロボットアーム仕様
	specifications := []model.RobotArmSpecification{
		{
			ID:             "teaching_free_arm_v1",
			ModelName:      "TF-ARM-001",
			DOF:            6,
			Reach:          850.0,
			Payload:        5.0,
			RepeatAccuracy: 0.02,
			MaxSpeed:       1000.0,
			WorkEnvelope: model.JSON(map[string]interface{}{
				"shape": "spherical",
				"radius": 850,
			}),
			JointLimits: model.JSON(map[string]interface{}{
				"J1": []int{-170, 170},
				"J2": []int{-90, 135},
				"J3": []int{-170, 90},
				"J4": []int{-180, 180},
				"J5": []int{-135, 135},
				"J6": []int{-360, 360},
			}),
			TeachingMethod: "vision",
			ControlSystem: model.JSON(map[string]interface{}{
				"type": "hybrid",
				"vision": "stereo_camera",
				"force": "6DOF_sensor",
				"ai": "reinforcement_learning",
			}),
			SafetyFeatures: model.JSON(map[string]interface{}{
				"collision_detection": true,
				"force_limiting": true,
				"speed_monitoring": true,
			}),
			MaintenanceSchedule: model.JSON(map[string]interface{}{
				"daily": []string{"visual_inspection"},
				"weekly": []string{"joint_lubrication"},
				"monthly": []string{"accuracy_check"},
			}),
		},
	}

	for _, spec := range specifications {
		var existing model.RobotArmSpecification
		if err := r.db.Where("id = ?", spec.ID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := r.db.Create(&spec).Error; err != nil {
				return err
			}
		}
	}

	// 言語最適化
	languageOpts := []model.LanguageOptimization{
		{
			ID:               "precision_language_opt",
			OriginalText:     "だいたいこのへん",
			OptimizedText:    "目標座標から半径5mm以内",
			Domain:           "robot_positioning",
			AbstractionLevel: "L0",
			Precision:        0.95,
			Clarity:          0.98,
			Completeness:     0.92,
			Context: model.JSON(map[string]interface{}{
				"task": "positioning",
				"tolerance": 5.0,
			}),
			Transformation: model.JSON(map[string]interface{}{
				"from": "ambiguous",
				"to": "quantified",
				"method": "context_based_inference",
			}),
			EvaluationScore: 0.95,
		},
	}

	for _, langOpt := range languageOpts {
		var existing model.LanguageOptimization
		if err := r.db.Where("id = ?", langOpt.ID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := r.db.Create(&langOpt).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// InitializeDatabase - データベース初期化
func (r *QuantificationLabelRepository) InitializeDatabase() error {
	// マイグレーション実行
	if err := r.AutoMigrate(); err != nil {
		return err
	}

	// インデックス作成
	if err := r.CreateIndexes(); err != nil {
		return err
	}

	// 制約作成
	if err := r.CreateConstraints(); err != nil {
		return err
	}

	// PostgreSQL の場合のみトリガーとビューを作成
	if r.db.Dialector.Name() == "postgres" {
		if err := r.CreateTriggers(); err != nil {
			return err
		}

		if err := r.CreateViews(); err != nil {
			return err
		}
	}

	// シードデータ投入
	if err := r.SeedData(); err != nil {
		return err
	}

	return nil
}