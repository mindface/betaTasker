package seed

import (
	"log"
	"time"

	"github.com/godotask/model"
	"gorm.io/gorm"
)

// SeedUsers - ユーザーデータのシード
func SeedUsers(db *gorm.DB) error {
	log.Println("Starting users seeding...")

	users := []model.User{
		{
			ID:           1,
			Username:     "tanaka_taro",
			Email:        "tanaka@example.com",
			PasswordHash: "hashed_password_1",
			Role:         "engineer",
			CreatedAt:    time.Now().Add(-30 * 24 * time.Hour),
			UpdatedAt:    time.Now().Add(-1 * 24 * time.Hour),
			IsActive:     true,
			Factor:       "manufacturing",
			Process:      "machining",
		},
		{
			ID:           2,
			Username:     "sato_hanako",
			Email:        "sato@example.com",
			PasswordHash: "hashed_password_2",
			Role:         "senior_engineer",
			CreatedAt:    time.Now().Add(-45 * 24 * time.Hour),
			UpdatedAt:    time.Now().Add(-2 * 24 * time.Hour),
			IsActive:     true,
			Factor:       "quality_control",
			Process:      "precision_machining",
		},
		{
			ID:           3,
			Username:     "yamada_ichiro",
			Email:        "yamada@example.com",
			PasswordHash: "hashed_password_3",
			Role:         "expert",
			CreatedAt:    time.Now().Add(-90 * 24 * time.Hour),
			UpdatedAt:    time.Now().Add(-1 * 24 * time.Hour),
			IsActive:     true,
			Factor:       "advanced_materials",
			Process:      "special_machining",
		},
		{
			ID:           4,
			Username:     "suzuki_misaki",
			Email:        "suzuki@example.com",
			PasswordHash: "hashed_password_4",
			Role:         "supervisor",
			CreatedAt:    time.Now().Add(-120 * 24 * time.Hour),
			UpdatedAt:    time.Now().Add(-3 * 24 * time.Hour),
			IsActive:     true,
			Factor:       "knowledge_transfer",
			Process:      "training",
		},
		{
			ID:           5,
			Username:     "takahashi_ken",
			Email:        "takahashi@example.com",
			PasswordHash: "hashed_password_5",
			Role:         "trainee",
			CreatedAt:    time.Now().Add(-7 * 24 * time.Hour),
			UpdatedAt:    time.Now().Add(-12 * time.Hour),
			IsActive:     true,
			Factor:       "basic_operations",
			Process:      "learning",
		},
		{
			ID:           6,
			Username:     "demo_user",
			Email:        "demo@example.com",
			PasswordHash: "demo_password",
			Role:         "demo",
			CreatedAt:    time.Now().Add(-1 * 24 * time.Hour),
			UpdatedAt:    time.Now(),
			IsActive:     true,
			Factor:       "demonstration",
			Process:      "demo",
		},
	}

	// Insert data with duplicate handling
	for _, user := range users {
		var existingUser model.User
		if err := db.Where("id = ?", user.ID).First(&existingUser).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// User doesn't exist, create it
				if err := db.Create(&user).Error; err != nil {
					log.Printf("Error inserting user %s: %v", user.ID, err)
					return err
				}
			} else {
				log.Printf("Error checking user %s: %v", user.ID, err)
				return err
			}
		} else {
			log.Printf("User %s already exists, skipping", user.ID)
		}
	}

	log.Printf("✓ Successfully seeded %d users", len(users))
	return nil
}