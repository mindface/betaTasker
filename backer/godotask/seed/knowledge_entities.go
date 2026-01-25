package seed

import (
    "encoding/csv"
    "os"
    "time"
    "log"
    "strconv"

    "github.com/godotask/infrastructure/db/model"
    "gorm.io/gorm"
)

// SeedKnowledgeEntities - 各CSVデータと紐付けたKnowledgeEntityをシード
func SeedKnowledgeEntities(db *gorm.DB) error {
    // 紐付けるCSVファイルのリスト
    csvFiles := []struct {
        Path        string
        EntityType  string
        ReferenceIDCol int
        DomainCol   int
    }{
        {"seed/data/heuristics_analysis.csv", "heuristics_analysis", 0, 1},
        {"seed/data/heuristics_insight.csv", "heuristics_insight", 0, 1},
        {"seed/data/memory_context.csv", "memory_context", 1, 4},
        {"seed/data/optimization_models.csv", "optimization_model", 0, 1},
        {"seed/data/phenomenological_frameworks.csv", "phenomenological_framework", 0, 1},
        {"seed/data/quantification_labels.csv", "quantification_label", 0, 1},
    }

    for _, file := range csvFiles {
        f, err := os.Open(file.Path)
        if err != nil {
			log.Printf("Failed to open %s: %v", file.Path, err)
            continue
        }
        defer f.Close()

        reader := csv.NewReader(f)
        records, err := reader.ReadAll()
        if err != nil {
            log.Printf("Failed to read %s: %v", file.Path, err)
            continue
        }

        // ヘッダーをスキップ
        for i, record := range records {
            if i == 0 {
                continue
            }
            // 必要なカラムを抽出
            referenceID := record[file.ReferenceIDCol]
            domain := record[file.DomainCol]
            taskID := uint(1)
            if len(record) > 1 {
                if tid, err := strconv.Atoi(record[1]); err == nil {
                    taskID = uint(tid)
                }
            }
            now := time.Now()
            entity := model.KnowledgeEntity{
                ID:              file.EntityType + "_" + referenceID,
                TaskID:          taskID,
                EntityType:      file.EntityType,
                ReferenceID:     referenceID,
                Domain:          domain,
                AbstractLevel:   "auto",
                Source:          "seed",
                Tags:            model.JSON(map[string]interface{}{"auto": true}),
                LinkedEntityIDs: model.JSON(map[string]interface{}{}),
                CreatedAt:       now,
                UpdatedAt:       now,
            }
            if err := db.Create(&entity).Error; err != nil {
                log.Printf("Error inserting KnowledgeEntity %s: %v", entity.ID, err)
            }
        }
    }
    return nil
}