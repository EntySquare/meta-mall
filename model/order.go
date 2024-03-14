package model

import (
	"gorm.io/gorm"
)

// Order struct
type Order struct {
	gorm.Model
	NftId        uint    //产品名字
	Price        float64 //购买价格
	BuyerId      uint    //`购买者ID`
	BuyerAddress string  //购买地址
	Flag         string  // 启用标志(1-处理中 2-已完成 0-取消)

	Buyer User //用户
}

func NewOrder(id int64) Order {
	return Order{Model: gorm.Model{ID: uint(id)}}
}
func NewOrder2(id uint) Order {
	return Order{Model: gorm.Model{ID: id}}
}

func (o *Order) GetById(db *gorm.DB) error {
	return db.First(&o, o.ID).Error
}
func (o *Order) GetByNftId(db *gorm.DB) error {
	return db.Model(&o).Where("nft_id = ? ", o.NftId).Take(&o).Error
}

func (o *Order) UpdateOrder(db *gorm.DB) error {
	return db.Model(&o).Updates(o).Error
}
func (o *Order) InsertNewOrder(db *gorm.DB) error {
	return db.Create(o).Error
}

// SelectMyOrder
//
//	@Description:
//	@param db
//	@return userId
//	@return err

func (o *Order) SelectMyOrder(db *gorm.DB) (os []Order, err error) {
	os = make([]Order, 0)
	err = db.Model(&Order{}).Where("buyer_id = ?", o.BuyerId).Find(&os).Error
	return os, err
}
