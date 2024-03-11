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

func (c *Order) GetById(db *gorm.DB) error {
	return db.First(&c, c.ID).Error
}

func (c *Order) UpdateOrder(db *gorm.DB) error {
	return db.Model(&c).Updates(c).Error
}
func (c *Order) InsertNewOrder(db *gorm.DB) error {
	return db.Create(c).Error
}

// SelectMyOrder
//
//	@Description:
//	@param db
//	@return userId
//	@return err

func (c *Order) SelectMyOrder(db *gorm.DB) (cs []Order, err error) {
	cs = make([]Order, 0)
	err = db.Model(&Order{}).Where("buyer_id = ?", c.BuyerId).Find(&cs).Error
	return cs, err
}
