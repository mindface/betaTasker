package seed

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/godotask/model"
	"gorm.io/gorm"
)

// DataAccumulator - データ蓄積管理
type DataAccumulator struct {
	db         *gorm.DB
	backupDir  string
	importDir  string
	exportDir  string
}

// NewDataAccumulator - 新しいデータ蓄積管理インスタンス
func NewDataAccumulator(db *gorm.DB) *DataAccumulator {
	return &DataAccumulator{
		db:        db,
		backupDir: "seed/backup",
		importDir: "seed/import",
		exportDir: "seed/export",
	}
}

// AccumulationStrategy - データ蓄積戦略
type AccumulationStrategy struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"` // incremental, snapshot, differential
	Schedule    string    `json:"schedule"` // cron expression
	LastRun     time.Time `json:"last_run"`
	NextRun     time.Time `json:"next_run"`
	Status      string    `json:"status"`
	Config      map[string]interface{} `json:"config"`
}

// 1. 実運用データの収集と蓄積
func (da *DataAccumulator) CollectProductionData() error {
	log.Println("Collecting production data...")
	
	// 実運用で生成された優良データの抽出
	var labels []model.QuantificationLabel
	err := da.db.Where("confidence > ? AND validated = ?", 0.8, true).
		Order("created_at DESC").
		Limit(100).
		Find(&labels).Error
	
	if err != nil {
		return err
	}
	
	// 高品質データをseedデータとして保存
	return da.SaveAsHighQualitySeed(labels)
}

// 2. 学習データの自動生成と蓄積
func (da *DataAccumulator) GenerateLearningData() error {
	// 既存パターンから新しいバリエーションを生成
	var patterns []model.KnowledgePattern
	err := da.db.Find(&patterns).Error
	if err != nil {
		return err
	}
	
	newPatterns := []model.KnowledgePattern{}
	for _, pattern := range patterns {
		// バリエーション生成ロジック
		variations := da.generatePatternVariations(pattern)
		newPatterns = append(newPatterns, variations...)
	}
	
	// 生成データの保存
	for _, np := range newPatterns {
		var existing model.KnowledgePattern
		if err := da.db.Where("id = ?", np.ID).First(&existing).Error; err == gorm.ErrRecordNotFound {
			da.db.Create(&np)
		}
	}
	
	return nil
}

// 3. 差分バックアップとバージョン管理
func (da *DataAccumulator) CreateDifferentialBackup() error {
	timestamp := time.Now().Format("20060102_150405")
	backupPath := filepath.Join(da.backupDir, timestamp)
	
	if err := os.MkdirAll(backupPath, 0755); err != nil {
		return err
	}
	
	// 最後のバックアップ以降の変更データを抽出
	lastBackup := da.getLastBackupTime()
	
	// 各テーブルの差分データをエクスポート
	tables := []struct {
		name  string
		model interface{}
	}{
		{"quantification_labels", &model.QuantificationLabel{}},
		{"phenomenological_frameworks", &model.PhenomenologicalFramework{}},
		{"knowledge_patterns", &model.KnowledgePattern{}},
		{"optimization_models", &model.OptimizationModel{}},
	}
	
	for _, table := range tables {
		query := da.db.Model(table.model).Where("updated_at > ?", lastBackup)
		
		var results []map[string]interface{}
		query.Find(&results)
		
		if len(results) > 0 {
			filename := filepath.Join(backupPath, fmt.Sprintf("%s_diff.json", table.name))
			data, _ := json.MarshalIndent(results, "", "  ")
			if err := os.WriteFile(filename, data, 0644); err != nil {
				log.Printf("Failed to backup %s: %v", table.name, err)
			}
		}
	}
	
	// バックアップメタデータの保存
	metadata := map[string]interface{}{
		"timestamp": timestamp,
		"type":      "differential",
		"from":      lastBackup,
		"to":        time.Now(),
	}
	
	metaFile := filepath.Join(backupPath, "metadata.json")
	data, _ := json.MarshalIndent(metadata, "", "  ")
	return os.WriteFile(metaFile, data, 0644)
}

// 4. インポート/エクスポート機能
func (da *DataAccumulator) ImportFromCSV(filename string, modelType string) error {
	file, err := os.Open(filepath.Join(da.importDir, filename))
	if err != nil {
		return err
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return err
	}
	
	var count int
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		
		// CSVデータをモデルに変換
		data := make(map[string]interface{})
		for i, header := range headers {
			if i < len(record) {
				data[header] = record[i]
			}
		}
		
		// モデルタイプに応じて保存
		switch modelType {
		case "quantification_label":
			label := da.csvToQuantificationLabel(data)
			da.db.Create(&label)
			count++
		case "knowledge_pattern":
			pattern := da.csvToKnowledgePattern(data)
			da.db.Create(&pattern)
			count++
		}
	}
	
	log.Printf("Imported %d records from %s", count, filename)
	return nil
}

// 5. データ品質向上のための自動精製
func (da *DataAccumulator) RefineData() error {
	// 低品質データの識別と改善
	var lowQualityLabels []model.QuantificationLabel
	da.db.Where("confidence < ?", 0.5).Find(&lowQualityLabels)
	
	for _, label := range lowQualityLabels {
		// 類似の高品質データから改善提案を生成
		improved := da.improveLabel(label)
		if improved != nil {
			da.db.Model(&label).Updates(improved)
		}
	}
	
	return nil
}

