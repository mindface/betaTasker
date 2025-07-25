package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"strings" 

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID      int    `gorm:"primaryKey"`
	Title   string
	Name    string
	Text    string
	Disc    string
	ImgPath string
	Status  string
}

type Memory struct {
	ID                 int       `gorm:"primaryKey"`
	UserID             int
	SourceType         string
	Title              string
	Author             string
	Notes              string
	Tags               string
	ReadStatus         string
	ReadDate           *time.Time
	Factor             string
	Process            string
	EvaluationAxis     string
	InformationAmount  string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Task struct {
	ID          int       `gorm:"primaryKey"`
	UserID      int
	MemoryID    int
	Title       string
	Description string
	Date        time.Time
	Status      string
	Priority    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Assessment struct {
	ID                  int       `gorm:"primaryKey"`
	TaskID              int
	UserID              int
	EffectivenessScore  int
	EffortScore         int
	ImpactScore         int
	QualitativeFeedback string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func classifyScore(score int) string {
	switch {
	case score >= 95:
		return "s_plus" // 卓越（非常に優れている）
	case score >= 90:
		return "s"      // 優秀（文句なし）
	case score >= 85:
		return "a_plus" // 良好（高評価）
	case score >= 80:
		return "a"      // 実用性あり（改善余地あり）
	case score >= 75:
		return "b_plus" // 満足（条件付きで応用可能）
	case score >= 70:
		return "b"
	case score >= 65:
		return "c_plus"
	case score >= 60:
		return "c"
	case score >= 55:
		return "d_plus"
	case score >= 50:
		return "d"
	default:
		return "e"      // 再考・再検証が必要
	}
}

func generateNoteText(scoreClass string, tags []string) string {
	tagText := ""
	if len(tags) > 0 {
		tagText = fmt.Sprintf(" 対象タグ: %s。", strings.Join(tags, ", "))
	}

	templates := map[string]string{
		"s_plus": "この素材は極めて高い評価を得ており、実務レベルでの即時活用が推奨されます。",
		"s":      "高評価の対象であり、すでに業務応用に十分な水準に達しています。",
		"a_plus": "専門的な観点からも優れた成果が得られており、今後の応用が期待されます。",
		"a":      "実用化に十分耐えうる性能を示しており、特定用途への展開が可能です。",
		"b_plus": "使用条件次第では実装可能と判断される水準にあります。",
		"b":      "中程度の評価結果であり、さらなる検証が推奨されます。",
		"c_plus": "限定的な条件下での使用にとどまりそうですが、将来的な改善が期待されます。",
		"c":      "現時点での有用性は限定的であり、慎重な評価が求められます。",
		"d_plus": "多くの課題が残っており、基本的な再評価が必要です。",
		"d":      "導入には大きなハードルが存在し、要再検討項目です。",
		"e":      "評価結果は不十分で、使用は推奨できません。",
	}

	return templates[scoreClass] + tagText
}

func generateTaskDescription(scoreClass string) string {
	desc := map[string]string{
		"s_plus": "即時展開を想定し、実装フェーズへの移行を進める。",
		"s":      "実証フェーズを省略し、直接的な運用評価へと移行する。",
		"a_plus": "応用領域を拡大するための展開計画を立案する。",
		"a":      "プロトタイプ環境にて早期適用を試みる。",
		"b_plus": "部分的な業務活用を見据え、パイロット導入を検討する。",
		"b":      "追加評価と検証を行い、使用可否を精査する。",
		"c_plus": "限定分野にて試行的に導入して経過観察する。",
		"c":      "関連文献を踏まえつつ、慎重な検証計画を策定する。",
		"d_plus": "基本性能の再検証および改善案の策定を行う。",
		"d":      "使用中止とともに代替案の模索を開始する。",
		"e":      "対象から除外し、他の候補へリソースを集中させる。",
	}
	return desc[scoreClass]
}

func generateTaskTitle(scoreClass, baseTitle string) string {
	prefix := map[string]string{
		"s_plus": "最優先評価",
		"s":      "優先評価",
		"a_plus": "高評価素材",
		"a":      "実用検証",
		"b_plus": "試験導入",
		"b":      "検討中素材",
		"c_plus": "条件付き評価",
		"c":      "再検証候補",
		"d_plus": "懸念要素あり",
		"d":      "非推奨候補",
		"e":      "使用不可",
	}
	return fmt.Sprintf("%s: %s", prefix[scoreClass], baseTitle)
}


func seed() {
	rand.Seed(time.Now().UnixNano())

	// PostgreSQL DSN
	dsn := "host=dbgodotask user=dbgodotask password=dbgodotask dbname=dbgodotask port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// AutoMigrate
	err = db.AutoMigrate(&Book{}, &Memory{}, &Task{}, &Assessment{})
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	fmt.Println("テーブルをクリアしています...")
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("DBの取得に失敗しました: %v", err)
	}

	// 外部キー制約を考慮して、子テーブルから削除
	_, err = sqlDB.Exec("DELETE FROM assessments")
	if err != nil {
		log.Printf("assessments削除エラー: %v", err)
	}

	_, err = sqlDB.Exec("DELETE FROM tasks")
	if err != nil {
		log.Printf("tasks削除エラー: %v", err)
	}

	_, err = sqlDB.Exec("DELETE FROM memories")
	if err != nil {
		log.Printf("memories削除エラー: %v", err)
	}

	// シーケンスをリセット（PostgreSQL）
	_, err = sqlDB.Exec("ALTER SEQUENCE assessments_id_seq RESTART WITH 1")
	if err != nil {
		log.Printf("assessments_id_seqリセットエラー: %v", err)
	}

	_, err = sqlDB.Exec("ALTER SEQUENCE tasks_id_seq RESTART WITH 1")
	if err != nil {
		log.Printf("tasks_id_seqリセットエラー: %v", err)
	}

	_, err = sqlDB.Exec("ALTER SEQUENCE memories_id_seq RESTART WITH 1")
	if err != nil {
		log.Printf("memories_id_seqリセットエラー: %v", err)
	}

	books := []Book{
		{Title: "Advanced Metal Printing Techniques", Name: "高性能金属プリント技術", Text: "Selective Laser Meltingの最前線", Disc: "航空・医療応用", ImgPath: "/img/metal1.jpg", Status: "published"},
		{Title: "Understanding Titanium Alloys", Name: "チタン合金の基礎と応用", Text: "軽量高強度材料の解説", Disc: "バイオ・航空", ImgPath: "/img/titanium.jpg", Status: "published"},
		{Title: "Polymer-Metal Composites", Name: "ポリマーメタル複合材料", Text: "新素材開発のための複合手法", Disc: "強化設計", ImgPath: "/img/composite.jpg", Status: "published"},
		{Title: "Ceramics in Biomedical Printing", Name: "生体用セラミック材料", Text: "骨再建と耐熱材料の3D応用", Disc: "医療", ImgPath: "/img/ceramic.jpg", Status: "published"},
		{Title: "Reinforced Aluminum", Name: "強化アルミニウム技術", Text: "繊維強化や熱処理プロセス", Disc: "自動車・宇宙産業", ImgPath: "/img/aluminum.jpg", Status: "published"},
		{Title: "Stainless Steel Printing", Name: "ステンレススチールプリント", Text: "腐食耐性と加工性の両立", Disc: "食品・機械加工", ImgPath: "/img/steel.jpg", Status: "published"},
		{Title: "Biodegradable Implants", Name: "生分解性インプラント材料", Text: "Mg系合金と体内吸収制御", Disc: "医療現場での展開", ImgPath: "/img/biodegrade.jpg", Status: "published"},
		{Title: "Surface Treatment of Metals", Name: "金属表面処理の全技術", Text: "研磨・酸化・コーティング", Disc: "耐久性と接着性向上", ImgPath: "/img/surface.jpg", Status: "published"},
	}
	db.Create(&books)

	now := time.Now()

	for i, book := range books {
		mem := Memory{
			UserID:             1,
			SourceType:         "book",
			Title:              book.Title,
			Author:             fmt.Sprintf("Dr. Author %d", i+1),
			Notes:              "技術適用性評価用ノート",
			Tags:               "3D,素材,評価",
			ReadStatus:         "finished",
			ReadDate:           &now,
			Factor:             "要素分析",
			Process:            "導入までのプロセス整理",
			EvaluationAxis:     "有効性・実装性",
			InformationAmount:  "中",
			CreatedAt:          now,
			UpdatedAt:          now,
		}
		db.Create(&mem)
	}

	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	totalWeeks := 52 * 5
	entriesPerWeek := 5

	for week := 0; week < totalWeeks; week++ {
		for day := 0; day < entriesPerWeek; day++ {
			date := startDate.AddDate(0, 0, week*7+day)
			bookIndex := rand.Intn(len(books))

			// スコア生成と分類
			effectiveness := rand.Intn(41) + 60
			effort := rand.Intn(51) + 40
			impact := rand.Intn(46) + 50
			scoreClass := classifyScore(effectiveness) // メインの指標に応じて分類

			// テキスト生成
			tags := []string{"3Dプリント", "材料", "Mg合金"}
			noteText := generateNoteText(scoreClass, tags)
			desc := generateTaskDescription(scoreClass)
			title := generateTaskTitle(scoreClass, books[bookIndex].Title)

			// Memory
			mem := Memory{
				UserID:     1,
				SourceType: "book",
				Title:      books[bookIndex].Title,
				Author:     "Researcher",
				Notes:      noteText,
				Tags:       strings.Join(tags, ","),
				ReadStatus: "finished",
				ReadDate:   &date,
				CreatedAt:  date,
				UpdatedAt:  date,
			}
			db.Create(&mem)

			// Task
			task := Task{
				UserID:      1,
				MemoryID:    mem.ID,
				Title:       title,
				Description: desc,
				Date:        date,
				Status:      "completed",
				Priority:    2,
				CreatedAt:   date,
				UpdatedAt:   date,
			}
			db.Create(&task)

			// Assessment
			assessment := Assessment{
				TaskID:              task.ID,
				UserID:              1,
				EffectivenessScore:  effectiveness,
				EffortScore:         effort,
				ImpactScore:         impact,
				QualitativeFeedback: noteText,
				CreatedAt:           date,
				UpdatedAt:           date,
			}
			db.Create(&assessment)
		}
	}

	fmt.Println("✅ PostgreSQL用のseedデータ生成が完了しました。")
}
