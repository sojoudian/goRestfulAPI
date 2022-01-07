package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sojoudian/goRestfulAPI/internal/comment"
)

// MigrateDB - migrates our database and tables
func MigrateDB(db, *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment()); result.Error != nil {
		return result.Error
	}
	return nil
}
