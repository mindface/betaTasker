package seed

import (
	"log"
	"time"

	"github.com/godotask/infrastructure/db/model"
	"gorm.io/gorm"
)

// SeedMemoryContexts - シンプルなメモリコンテキストのシードデータ
func SeedMemoryContexts(db *gorm.DB) error {
	log.Println("Starting memory contexts seeding...")

	contexts := []model.MemoryContext{
		{
			UserID:       1,
			TaskID:       1,
			Level:        1,
			WorkTarget:   "[MA-C-01] 初品加工・基本寸法確認",
			Machine:      "NC旋盤（Mazak QT-200）",
			MaterialSpec: "AISI 4340鋼材",
			ChangeFactor: "材料の初期バラツキ",
			Goal:         "基本的な切削条件と品質の関係を理解",
			CreatedAt:    time.Now().Add(-30 * 24 * time.Hour),
		},
		{
			UserID:       1,
			TaskID:       2,
			Level:        2,
			WorkTarget:   "[MA-Q-02] 材料硬度変動への対応",
			Machine:      "NC旋盤（Mazak QT-200）",
			MaterialSpec: "AISI 4340鋼材（HRC 28-32）",
			ChangeFactor: "材料硬度の変動",
			Goal:         "切削条件の最適化",
			CreatedAt:    time.Now().Add(-25 * 24 * time.Hour),
		},
		{
			UserID:       2,
			TaskID:       3,
			Level:        3,
			WorkTarget:   "[MA-H-03] 真円度0.005mm以下の高精度加工",
			Machine:      "5軸マシニングセンタ（DMG MORI）",
			MaterialSpec: "チタン合金 Ti-6Al-4V",
			ChangeFactor: "温度変化と熱変形",
			Goal:         "高精度加工の実現",
			CreatedAt:    time.Now().Add(-20 * 24 * time.Hour),
		},
		{
			UserID:       3,
			TaskID:       4,
			Level:        4,
			WorkTarget:   "[MA-S-04] インコネル718の高効率加工",
			Machine:      "5軸マシニングセンタ（DMG MORI）",
			MaterialSpec: "インコネル718",
			ChangeFactor: "難削材特性",
			Goal:         "難削材の効率的加工",
			CreatedAt:    time.Now().Add(-15 * 24 * time.Hour),
		},
		{
			UserID:       3,
			TaskID:       5,
			Level:        5,
			WorkTarget:   "[MA-T-05] 若手技術者への体系的指導",
			Machine:      "全設備",
			MaterialSpec: "全材料",
			ChangeFactor: "技術レベルの差",
			Goal:         "技術伝承の体系化",
			CreatedAt:    time.Now().Add(-10 * 24 * time.Hour),
		},
	}

	// Insert data with duplicate handling
	for _, context := range contexts {
		var existingContext model.MemoryContext
		if err := db.Where("user_id = ? AND task_id = ?", context.UserID, context.TaskID).First(&existingContext).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Context doesn't exist, create it
				if err := db.Create(&context).Error; err != nil {
					log.Printf("Error inserting memory context for user %d, task %d: %v", context.UserID, context.TaskID, err)
					return err
				}
			} else {
				log.Printf("Error checking memory context for user %d, task %d: %v", context.UserID, context.TaskID, err)
				return err
			}
		} else {
			log.Printf("Memory context for user %d, task %d already exists, skipping", context.UserID, context.TaskID)
		}
	}

	log.Printf("✓ Successfully seeded %d memory contexts", len(contexts))
	return nil
}