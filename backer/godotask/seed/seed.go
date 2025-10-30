package seed

import (
	"log"

	"github.com/godotask/model"
	"gorm.io/gorm"
)

// SeedBooksSimple - 書籍のシンプルなシードデータ
func SeedBooksSimple() error {
	db := model.DB
	log.Println("Starting books and tasks seeding...")

	// Books
	books := []model.Book{
		{
			ID:      1,
			Title:   "Go Programming Language",
			Name:    "Go Programming",
			Text:    "The authoritative resource on Go programming",
			Disc:    "Programming guide for Go language",
			ImgPath: "/images/go-book.jpg",
			Status:  "available",
		},
		{
			ID:      2,
			Title:   "Clean Code",
			Name:    "Clean Code Handbook",
			Text:    "A handbook of agile software craftsmanship",
			Disc:    "Software engineering best practices",
			ImgPath: "/images/clean-code.jpg",
			Status:  "available",
		},
		{
			ID:      3,
			Title:   "Design Patterns",
			Name:    "GoF Design Patterns",
			Text:    "Elements of reusable object-oriented software",
			Disc:    "Software design pattern reference",
			ImgPath: "/images/design-patterns.jpg",
			Status:  "available",
		},
	}

	// Insert books with duplicate handling
	for _, book := range books {
		var existingBook model.Book
		if err := db.Where("id = ?", book.ID).First(&existingBook).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&book).Error; err != nil {
					log.Printf("Error inserting book %d: %v", book.ID, err)
				}
			}
		} else {
			log.Printf("Book %d already exists, skipping", book.ID)
		}
	}

	log.Println("✓ Books and tasks seeded successfully")
	return nil
}