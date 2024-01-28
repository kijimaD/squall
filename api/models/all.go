package models

import (
	"fmt"

	"gorm.io/gorm"
)

var All = []any{
	Entry{},
}

const defaultIncrementVal = 0

// レコードの中から最大のIDを返す
func GetMaxID(db *gorm.DB, modelVar interface{}) uint {
	var maxID uint
	err := db.Model(&modelVar).Select(fmt.Sprintf("COALESCE(MAX(id), %d)", defaultIncrementVal)).Scan(&maxID).Error
	if err != nil {
		return defaultIncrementVal
	}

	return maxID
}
