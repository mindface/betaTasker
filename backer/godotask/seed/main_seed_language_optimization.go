package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/godotask/model"
	"github.com/godotask/lib"
)

func main() {
	dsn := "host=db user=dbgodotask password=dbgodotask dbname=dbgodotask port=5432 sslmode=disable"
	var err error
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// シードデータ生成
	now := time.Now()
	abstractionLevels := []string{"low", "medium", "high"}
	domains := []string{"technical", "medical", "business", "legal", "casual"}
	ctx := lib.JSON{
		"source": "example",
		"lang":   "en",
	}

	trans := lib.JSON{
		"method": "rewrite",
		"focus":  "clarity",
	}
	for i := 1; i <= 50; i++ {
		original := fmt.Sprintf("This is a sample text number %d for optimization.", i)
		optimized := fmt.Sprintf("Optimized version of text number %d with improved clarity.", i)

		precision := rand.Float64()*0.4 + 0.6     // 0.6 - 1.0
		clarity := rand.Float64()*0.4 + 0.6       // 0.6 - 1.0
		completeness := rand.Float64()*0.4 + 0.6  // 0.6 - 1.0
		evalScore := (precision + clarity + completeness) / 3.0

    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

		lo := model.LanguageOptimization{
			ID:               fmt.Sprintf("lo_%03d", i),
			TaskID:           rand.Intn(20) + 1, // 適当な Task ID に紐づけ
			OriginalText:     original,
			OptimizedText:    optimized,
			Domain:           domains[rand.Intn(len(domains))],
			AbstractionLevel: abstractionLevels[rand.Intn(len(abstractionLevels))],
			Precision:        precision,
			Clarity:          clarity,
			Completeness:     completeness,
			Context:          ctx,
			Transformation:   trans,
			EvaluationScore:  evalScore,
			CreatedAt:        now,
			UpdatedAt:        now,
		}

		if err := db.Create(&lo).Error; err != nil {
			log.Printf("insert error for %s: %v", lo.ID, err)
		}
	}

	fmt.Println("✅ LanguageOptimization の seed データ生成が完了しました。")
}