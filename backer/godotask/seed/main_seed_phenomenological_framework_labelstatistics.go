package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/godotask/lib"
	"github.com/godotask/model"
)

func main() {
	// PostgreSQL DSN
	dsn := "host=db user=dbgodotask password=dbgodotask dbname=dbgodotask port=5432 sslmode=disable"

	// DB接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// AutoMigrate
	if err := db.AutoMigrate(&model.PhenomenologicalFramework{}, &model.KnowledgePattern{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	now := time.Now()

	// 10件ループでSeed
	for i := 1; i <= 10; i++ {
		// 段階的な数値生成（回を追うごとに改善）
		accuracy := 0.6 + float64(i)*0.04 + rand.Float64()*0.05   // 0.65〜1.0へ向上

		coverage := 0.5 + float64(i)*0.05 + rand.Float64()*0.05   // 0.55〜1.0へ向上
		consistency := 0.4 + float64(i)*0.06 + rand.Float64()*0.05 // 0.46〜1.0へ向上

		// PhenomenologicalFramework
		pf := model.PhenomenologicalFramework{
			ID:          fmt.Sprintf("pf_%02d", i),
			TaskID:      rand.Intn(20) + 1,
			Name:        fmt.Sprintf("ロボット精度フレームワーク v%d", i),
			Description: fmt.Sprintf("試行回数 %d に基づいた段階的改善モデル", i),
			Goal:        "G: 位置決め精度±0.01mm達成",
			Scope:       "A: 6軸ロボットアームの動作範囲全体",
			Process: lib.JSON{
				"Pa":    "キャリブレーション→測定→補正→検証の反復プロセス",
				"steps": []string{"初期測定", "誤差解析", "補正値計算", "適用", "再測定"},
			},
			Result:       lib.JSON{"expected": fmt.Sprintf("精度向上ステージ%dでロボット位置決め精度を改善", i)},
			Feedback:     lib.JSON{"loop": "測定結果を補正プロセスにフィードバック"},
			LimitMin:     -0.01,
			LimitMax:     0.01,
			GoalFunction: "minimize(abs(measured_position - target_position))",
			AbstractLevel: []string{"low", "medium", "high"}[min(i/4, 2)], // 段階的に抽象度アップ
			Domain:       "robotics",
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		if err := db.Create(&pf).Error; err != nil {
			log.Printf("insert error PF %s: %v", pf.ID, err)
		}

		// KnowledgePattern
		kp := model.KnowledgePattern{
			ID:             fmt.Sprintf("kp_%02d", i),
			Type:           []string{"tacit", "explicit", "hybrid"}[i%3],
			Domain:         "robotics",
			TacitKnowledge: fmt.Sprintf("熟練工の『しっくりくる』感覚 (段階 %d)", i),
			ExplicitForm:   fmt.Sprintf("力覚センサ閾値: Fx<%.2fN, Fy<%.2fN, Tz<%.2fNm", 1.0-accuracy/2, 1.0-coverage/2, 1.0-consistency/2),
			ConversionPath: lib.JSON{
				"SECI":   []string{"共同化", "表出化", "連結化", "内面化"},
				"method": fmt.Sprintf("データ収集→分析→ルール化 (ステージ%d)", i),
			},
			Accuracy:      minf(accuracy, 1.0),
			Coverage:      minf(coverage, 1.0),
			Consistency:   minf(consistency, 1.0),
			AbstractLevel: []string{"low", "medium", "high"}[min(i/4, 2)],
			CreatedAt:     now,
			UpdatedAt:     now,
		}

		if err := db.Create(&kp).Error; err != nil {
			log.Printf("insert error KP %s: %v", kp.ID, err)
		}
	}

	fmt.Println("✅ PhenomenologicalFramework と KnowledgePattern の段階的Seed(10件)が完了しました。")
}

// min のユーティリティ
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minf(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
