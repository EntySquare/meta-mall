package model

import (
	"gorm.io/gorm"
	"time"
)

// ContractFlow struct
type ContractFlow struct {
	gorm.Model
	AccountId   uint
	ContractId  uint
	Contract    Contract
	ReleaseNum  float64 //已领取数量
	ReleaseDate *time.Time
	Flag        string // 启用标志(0-未释放 1-已释放)
}

func (c *ContractFlow) GetById(db *gorm.DB) error {
	return db.First(&c, c.ID).Error
}
func (c *ContractFlow) GetByAccountIdAndReleaseDate(db *gorm.DB) ([]ContractFlow, error) {
	flowList := make([]ContractFlow, 0)
	db.Model(&c).Where("account_id = ? and release_date = ? ", c.AccountId, c.ReleaseDate).Find(&flowList)
	return flowList, nil
}
func (c *ContractFlow) GetByContractId(db *gorm.DB) ([]ContractFlow, error) {
	flowList := make([]ContractFlow, 0)
	db.Model(&c).Where("Contract_id = ? ", c.ContractId).Find(&flowList)
	return flowList, nil
}
func (c *ContractFlow) InsertNewContractFlow(db *gorm.DB) error {
	return db.Create(c).Error
}
