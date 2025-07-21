package main

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
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
	ID         int        `gorm:"primaryKey"`
	UserID     int
	SourceType string
	Title      string
	Author     string
	Notes      string
	Tags       string
	ReadStatus string
	ReadDate   *time.Time
	Factor           string
	Process          string
	EvaluationAxis   string
	InformationAmount string
	CreatedAt  time.Time
	UpdatedAt  time.Time
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
	EffectivenessScore int
	EffortScore        int
	ImpactScore        int
	QualitativeFeedback string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dsn := "root:dbgodotask@tcp(127.0.0.1:3306)/dbgodotask?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connection failed")
	}

	db.AutoMigrate(&Book{}, &Memory{}, &Task{}, &Assessment{})

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
			UserID:     1,
			SourceType: "book",
			Title:      book.Title,
			Author:     fmt.Sprintf("Dr. Author %d", i+1),
			Notes:      "技術適用性評価用ノート",
			Tags:       "3D,素材,評価",
			ReadStatus: "finished",
			ReadDate:   &now,
			Factor:           "要素分析",
			Process:          "導入までのプロセス整理",
			EvaluationAxis:   "有効性・実装性",
			InformationAmount: "中",
			CreatedAt:  now,
			UpdatedAt:  now,
		}
		db.Create(&mem)
	}

	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	totalWeeks := 52 * 5
	entriesPerWeek := 5

	entryID := 1
	for week := 0; week < totalWeeks; week++ {
		for day := 0; day < entriesPerWeek; day++ {
			date := startDate.AddDate(0, 0, week*7+day)
			bookIndex := rand.Intn(len(books))

			mem := Memory{
				UserID:     1,
				SourceType: "book",
				Title:      books[bookIndex].Title,
				Author:     "Researcher",
				Notes:      "Study on 3D printing material",
				Tags:       "3D,Material,Alloy",
				ReadStatus: "finished",
				ReadDate:   &date,
				CreatedAt:  date,
				UpdatedAt:  date,
			}
			db.Create(&mem)

			task := Task{
				UserID:      1,
				MemoryID:    mem.ID,
				Title:       fmt.Sprintf("Evaluate %s", mem.Title),
				Description: "Evaluating adoption of material in job task",
				Date:        date,
				Status:      "completed",
				Priority:    2,
				CreatedAt:   date,
				UpdatedAt:   date,
			}
			db.Create(&task)

			assessment := Assessment{
				TaskID:              task.ID,
				UserID:              1,
				EffectivenessScore: rand.Intn(41) + 60, // 60-100
				EffortScore:        rand.Intn(51) + 40, // 40-90
				ImpactScore:        rand.Intn(46) + 50, // 50-95
				QualitativeFeedback: "Relevant to 3D alloy application.",
				CreatedAt:           date,
				UpdatedAt:           date,
			}
			db.Create(&assessment)

			entryID++
		}
	}

	fmt.Println("Seed data generation completed.")
}

