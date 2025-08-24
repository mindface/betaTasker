package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"database/sql"
	"fmt"
	"log"
	"time"
	"github.com/godotask/model"
	// "github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	dsn := "host=db user=dbgodotask password=dbgodotask dbname=dbgodotask port=5432 sslmode=disable"
	var err error
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	// シードデータ投入前にテーブルを空にする
	// fmt.Println("テーブルをクリアしています...")

	// // 外部キー制約を考慮して、子テーブルから削除
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

	// シーケンスをリセット（PostgreSQL）
	_, err = sqlDB.Exec("ALTER SEQUENCE memory_contexts_id_seq RESTART WITH 1")
	if err != nil {
		log.Printf("memory_contexts_id_seqリセットエラー: %v", err)
	}
	
	_, err = sqlDB.Exec("ALTER SEQUENCE technical_factors_id_seq RESTART WITH 1")
	if err != nil {
		log.Printf("technical_factors_id_seqリセットエラー: %v", err)
	}
	
	_, err = sqlDB.Exec("ALTER SEQUENCE knowledge_transformations_id_seq RESTART WITH 1")
	if err != nil {
		log.Printf("knowledge_transformations_id_seqリセットエラー: %v", err)
	}

	fmt.Println("テーブルクリア完了。シードデータを投入します...")

	// --- Level 1 ---
	level1Data := []struct {
		workTarget string
		changeFactor string
		goal string
		toolSpec string
		concern string
		countermeasure string
		learnedKnowledge string
	}{
		{
			workTarget: "[職務カテゴリ: 切削・初品確認 / 分類コード: MA-C-01] 対象工程 L1-1: 初品加工・基本寸法確認",
			changeFactor: "新規ロット材導入（ロット番号: A-123）",
			goal: "初品寸法公差内維持、不良率5%以下",
			toolSpec: "TNMG160408 (汎用), 標準コーティング",
			concern: "初回切削でのバリ発生（要因: 切削条件未調整）",
			countermeasure: "メーカー推奨値での標準切削開始、目視での品質確認",
			learnedKnowledge: "基本的な切削条件とバリ発生の関係を理解。目視確認の重要性を認識。",
		},
		{
			workTarget: "[職務カテゴリ: 切削・初品確認 / 分類コード: MA-C-01] 対象工程 L1-2: 寸法精度の安定化",
			changeFactor: "加工開始後の寸法ばらつき発生",
			goal: "寸法公差±0.02mm以内、バリ発生抑制",
			toolSpec: "TNMG160408 (汎用), 標準コーティング",
			concern: "加工開始時の寸法不安定（要因: 機械暖機不足）",
			countermeasure: "機械暖機15分実施、初品3個での寸法確認",
			learnedKnowledge: "機械の暖機時間と寸法精度の関係を理解。初品確認の重要性を認識。",
		},
		{
			workTarget: "[職務カテゴリ: 切削・初品確認 / 分類コード: MA-C-01] 対象工程 L1-3: 表面品質の改善",
			changeFactor: "表面粗さの悪化傾向",
			goal: "Ra 3.2以下達成、外観品質向上",
			toolSpec: "TNMG160408 (汎用), 標準コーティング",
			concern: "表面に縦筋発生（要因: 送り速度過大）",
			countermeasure: "送り速度を20%減少、切削油流量増加",
			learnedKnowledge: "送り速度と表面粗さの基本的な関係を理解。切削油の効果を実感。",
		},
		{
			workTarget: "[職務カテゴリ: 切削・初品確認 / 分類コード: MA-C-01] 対象工程 L1-4: 工具寿命の認識",
			changeFactor: "工具摩耗による品質低下",
			goal: "工具寿命の把握、交換タイミング習得",
			toolSpec: "TNMG160408 (汎用), 標準コーティング",
			concern: "工具摩耗による寸法精度悪化",
			countermeasure: "100個加工ごとの工具点検、摩耗状況記録",
			learnedKnowledge: "工具摩耗の進行と品質への影響を理解。定期点検の重要性を認識。",
		},
		{
			workTarget: "[職務カテゴリ: 切削・初品確認 / 分類コード: MA-C-01] 対象工程 L1-5: 基本加工条件の習得",
			changeFactor: "材料硬度のばらつき",
			goal: "安定した加工条件の確立、不良率3%以下",
			toolSpec: "TNMG160408 (汎用), 標準コーティング",
			concern: "材料硬度変化による切削抵抗増加",
			countermeasure: "材料硬度測定、切削条件の微調整",
			learnedKnowledge: "材料特性の違いと加工条件の関係を理解。測定の重要性を認識。",
		},
	}

	for i, data := range level1Data {
		ctx := model.MemoryContext{
			UserID:     1,
			TaskID:     i + 1,
			Level:      1,
			WorkTarget: data.workTarget,
			Machine:    "NC旋盤（Mazak QT-200）",
			MaterialSpec:   "AISI 4340鋼材",
			ChangeFactor:     data.changeFactor,
			Goal:       data.goal,
			CreatedAt:  time.Now(),
		}

		var contextID int
		err := sqlDB.QueryRow(`
			INSERT INTO memory_contexts (user_id, task_id, level, work_target, machine, material_spec, change_factor, goal, created_at)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
			RETURNING id
		`, ctx.UserID, ctx.TaskID, ctx.Level, ctx.WorkTarget, ctx.Machine, ctx.MaterialSpec, ctx.ChangeFactor, ctx.Goal, ctx.CreatedAt).Scan(&contextID)
		if err != nil {
			log.Printf("insert context err: %v", err)
			continue
		}

		tf := model.TechnicalFactor{
			ContextID:   contextID,
			ToolSpec:    data.toolSpec,
			EvalFactors: "寸法精度, バリの有無, 面粗度",
			Measurement: "ノギス, マイクロメータ, 目視確認",
			Concern:     data.concern,
			CreatedAt:   time.Now(),
		}

		_, err = sqlDB.Exec(`
			INSERT INTO technical_factors (context_id, tool_spec, eval_factors, measurement_method, concern, created_at)
			VALUES ($1,$2,$3,$4,$5,$6)
		`, tf.ContextID, tf.ToolSpec, tf.EvalFactors, tf.Measurement, tf.Concern, tf.CreatedAt)
		if err != nil {
			log.Printf("insert technical_factor err: %v", err)
		}

		kt := model.KnowledgeTransformation{
			ContextID:       contextID,
			Transformation:  "基本的な加工現象の理解",
			Countermeasure:  data.countermeasure,
			ModelFeedback:   "基本的な切削理論の適用",
			LearnedKnowledge: data.learnedKnowledge,
			CreatedAt:       time.Now(),
		}

		_, err = sqlDB.Exec(`
			INSERT INTO knowledge_transformations (context_id, transformation, countermeasure, model_feedback, learned_knowledge, created_at)
			VALUES ($1,$2,$3,$4,$5,$6)
		`, kt.ContextID, kt.Transformation, kt.Countermeasure, kt.ModelFeedback, kt.LearnedKnowledge, kt.CreatedAt)
		if err != nil {
			log.Printf("insert knowledge_transformation err: %v", err)
		}
	}

	fmt.Println("Level 1データ投入完了")

	// --- Level 2 ---
	level2Data := []struct {
		workTarget string
		changeFactor string
		goal string
		materialSpec string
		toolSpec string
		concern string
		countermeasure string
		learnedKnowledge string
	}{
		{
			workTarget: "[職務カテゴリ: 品質改善・条件最適化 / 分類コード: MA-Q-02] 対象工程 L2-1: 定期点検後の加工条件の再確認と簡易的な最適化",
			changeFactor: "工具交換後に加工面がやや荒れる傾向が見られた",
			goal: "Ra 1.2以下の安定維持",
			materialSpec: "AISI 4340鋼材",
			toolSpec: "汎用チップ (ISO CNMG), Al2O3コーティング",
			concern: "摩耗による面粗度のばらつき",
			countermeasure: "切削速度を10%下げ、工具寿命を延ばす方向で設定変更",
			learnedKnowledge: "工具寿命と粗さのトレードオフを初めて意識",
		},
		{
			workTarget: "[職務カテゴリ: 品質改善・条件最適化 / 分類コード: MA-Q-02] 対象工程 L2-2: 再研磨工具使用後の加工品質改善と標準条件の見直し",
			changeFactor: "研磨済み工具による面粗度の悪化傾向",
			goal: "Ra 1.0以下を安定して達成",
			materialSpec: "AISI 4340鋼材",
			toolSpec: "再研磨チップ + AlTiNコーティング",
			concern: "チップごとの仕上がりのばらつき",
			countermeasure: "仕上げパスで送りを15%低減して、粗さ改善を狙う",
			learnedKnowledge: "送り速度と粗さに直線的な関係がないことを実感",
		},
		{
			workTarget: "[職務カテゴリ: 品質改善・条件最適化 / 分類コード: MA-Q-02] 対象工程 L2-3: SUS304への材質変更後の加工条件最適化",
			changeFactor: "新材料により切りくず絡みと摩耗進行が急増",
			goal: "Ra 0.8以下、工具寿命の延伸と切りくず処理の安定化",
			materialSpec: "SUS304ステンレス鋼材",
			toolSpec: "TNMG160408 (Sumitomo), AlTiNコーティング, SUS304向けブレーカ形状",
			concern: "高回転での振動、冷却不足、粘り材質の切りくず絡み",
			countermeasure: "切削速度15%低減、クーラント圧1.5倍、送り速度を20%増加して切りくず分断を狙う",
			learnedKnowledge: "送りをあえて増やすことで切りくずが分断しやすくなるという逆転の発想",
		},
		{
			workTarget: "[職務カテゴリ: 品質改善・条件最適化 / 分類コード: MA-Q-02] 対象工程 L2-4: SUS304対応の複合条件下における安定加工の探索",
			changeFactor: "振動＋冷却不足＋材質粘性が同時に悪影響",
			goal: "粗さRa 0.8以下とZ軸送り精度0.01mm以下の両立",
			materialSpec: "SUS304ステンレス鋼材",
			toolSpec: "特殊形状ブレーカー付きPVDチップ、内径冷却対応",
			concern: "送り方向の振動、Z軸送りずれによる精度不良",
			countermeasure: "送り速度と切削深さの最適交差点を探索（DoE実施）",
			learnedKnowledge: "多変量条件下では単一要因での最適化は不可能であり、交互作用を前提とした設計が必要",
		},
		{
			workTarget: "[職務カテゴリ: 品質改善・条件最適化 / 分類コード: MA-Q-02] 対象工程 L2-5: SUS304量産加工に向けた自動補正・異常予測モデルの導入検証",
			changeFactor: "加工途中での条件逸脱と製品ばらつきが発生",
			goal: "Ra 0.8以下を24時間安定保持、自動補正アルゴリズムの有効性確認",
			materialSpec: "SUS304ステンレス鋼材",
			toolSpec: "センサ付き工具（摩耗・振動計測）、AI対応切削条件マップ搭載",
			concern: "加工状態の自律変化に対するモデル遅延",
			countermeasure: "AIモデルによるリアルタイム補正の導入、履歴データによる閾値学習",
			learnedKnowledge: "加工プロセスを静的ではなく動的・予測的に捉える視点が重要。センサ統合による知識拡張が次世代標準となる",
		},
	}

	for i, data := range level2Data {
		ctx := model.MemoryContext{
			UserID:     1,
			TaskID:     i + 1,
			Level:      2,
			WorkTarget: data.workTarget,
			Machine:    "NC旋盤（Mazak QT-200）",
			MaterialSpec:   data.materialSpec,
			ChangeFactor:     data.changeFactor,
			Goal:       data.goal,
			CreatedAt:  time.Now(),
		}

		var contextID int
		err := sqlDB.QueryRow(`
			INSERT INTO memory_contexts (user_id, task_id, level, work_target, machine, material_spec, change_factor, goal, created_at)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
			RETURNING id
		`, ctx.UserID, ctx.TaskID, ctx.Level, ctx.WorkTarget, ctx.Machine, ctx.MaterialSpec, ctx.ChangeFactor, ctx.Goal, ctx.CreatedAt).Scan(&contextID)
		if err != nil {
			log.Printf("insert context err: %v", err)
			continue
		}

		tf := model.TechnicalFactor{
			ContextID:   contextID,
			ToolSpec:    data.toolSpec,
			EvalFactors: "加工面粗度, 工具刃先の摩耗状態, 切削音の変化, 切りくず形状",
			Measurement: "表面粗さ計, 顕微鏡観察, 騒音計データ, 切りくず形状観察",
			Concern:     data.concern,
			CreatedAt:   time.Now(),

			// Factor:            "材料硬度",          // 任意の初期値（レベル別に変えてもOK）
			// Process:           "切削→評価→補正",
			// EvaluationAxis:    "面粗度, 寸法精度",
			// InformationAmount: "3件の加工ログ、1件の参考資料",

		}

		_, err = sqlDB.Exec(`
			INSERT INTO technical_factors (context_id, tool_spec, eval_factors, measurement_method, concern, created_at)
			VALUES ($1,$2,$3,$4,$5,$6)
		`, tf.ContextID, tf.ToolSpec, tf.EvalFactors, tf.Measurement, tf.Concern, tf.CreatedAt)
		if err != nil {
			log.Printf("insert technical_factor err: %v", err)
		}

		kt := model.KnowledgeTransformation{
			ContextID:       contextID,
			Transformation:  "条件最適化と材料特性の理解",
			Countermeasure:  data.countermeasure,
			ModelFeedback:   "工具摩耗パターンと加工条件の相関関係知識追加",
			LearnedKnowledge: data.learnedKnowledge,
			CreatedAt:       time.Now(),
		}

		_, err = sqlDB.Exec(`
			INSERT INTO knowledge_transformations (context_id, transformation, countermeasure, model_feedback, learned_knowledge, created_at)
			VALUES ($1,$2,$3,$4,$5,$6)
		`, kt.ContextID, kt.Transformation, kt.Countermeasure, kt.ModelFeedback, kt.LearnedKnowledge, kt.CreatedAt)
		if err != nil {
			log.Printf("insert knowledge_transformation err: %v", err)
		}
	}

	fmt.Println("Level 2データ投入完了")

	// --- Level 3 ---
	level3Data := []struct {
		workTarget string
		changeFactor string
		goal string
		materialSpec string
		toolSpec string
		concern string
		countermeasure string
		learnedKnowledge string
	}{
		{
			workTarget: "[職務カテゴリ: プロセス改善・予知保全 / 分類コード: PM-P-03] 対象工程 L3-1: 高硬度材（Inconel 718）の基礎加工条件確立",
			changeFactor: "新規高硬度材導入による従来条件の適用困難",
			goal: "基本的な加工可能性の確認、初期条件の設定",
			materialSpec: "Inconel 718合金",
			toolSpec: "CNMG120408 (Kennametal), 超硬合金基本コーティング",
			concern: "加工硬化による工具寿命の極端な短縮",
			countermeasure: "超低速・軽切削での条件探索、頻繁な工具交換での対応",
			learnedKnowledge: "高硬度材では従来の加工理論が通用しない。材料特性の事前理解が必須",
		},
		{
			workTarget: "[職務カテゴリ: プロセス改善・予知保全 / 分類コード: PM-P-03] 対象工程 L3-2: Inconel 718専用工具による加工条件最適化",
			changeFactor: "専用工具導入による加工特性の変化",
			goal: "工具寿命30%向上、加工精度の安定化",
			materialSpec: "Inconel 718合金",
			toolSpec: "CNMG120408 (Kennametal), CBNコーティング、専用チップブレーカー",
			concern: "高コスト工具の効果的活用法の確立",
			countermeasure: "工具特性を活かした条件設定、段階的な条件調整",
			learnedKnowledge: "工具選択が加工結果に決定的な影響を与える。投資対効果の計算が重要",
		},
		{
			workTarget: "[職務カテゴリ: プロセス改善・予知保全 / 分類コード: PM-P-03] 対象工程 L3-3: 熱変位補正システムの導入と検証",
			changeFactor: "長時間加工での熱変位による精度悪化",
			goal: "24時間連続加工での精度維持、熱変位±0.005mm以下",
			materialSpec: "Inconel 718合金",
			toolSpec: "温度センサ内蔵工具、熱変位補正対応チップ",
			concern: "熱変位による寸法精度の経時変化",
			countermeasure: "リアルタイム熱変位補正、予測モデルによる事前補正",
			learnedKnowledge: "機械の熱的特性を理解し、予測制御することで精度を維持できる",
		},
		{
			workTarget: "[職務カテゴリ: プロセス改善・予知保全 / 分類コード: PM-P-03] 対象工程 L3-4: AI予測モデルによる工具交換時期の最適化",
			changeFactor: "工具摩耗の予測困難による突発的な品質問題",
			goal: "工具交換時期の予測精度90%以上、突発停止50%削減",
			materialSpec: "Inconel 718合金",
			toolSpec: "IoTセンサ内蔵工具、AI解析対応データ収集システム",
			concern: "複雑な摩耗パターンの予測精度向上",
			countermeasure: "機械学習モデルによる多変量解析、リアルタイム予測システム構築",
			learnedKnowledge: "データ駆動型のアプローチにより、経験に頼らない予測が可能。継続的学習が鍵",
		},
		{
			workTarget: "[職務カテゴリ: プロセス改善・予知保全 / 分類コード: PM-P-03] 対象工程 L3-5: 自律制御システムによる無人加工の実現",
			changeFactor: "人的リソースの制約と24時間稼働の要求",
			goal: "完全無人加工24時間稼働、品質維持率98%以上、機材調整回数70%削減",
			materialSpec: "Inconel 718合金",
			toolSpec: "完全自律制御対応工具システム、ロボット工具交換対応",
			concern: "予期しない異常への自律対応能力",
			countermeasure: "多重フェールセーフシステム、異常パターン学習による自律回復機能",
			learnedKnowledge: "単なる自動化ではなく、自律的な判断・学習・適応能力を持つシステムが次世代製造の基盤。人間の役割は監視から戦略立案へシフト",
		},
	}

	for i, data := range level3Data {
		ctx := model.MemoryContext{
			UserID:     1,
			TaskID:     i + 1,
			Level:      3,
			WorkTarget: data.workTarget,
			Machine:    "NC旋盤（Mazak QT-200）",
			MaterialSpec:   data.materialSpec,
			ChangeFactor:     data.changeFactor,
			Goal:       data.goal,
			CreatedAt:  time.Now(),
		}

		var contextID int
		err := sqlDB.QueryRow(`
			INSERT INTO memory_contexts (user_id, task_id, level, work_target, machine, material_spec, change_factor, goal, created_at)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
			RETURNING id
		`, ctx.UserID, ctx.TaskID, ctx.Level, ctx.WorkTarget, ctx.Machine, ctx.MaterialSpec, ctx.ChangeFactor, ctx.Goal, ctx.CreatedAt).Scan(&contextID)
		if err != nil {
			log.Printf("insert context err: %v", err)
			continue
		}

		tf := model.TechnicalFactor{
			ContextID:   contextID,
			ToolSpec:    data.toolSpec,
			EvalFactors: "加工サイクルタイム, 工具交換頻度, 設備稼働率, 異常振動パターン, 工具摩耗予測精度",
			Measurement: "AI画像解析, IoTセンサーモニタリング, 自動データ収集システム, 予測モデル精度評価",
			Concern:     data.concern,
			CreatedAt:   time.Now(),
		}

		_, err = sqlDB.Exec(`
			INSERT INTO technical_factors (context_id, tool_spec, eval_factors, measurement_method, concern, created_at)
			VALUES ($1,$2,$3,$4,$5,$6)
		`, tf.ContextID, tf.ToolSpec, tf.EvalFactors, tf.Measurement, tf.Concern, tf.CreatedAt)
		if err != nil {
			log.Printf("insert technical_factor err: %v", err)
		}

		kt := model.KnowledgeTransformation{
			ContextID:       contextID,
			Transformation:  "高度な制御システムと予測技術の統合",
			Countermeasure:  data.countermeasure,
			ModelFeedback:   "多変量解析モデルの構築、予測精度向上、自律制御システムの学習",
			LearnedKnowledge: data.learnedKnowledge,
			CreatedAt:       time.Now(),
		}

		_, err = sqlDB.Exec(`
			INSERT INTO knowledge_transformations (context_id, transformation, countermeasure, model_feedback, learned_knowledge, created_at)
			VALUES ($1,$2,$3,$4,$5,$6)
		`, kt.ContextID, kt.Transformation, kt.Countermeasure, kt.ModelFeedback, kt.LearnedKnowledge, kt.CreatedAt)
		if err != nil {
			log.Printf("insert knowledge_transformation err: %v", err)
		}
	}

	fmt.Println("Level 3データ投入完了")
	fmt.Println("全てのシードデータの投入が完了しました")
}
