package model

// import "gorm.io/gorm"

type Student struct {
	// gorm.Model        // Includes fields like ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"column:name;primaryKey"` // Unique and cannot be null
	Address    string `gorm:"column:address"`
	City       string `gorm:"column:city"`
	Country    string `gorm:"column:country"`
	Pincode    string `gorm:"column:pincode"`
	SATScore   int    `gorm:"column:sat_score;not null"`
	Passed     string `gorm:"column:passed;not null"`
}
