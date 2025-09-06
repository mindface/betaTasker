package seed

import (
	"log"
	"time"

	"github.com/godotask/model"
	"gorm.io/gorm"
)

// SeedTasks - タスクデータのシード（正しい構造）
func SeedTasks(db *gorm.DB) error {
	log.Println("Starting tasks seeding...")

	// 日付のポインタを作成するヘルパー関数
	timePtr := func(t time.Time) *time.Time {
		return &t
	}

	tasks := []model.Task{
		{
			ID:          1,
			UserID:      1,
			Title:       "初品加工・基本寸法確認",
			Description: "NC旋盤を使用した基本的な切削加工と寸法確認作業",
			Date:        timePtr(time.Now().Add(-30 * 24 * time.Hour)),
			Status:      "completed",
			Priority:    3,
			CreatedAt:   time.Now().Add(-30 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-30 * 24 * time.Hour),
		},
		{
			ID:          2,
			UserID:      1,
			Title:       "材料硬度変動への対応",
			Description: "材料の硬度変動に対する切削条件の最適化",
			Date:        timePtr(time.Now().Add(-25 * 24 * time.Hour)),
			Status:      "completed",
			Priority:    4,
			CreatedAt:   time.Now().Add(-25 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-25 * 24 * time.Hour),
		},
		{
			ID:          3,
			UserID:      2,
			Title:       "真円度0.005mm以下の高精度加工",
			Description: "超高精度が要求される部品の加工技術",
			Date:        timePtr(time.Now().Add(-20 * 24 * time.Hour)),
			Status:      "completed",
			Priority:    5,
			CreatedAt:   time.Now().Add(-20 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-20 * 24 * time.Hour),
		},
		{
			ID:          4,
			UserID:      3,
			Title:       "インコネル718の高効率加工",
			Description: "難削材であるインコネル718の高効率加工技術",
			Date:        timePtr(time.Now().Add(-15 * 24 * time.Hour)),
			Status:      "in_progress",
			Priority:    5,
			CreatedAt:   time.Now().Add(-15 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-15 * 24 * time.Hour),
		},
		{
			ID:          5,
			UserID:      3,
			Title:       "若手技術者への体系的指導",
			Description: "技術知識とスキルの体系的な伝承",
			Date:        timePtr(time.Now().Add(-10 * 24 * time.Hour)),
			Status:      "pending",
			Priority:    4,
			CreatedAt:   time.Now().Add(-10 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-10 * 24 * time.Hour),
		},
		{
			ID:          6,
			UserID:      2,
			Title:       "ロボット協働組立作業",
			Description: "人間とロボットが協働する組立作業の最適化",
			Date:        timePtr(time.Now().Add(-12 * 24 * time.Hour)),
			Status:      "completed",
			Priority:    4,
			CreatedAt:   time.Now().Add(-12 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-12 * 24 * time.Hour),
		},
		{
			ID:          7,
			UserID:      3,
			Title:       "AI駆動品質検査システム",
			Description: "AIを活用した自動品質検査システムの構築",
			Date:        timePtr(time.Now().Add(-8 * 24 * time.Hour)),
			Status:      "in_progress",
			Priority:    5,
			CreatedAt:   time.Now().Add(-8 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-8 * 24 * time.Hour),
		},
		{
			ID:          8,
			UserID:      2,
			Title:       "予知保全システム構築",
			Description: "IoTセンサーを活用した設備の予知保全システム",
			Date:        timePtr(time.Now().Add(-6 * 24 * time.Hour)),
			Status:      "in_progress",
			Priority:    4,
			CreatedAt:   time.Now().Add(-6 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-6 * 24 * time.Hour),
		},
		{
			ID:          9,
			UserID:      3,
			Title:       "デジタルツイン活用最適化",
			Description: "デジタルツイン技術を活用した製造プロセス最適化",
			Date:        timePtr(time.Now().Add(-4 * 24 * time.Hour)),
			Status:      "pending",
			Priority:    3,
			CreatedAt:   time.Now().Add(-4 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-4 * 24 * time.Hour),
		},
		{
			ID:          10,
			UserID:      1,
			Title:       "サスティナブル製造システム",
			Description: "環境負荷を最小化する持続可能な製造システムの設計",
			Date:        timePtr(time.Now().Add(-2 * 24 * time.Hour)),
			Status:      "pending",
			Priority:    3,
			CreatedAt:   time.Now().Add(-2 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-2 * 24 * time.Hour),
		},
	}

	// Insert data with duplicate handling
	for _, task := range tasks {
		var existingTask model.Task
		if err := db.Where("id = ?", task.ID).First(&existingTask).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Task doesn't exist, create it
				if err := db.Create(&task).Error; err != nil {
					log.Printf("Error inserting task %d: %v", task.ID, err)
					return err
				}
			} else {
				log.Printf("Error checking task %d: %v", task.ID, err)
				return err
			}
		} else {
			log.Printf("Task %d already exists, skipping", task.ID)
		}
	}

	log.Printf("✓ Successfully seeded %d tasks", len(tasks))
	return nil
}