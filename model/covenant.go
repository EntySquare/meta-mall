package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// Covenant struct
type Covenant struct {
	gorm.Model
	NFTName            string     //产品名字
	PledgeId           string     //质押ID
	ChainName          string     //公链名称
	Duration           string     //质押期限
	Hash               string     //交易哈希
	InterestRate       float64    //日利率
	AccumulatedBenefit float64    //累计收益
	PledgeFee          float64    //质押费用
	ReleaseFee         float64    //释放费用
	StartTime          *time.Time //开始执行时间
	ExpireTime         *time.Time //结束时间
	NFTReleaseTime     *time.Time //NFT释放时间
	Flag               string     // 启用标志(1-质押中 2-已完成 0-取消中)

	ChainUnix int64 //链上质押时间
	OwnerId   uint  //用户id
	Owner     User  //用户
}

func NewCovenant(id int64) Covenant {
	return Covenant{Model: gorm.Model{ID: uint(id)}}
}
func NewCovenant2(id uint) Covenant {
	return Covenant{Model: gorm.Model{ID: id}}
}

func (c *Covenant) GetById(db *gorm.DB) error {
	return db.First(&c, c.ID).Error
}

func (c *Covenant) UpdateCovenant(db *gorm.DB) error {
	return db.Model(&c).Updates(c).Error
}
func (c *Covenant) InsertNewCovenant(db *gorm.DB) error {
	return db.Create(c).Error
}
func (c *Covenant) GetByHash(db *gorm.DB) error {
	return db.Model(&c).Where("hash = ? ", c.Hash).Take(&c).Error
}

// SelectMyCovenant
//
//	@Description:
//	@param db
//	@return userId
//	@return err

func (c *Covenant) SelectMyCovenant(db *gorm.DB) (cs []Covenant, err error) {
	cs = make([]Covenant, 0)
	err = db.Model(&Covenant{}).Where("owner_id = ?", c.OwnerId).Find(&cs).Error
	return cs, err
}

func (c *Covenant) GetUserAccumulatedBenefit(db *gorm.DB) (float64, error) {
	var accumulatedBenefit sql.NullFloat64
	err := db.Model(&c).Select("sum(release_fee)").Where("owner_id = ? ", c.OwnerId).Scan(&accumulatedBenefit).Error
	if err != nil {
		return 0, err
	}
	if accumulatedBenefit.Valid {
		return accumulatedBenefit.Float64, nil
	} else {
		return 0, err
	}
}

func SelectMyCovenantByFlag(db *gorm.DB, flag string) (cs []Covenant, err error) {
	cs = make([]Covenant, 0)
	err = db.Model(&Covenant{}).Where("flag = ?", flag).Find(&cs).Error
	return cs, err
}
