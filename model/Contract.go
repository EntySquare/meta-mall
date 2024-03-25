package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// Contract struct
type Contract struct {
	gorm.Model
	NFTName            string //产品名字
	NftId              uint   //nftID
	TokenName          string
	Hash               string     //交易哈希
	AccumulatedBenefit float64    //累计收益
	Power              float64    //算力
	BenefitLimit       float64    //收益上限
	StartTime          *time.Time //开始执行时间
	ExpireTime         *time.Time //结束时间
	Flag               string     // 启用标志(1-确认中 2-进行中 3- 已完成 0-取消中)
	OwnerId            uint       //用户id
	Owner              User       //用户
}

func NewContract(id int64) Contract {
	return Contract{Model: gorm.Model{ID: uint(id)}}
}
func NewContract2(id uint) Contract {
	return Contract{Model: gorm.Model{ID: id}}
}

func (c *Contract) GetById(db *gorm.DB) error {
	return db.First(&c, c.ID).Error
}

func (c *Contract) UpdateContract(db *gorm.DB) error {
	return db.Model(&c).Updates(c).Error
}
func (c *Contract) InsertNewContract(db *gorm.DB) error {
	return db.Create(c).Error
}
func (c *Contract) GetByHash(db *gorm.DB) error {
	return db.Model(&c).Where("hash = ? ", c.Hash).Take(&c).Error
}

// SelectMyContract
//
//	@Description:
//	@param db
//	@return userId
//	@return err

func (c *Contract) SelectMyContract(db *gorm.DB) (cs []Contract, err error) {
	cs = make([]Contract, 0)
	err = db.Model(&Contract{}).Where("owner_id = ? and flag in ('1','2')", c.OwnerId).Find(&cs).Error
	return cs, err
}

func (c *Contract) GetUserAccumulatedBenefitByTokenName(db *gorm.DB) (float64, error) {
	var accumulatedBenefit sql.NullFloat64
	err := db.Model(&c).Select("sum(accumulated_benefit)").Where("owner_id = ? and token_name = ? and flag = ?", c.OwnerId, c.TokenName, "2").Scan(&accumulatedBenefit).Error
	if err != nil {
		return 0, err
	}
	if accumulatedBenefit.Valid {
		return accumulatedBenefit.Float64, nil
	} else {
		return 0, err
	}
}
func (c *Contract) GetAllBenefitLimitByTokenName(db *gorm.DB) (float64, error) {
	var bf sql.NullFloat64
	err := db.Model(&c).Select("sum(benefit_limit)").Where("flag = ? and token_name = ?", c.Flag, c.TokenName).Scan(&bf).Error
	if err != nil {
		return 0, err
	}
	if bf.Valid {
		return bf.Float64, nil
	} else {
		return 0, err
	}
}
func SelectContractByFlag(db *gorm.DB, flag string) (cs []Contract, err error) {
	cs = make([]Contract, 0)
	err = db.Model(&Contract{}).Where("flag = ?", flag).Find(&cs).Error
	return cs, err
}
func (c *Contract) GetAllPowers(db *gorm.DB) (float64, error) {
	var accumulatedBenefit sql.NullFloat64
	err := db.Model(&c).Select("sum(power)").Where("flag = ?", "2").Scan(&accumulatedBenefit).Error
	if err != nil {
		return 0, err
	}
	if accumulatedBenefit.Valid {
		return accumulatedBenefit.Float64, nil
	} else {
		return 0, err
	}
}
