package model

type PLog struct {
	DateStr string `gorm:"unique"`
	Log     string
}
