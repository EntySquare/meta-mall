package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// ContractFlow struct
type ContractFlow struct {
	gorm.Model
	UserId      uint
	ContractId  uint
	Contract    Contract
	ReleaseNum  float64 //可领取数量
	ReleaseDate *time.Time
	TokenName   string
	Flag        string // 启用标志(0-失效 1-可领取自产 2-已领取自产 3-可领取推广 4-已领取推广)
}

func (c *ContractFlow) GetById(db *gorm.DB) error {
	return db.First(&c, c.ID).Error
}
func (cf *ContractFlow) UpdateContractFlow(db *gorm.DB) error {
	return db.Model(&cf).Updates(cf).Error
}
func (c *ContractFlow) GetByContractFlowByUserId(db *gorm.DB) ([]ContractFlow, error) {
	flowList := make([]ContractFlow, 0)
	db.Model(&c).Where("user_id = ? and token_name = ?  and flag = ?", c.UserId, c.TokenName, c.Flag).Find(&flowList)
	return flowList, nil
}
func (c *ContractFlow) GetByContractId(db *gorm.DB) ([]ContractFlow, error) {
	flowList := make([]ContractFlow, 0)
	db.Model(&c).Where("Contract_id = ?  ", c.ContractId).Find(&flowList)
	return flowList, nil
}
func (c *ContractFlow) InsertNewContractFlow(db *gorm.DB) error {
	return db.Create(c).Error
}
func (c *ContractFlow) GetUserReleaseBenefitByTokenName(db *gorm.DB) (float64, error) {
	var releaseBenefit sql.NullFloat64
	err := db.Model(&c).Select("sum(release_num)").Where("user_id = ? and token_name = ? and flag = ?", c.UserId, c.TokenName, c.Flag).Scan(&releaseBenefit).Error
	if err != nil {
		return 0, err
	}
	if releaseBenefit.Valid {
		return releaseBenefit.Float64, nil
	} else {
		return 0, err
	}
}
