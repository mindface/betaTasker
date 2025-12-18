package seed

import (
	"fmt"
	"log"
	"time"

	"github.com/godotask/model"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

// SeedMemoryContextsLegacy - レガシーメモリコンテキストのシードデータ
func SeedMemoryContextsLegacy() error {
	dsn := "host=db user=dbgodotask password=dbgodotask dbname=dbgodotask port=5432 sslmode=disable"
	var err error
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
		return err
	}

	// シードデータ投入前にテーブルを空にする
	fmt.Println("テーブルをクリアしています...")

	// 外部キー制約を考慮して、子テーブルから削除
	_, err = sqlDB.Exec("DELETE FROM knowledge_transformations")
	if err != nil {
		log.Printf("knowledge_transformations削除エラー: %v", err)
	}

	_, err = sqlDB.Exec("DELETE FROM technical_factors")
	if err != nil {
		log.Printf("technical_factors削除エラー: %v", err)
	}

	_, err = sqlDB.Exec("DELETE FROM memory_contexts")
	if err != nil {
		log.Printf("memory_contexts削除エラー: %v", err)
	}

	fmt.Println("テーブルクリア完了")

	// Level 1のメモリコンテキスト
	level1Context := model.MemoryContext{
		UserID:       1,
		TaskID:       1,
		Level:        1,
		WorkTarget:   "初品加工・基本寸法確認",
		Machine:      "NC旋盤（Mazak QT-200）",
		MaterialSpec: "AISI 4340鋼材",
		ChangeFactor: "材料の初期バラツキ",
		Goal:         "基本的な切削条件と品質の関係を理解",
		CreatedAt:    time.Now(),
	}

	if err := db.Create(&level1Context).Error; err != nil {
		log.Printf("insert memory_context err: %v", err)
		return err
	}

	// Level 1のTechnicalFactors
	level1TechnicalFactors := []model.TechnicalFactor{
		{
			ContextID:         level1Context.ID,
			ToolSpec:          "TNMG160408 (汎用) 標準コーティング",
			EvalFactors:       "表面粗さ: Ra3.2→1.6",
			MeasurementMethod: "ノギス計測（サンプル頻度改善）",
			Concern:           "バリ発生、寸法精度",
			CreatedAt:         time.Now(),
		},
		{
			ContextID:         level1Context.ID,
			ToolSpec:          "ノギス（デジタル、0.01mm精度）",
			EvalFactors:       "測定精度±0.02mm",
			MeasurementMethod: "3回測定の平均値採用",
			Concern:           "測定ポイントのばらつき",
			CreatedAt:         time.Now(),
		},
	}

	for _, tf := range level1TechnicalFactors {
		if err := db.Create(&tf).Error; err != nil {
			log.Printf("insert technical_factor err: %v", err)
		}
	}

	// Level 1のKnowledgeTransformations
	level1KnowledgeTransformations := []model.KnowledgeTransformation{
		{
			ContextID:         level1Context.ID,
			Transformation:    "切削速度V: 100→120m/min",
			Countermeasure:    "バリ発生→面取り追加",
			ModelFeedback:     "切削条件DB更新",
			LearnedKnowledge:  "材料・工具の基本相性理解",
			CreatedAt:         time.Now(),
		},
		{
			ContextID:         level1Context.ID,
			Transformation:    "送り速度f: 0.2→0.15mm/rev",
			Countermeasure:    "表面粗さ改善",
			ModelFeedback:     "加工時間増加を許容",
			LearnedKnowledge:  "品質と効率のトレードオフ基礎",
			CreatedAt:         time.Now(),
		},
	}

	for _, kt := range level1KnowledgeTransformations {
		if err := db.Create(&kt).Error; err != nil {
			log.Printf("insert knowledge_transformation err: %v", err)
		}
	}

	fmt.Println("Level 1データ投入完了")

	// Level 2のメモリコンテキスト
	level2Context := model.MemoryContext{
		UserID:       1,
		TaskID:       2,
		Level:        2,
		WorkTarget:   "材料硬度変動への対応",
		Machine:      "NC旋盤（Mazak QT-200）",
		MaterialSpec: "AISI 4340鋼材（HRC 28-32）",
		ChangeFactor: "材料硬度の変動（±2HRC）",
		Goal:         "硬度変動に応じた切削条件の最適化",
		CreatedAt:    time.Now(),
	}

	if err := db.Create(&level2Context).Error; err != nil {
		log.Printf("insert memory_context err: %v", err)
		return err
	}

	// Level 2のTechnicalFactors
	level2TechnicalFactors := []model.TechnicalFactor{
		{
			ContextID:         level2Context.ID,
			ToolSpec:          "汎用チップ (ISO CNMG) Al2O3コーティング",
			EvalFactors:       "工具摩耗速度の変動",
			MeasurementMethod: "工具顕微鏡による摩耗量測定",
			Concern:           "予期せぬ工具寿命短縮",
			CreatedAt:         time.Now(),
		},
		{
			ContextID:         level2Context.ID,
			ToolSpec:          "硬度計（ロックウェル）",
			EvalFactors:       "材料硬度HRC 28-32",
			MeasurementMethod: "ロット毎の硬度測定",
			Concern:           "硬度分布の不均一性",
			CreatedAt:         time.Now(),
		},
	}

	for _, tf := range level2TechnicalFactors {
		if err := db.Create(&tf).Error; err != nil {
			log.Printf("insert technical_factor err: %v", err)
		}
	}

	// Level 2のKnowledgeTransformations
	level2KnowledgeTransformations := []model.KnowledgeTransformation{
		{
			ContextID:         level2Context.ID,
			Transformation:    "硬度HRC30以上で切削速度10%減",
			Countermeasure:    "工具摩耗監視頻度増加",
			ModelFeedback:     "硬度-切削条件マトリクス作成",
			LearnedKnowledge:  "材料特性に応じた動的調整の重要性",
			CreatedAt:         time.Now(),
		},
		{
			ContextID:         level2Context.ID,
			Transformation:    "切込み深さ調整: 2.0→1.5mm",
			Countermeasure:    "パス数増加による負荷分散",
			ModelFeedback:     "工具寿命20%延長確認",
			LearnedKnowledge:  "工具寿命と生産性のバランス最適化",
			CreatedAt:         time.Now(),
		},
	}

	for _, kt := range level2KnowledgeTransformations {
		if err := db.Create(&kt).Error; err != nil {
			log.Printf("insert knowledge_transformation err: %v", err)
		}
	}

	fmt.Println("Level 2データ投入完了")

	// Level 3のメモリコンテキスト
	level3Context := model.MemoryContext{
		UserID:       2,
		TaskID:       3,
		Level:        3,
		WorkTarget:   "真円度0.005mm以下の高精度加工",
		Machine:      "5軸マシニングセンタ（DMG MORI）",
		MaterialSpec: "チタン合金 Ti-6Al-4V",
		ChangeFactor: "温度変化による熱変形",
		Goal:         "高精度形状の安定的実現",
		CreatedAt:    time.Now(),
	}

	if err := db.Create(&level3Context).Error; err != nil {
		log.Printf("insert memory_context err: %v", err)
		return err
	}

	// Level 3のTechnicalFactors
	level3TechnicalFactors := []model.TechnicalFactor{
		{
			ContextID:         level3Context.ID,
			ToolSpec:          "CBN工具 高速切削対応",
			EvalFactors:       "真円度0.005mm達成",
			MeasurementMethod: "三次元測定機による形状測定",
			Concern:           "熱変形による精度劣化",
			CreatedAt:         time.Now(),
		},
		{
			ContextID:         level3Context.ID,
			ToolSpec:          "温度センサー（非接触式）",
			EvalFactors:       "加工点温度±5℃管理",
			MeasurementMethod: "赤外線温度計によるリアルタイム監視",
			Concern:           "局所的な温度上昇",
			CreatedAt:         time.Now(),
		},
		{
			ContextID:         level3Context.ID,
			ToolSpec:          "高圧クーラント（70bar）",
			EvalFactors:       "切屑排出性向上",
			MeasurementMethod: "切屑形状・色の観察",
			Concern:           "切屑の巻き付き防止",
			CreatedAt:         time.Now(),
		},
	}

	for _, tf := range level3TechnicalFactors {
		if err := db.Create(&tf).Error; err != nil {
			log.Printf("insert technical_factor err: %v", err)
		}
	}

	// Level 3のKnowledgeTransformations
	level3KnowledgeTransformations := []model.KnowledgeTransformation{
		{
			ContextID:         level3Context.ID,
			Transformation:    "温度補正アルゴリズム実装",
			Countermeasure:    "リアルタイム座標補正",
			ModelFeedback:     "温度-変位相関モデル構築",
			LearnedKnowledge:  "環境制御の重要性と補正技術",
			CreatedAt:         time.Now(),
		},
		{
			ContextID:         level3Context.ID,
			Transformation:    "5軸同時制御による一発加工",
			Countermeasure:    "段取り誤差の排除",
			ModelFeedback:     "加工プログラム最適化",
			LearnedKnowledge:  "複合加工による累積誤差低減",
			CreatedAt:         time.Now(),
		},
		{
			ContextID:         level3Context.ID,
			Transformation:    "仕上げ代0.02mm均一化",
			Countermeasure:    "超精密仕上げパス追加",
			ModelFeedback:     "仕上げ条件データベース更新",
			LearnedKnowledge:  "段階的精度向上アプローチ",
			CreatedAt:         time.Now(),
		},
	}

	for _, kt := range level3KnowledgeTransformations {
		if err := db.Create(&kt).Error; err != nil {
			log.Printf("insert knowledge_transformation err: %v", err)
		}
	}

	fmt.Println("Level 3データ投入完了")
	fmt.Println("全てのシードデータの投入が完了しました")
	return nil
}