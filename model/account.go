package model

import (
	"gorm.io/gorm"
)

// Account struct
type Account struct {
	gorm.Model
	UserId            uint
	USDTBalance       float64
	USDTFrozenBalance float64
	UNCBalance        float64
	UNCFrozenBalance  float64
	Flag              string // 启用标志(1-启用 0-停用)
}

func (ac *Account) GetById(db *gorm.DB) error {
	return db.First(&ac, ac.ID).Error
}
func (ac *Account) GetByUserId(db *gorm.DB) error {
	return db.Model(&ac).Where("user_id = ? ", ac.UserId).Take(&ac).Error
}
func (ac *Account) UpdateAccount(db *gorm.DB) error {
	return db.Model(&ac).Updates(ac).Error
}
func (ac *Account) InsertNewAccount(db *gorm.DB) error {
	return db.Create(ac).Error
}
