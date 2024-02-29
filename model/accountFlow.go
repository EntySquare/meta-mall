package model

import (
	"gorm.io/gorm"
	"time"
)

// AccountFlow struct
type AccountFlow struct {
	gorm.Model
	AccountId       uint
	Account         Account
	Num             float64
	Chain           string
	Address         string
	Hash            string
	AskForTime      *time.Time //开始申请时间
	AchieveTime     *time.Time //实现时间
	TransactionType string     // 交易类型(1-充值 2-提现)
	Flag            string     // 启用标志(1-启用 0-停用)
}

func (a *AccountFlow) GetById(db *gorm.DB) error {
	return db.First(&a, a.ID).Error
}
func (a *AccountFlow) GetByAccountId(db *gorm.DB) ([]AccountFlow, error) {
	flowList := make([]AccountFlow, 0)
	db.Model(&a).Where("account_id = ? ", a.AccountId).Order("ask_for_time").Find(&flowList)
	return flowList, nil
}
func (a *AccountFlow) InsertNewAccountFlow(db *gorm.DB) error {
	return db.Create(a).Error
}
