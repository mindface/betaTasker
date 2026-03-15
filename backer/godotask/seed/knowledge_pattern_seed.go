package seed

import (
  "encoding/csv"
  "encoding/json"
  "os"
  "io"
  "time"
  "log"
  "strconv"
  "fmt"

	"github.com/godotask/seed/utils"
  "github.com/godotask/infrastructure/db/model"
  "gorm.io/gorm"
)

type ConversionPath struct {
    PatternType    string   `json:"pattern_type"`
    Level          int      `json:"level"`
    Coefficient    string   `json:"coefficient"`
    TaskType       string   `json:"task_type"`
    Characteristics []string `json:"characteristics"`
    Triggers        []string `json:"triggers"`
    Outcomes        []string `json:"outcomes"`
}

// SeedKnowledgeEntities - 各CSVデータと紐付けたKnowledgeEntityをシード
func SeedKnowledgePattern(db *gorm.DB) error {
    path := utils.GetSeedPath()
    filePath := fmt.Sprintf("seed/%s/knowledge_patterns.csv", path)

    file, err := os.Open(filePath)
    // 紐付けるCSVファイルのリスト
    if err != nil {
      return fmt.Errorf("could not open knowledge_patterns.csv: %v", err)
    }

    reader := csv.NewReader(file)
    _, err = reader.Read()
    if err != nil {
        return err
    }
    var count int
    var models []model.KnowledgePattern
	  defer file.Close()

    for {
      record, err := reader.Read()
      if err == io.EOF {
        break
      }
      if err != nil {
        return fmt.Errorf("failed to read CSV record: %w", err)
      }

      id := record[0]
      if record[6] == "pattern" {
        continue
      }
      // fmt.Println("conversion_path:", record[6])

		  taskID, _ := strconv.Atoi(record[1])

      typeVal := record[2]
      domain := record[3]
      tacitKnowledge := record[4]
      explicitForm := record[5]

      accuracy, _ := strconv.ParseFloat(record[7], 64)
      coverage, _ := strconv.ParseFloat(record[8], 64)
      consistency, _ := strconv.ParseFloat(record[9], 64)

      abstractLevel := record[10]

      createdAt, _ := time.Parse("2006-01-02", record[11])
      updatedAt, _ := time.Parse("2006-01-02", record[12])

      // JSON decode (conversion_path)
      // var conversionPath model.JSON
      var conversionPath ConversionPath

      // ダミーデータ
      conversionPath = ConversionPath{
          PatternType: "work_rhythm",
          Level: 1,
          Coefficient: "理解度",
          TaskType: "machining",
          Characteristics: []string{
              "集中時間の最適化",
              "休憩タイミングの確立",
              "効率的作業フロー",
          },
          Triggers: []string{
              "タスク開始時",
              "問題発生時",
              "レビュー時",
          },
          Outcomes: []string{
              "基本的な成果達成",
              "学習進捗",
              "スキル向上",
          },
      }
      jsonBytes, err := json.Marshal(conversionPath)
      if err != nil {
          return err
      }
      var conversionPath2 model.JSON
      err = json.Unmarshal(jsonBytes, &conversionPath2)
      if err != nil {
          log.Fatal(err)
      }
      models = append(models, model.KnowledgePattern{
        ID:             id,
        TaskID:         taskID,
        Type:           typeVal,
        Domain:         domain,
        TacitKnowledge: tacitKnowledge,
        ExplicitForm:   explicitForm,
        ConversionPath: conversionPath2,
        Accuracy:       accuracy,
        Coverage:       coverage,
        Consistency:    consistency,
        AbstractLevel:  abstractLevel,
        CreatedAt:      createdAt,
        UpdatedAt:      updatedAt,
      })
    }

    // バッチインサート
    if len(models) > 0 {
      if err := db.CreateInBatches(&models, 1000).Error; err != nil {
        return fmt.Errorf("failed to insert optimization models: %w", err)
        count ++
      }
      fmt.Printf("Successfully seeded %d knowledge pattern models\n", len(models))
    }
    log.Printf("🎉 Import finished. Total inserted: len(models) %d", len(models))
    return nil
}