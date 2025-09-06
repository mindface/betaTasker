package migration

import (
	"log"
	"gorm.io/gorm"
)

// RunMigrations executes all database migrations
func RunMigrations(db *gorm.DB) error {
	log.Println("Starting database migrations...")

	// Run migration 001: Create state evaluation tables
	if err := CreateStateEvaluationTables(db); err != nil {
		log.Printf("Error running migration 001: %v", err)
		return err
	}
	log.Println("Migration 001 completed: State evaluation tables created")

	log.Println("All migrations completed successfully")
	return nil
}