// 6. データマイグレーション戦略
func (da *DataAccumulator) MigrateData(fromVersion, toVersion string) error {
	log.Printf("Migrating data from %s to %s", fromVersion, toVersion)
	
	migrations := map[string]func() error{
		"1.0.0->1.1.0": da.migrateV100ToV110,
		"1.1.0->1.2.0": da.migrateV110ToV120,
	}
	
	key := fmt.Sprintf("%s->%s", fromVersion, toVersion)
	if migration, exists := migrations[key]; exists {
		return migration()
	}
	
	return fmt.Errorf("no migration path from %s to %s", fromVersion, toVersion)
}

// Helper functions

func (da *DataAccumulator) SaveAsHighQualitySeed(labels []model.QuantificationLabel) error {
	timestamp := time.Now().Format("20060102")
	filename := filepath.Join(da.exportDir, fmt.Sprintf("high_quality_seeds_%s.json", timestamp))
	
	data, err := json.MarshalIndent(labels, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(filename, data, 0644)
}

func (da *DataAccumulator) generatePatternVariations(pattern model.KnowledgePattern) []model.KnowledgePattern {
	variations := []model.KnowledgePattern{}
	
	// ドメイン横展開
	domains := []string{"robot_assembly", "robot_inspection", "robot_welding"}
	for _, domain := range domains {
		if domain != pattern.Domain {
			variation := pattern
			variation.ID = fmt.Sprintf("%s_%s", pattern.ID, domain)
			variation.Domain = domain
			variation.Accuracy *= 0.9 // 転移学習のため精度は若干低下
			variations = append(variations, variation)
		}
	}
	
	return variations
}

func (da *DataAccumulator) getLastBackupTime() time.Time {
	// 最後のバックアップ時刻を取得
	var lastBackup time.Time
	
	entries, err := os.ReadDir(da.backupDir)
	if err != nil {
		return time.Now().AddDate(0, -1, 0) // デフォルト1ヶ月前
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			// ディレクトリ名から時刻を解析
			if t, err := time.Parse("20060102_150405", entry.Name()); err == nil {
				if t.After(lastBackup) {
					lastBackup = t
				}
			}
		}
	}
	
	if lastBackup.IsZero() {
		return time.Now().AddDate(0, -1, 0)
	}
	
	return lastBackup
}

func (da *DataAccumulator) csvToQuantificationLabel(data map[string]interface{}) model.QuantificationLabel {
	// CSV データを QuantificationLabel モデルに変換
	label := model.QuantificationLabel{}
	
	if id, ok := data["id"].(string); ok {
		label.ID = id
	}
	if text, ok := data["original_text"].(string); ok {
		label.OriginalText = text
	}
	// ... 他のフィールドも同様に変換
	
	return label
}

func (da *DataAccumulator) csvToKnowledgePattern(data map[string]interface{}) model.KnowledgePattern {
	// CSV データを KnowledgePattern モデルに変換
	pattern := model.KnowledgePattern{}
	
	if id, ok := data["id"].(string); ok {
		pattern.ID = id
	}
	// ... 他のフィールドも同様に変換
	
	return pattern
}

func (da *DataAccumulator) improveLabel(label model.QuantificationLabel) map[string]interface{} {
	// 類似の高品質データから改善案を生成
	var similarLabels []model.QuantificationLabel
	da.db.Where("domain = ? AND confidence > ?", label.Domain, 0.8).
		Limit(5).
		Find(&similarLabels)
	
	if len(similarLabels) == 0 {
		return nil
	}
	
	// 改善案の生成
	improvements := map[string]interface{}{}
	
	// 信頼度の平均値で更新
	var avgConfidence float64
	for _, sl := range similarLabels {
		avgConfidence += sl.Confidence
	}
	improvements["confidence"] = avgConfidence / float64(len(similarLabels))
	
	return improvements
}

func (da *DataAccumulator) migrateV100ToV110() error {
	// v1.0.0 から v1.1.0 へのマイグレーション
	log.Println("Migrating from v1.0.0 to v1.1.0...")
	
	// 新しいフィールドの追加など
	return da.db.Exec(`
		ALTER TABLE quantification_labels 
		ADD COLUMN IF NOT EXISTS migration_version VARCHAR(20) DEFAULT '1.1.0'
	`).Error
}

func (da *DataAccumulator) migrateV110ToV120() error {
	// v1.1.0 から v1.2.0 へのマイグレーション
	log.Println("Migrating from v1.1.0 to v1.2.0...")
	
	return nil
}

// ScheduledAccumulation - 定期的なデータ蓄積
func (da *DataAccumulator) ScheduledAccumulation() error {
	strategies := []AccumulationStrategy{
		{
			ID:       "daily_backup",
			Name:     "日次バックアップ",
			Type:     "snapshot",
			Schedule: "0 2 * * *", // 毎日午前2時
		},
		{
			ID:       "weekly_refinement",
			Name:     "週次データ精製",
			Type:     "incremental",
			Schedule: "0 3 * * 0", // 毎週日曜午前3時
		},
		{
			ID:       "monthly_export",
			Name:     "月次エクスポート",
			Type:     "differential",
			Schedule: "0 4 1 * *", // 毎月1日午前4時
		},
	}
	
	for _, strategy := range strategies {
		log.Printf("Executing strategy: %s", strategy.Name)
		
		switch strategy.Type {
		case "snapshot":
			if err := da.CreateDifferentialBackup(); err != nil {
				log.Printf("Backup failed: %v", err)
			}
		case "incremental":
			if err := da.RefineData(); err != nil {
				log.Printf("Refinement failed: %v", err)
			}
		case "differential":
			if err := da.CollectProductionData(); err != nil {
				log.Printf("Collection failed: %v", err)
			}
		}
		
		// 実行記録の更新
		strategy.LastRun = time.Now()
		strategy.Status = "completed"
	}
	
	return nil
}