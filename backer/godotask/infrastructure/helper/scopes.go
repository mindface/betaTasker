package helper

import (
  "gorm.io/gorm"
  "fmt"
)

func WithUserFilter(userID uint) func(db *gorm.DB) *gorm.DB {
  return func(db *gorm.DB) *gorm.DB {
    fmt.Printf("llll---------- %d",userID)
    if userID == 0 {
      return db
    }
    fmt.Printf("none----------")
    return db.Where("user_id = ?", userID)
  }
}

func BuildPaginationQuery(db *gorm.DB, userID uint, offset, limit int) (*gorm.DB, error) {
    q := db.Scopes(WithUserFilter(userID))
    return q, nil
}