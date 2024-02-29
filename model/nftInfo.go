package model

import (
	"gorm.io/gorm"
)

// NftInfo struct
type NftInfo struct {
	gorm.Model
	Name      string
	DayRate   string
	ChainName string //公链名 //Polygon
	TypeNum   int64  //nft种类
	ImgUrl    string
	Flag      string // // 启用标志(1-充值 2-提现 0-取消中)
}

func NewNftInfo(id int64) NftInfo {
	return NftInfo{Model: gorm.Model{ID: uint(id)}}
}

func (ni *NftInfo) GetById(db *gorm.DB) error {
	return db.First(&ni, ni.ID).Error
}

func (ni *NftInfo) UpdateNftInfo(db *gorm.DB) error {
	return db.Model(&ni).Updates(ni).Error
}
func (ni *NftInfo) InsertNewNftInfo(db *gorm.DB) error {
	return db.Create(ni).Error
}
func (ni *NftInfo) GetByTypeNum(db *gorm.DB) error {
	return db.Model(&ni).Where("type_num = ? ", ni.TypeNum).First(&ni).Error
}
func (ni *NftInfo) GetAllNftInfoByFlag(db *gorm.DB) (ns []NftInfo, err error) {
	ns = make([]NftInfo, 0)
	err = db.Model(&NftInfo{}).Where("flag = ?", "1").Find(&ns).Error
	return ns, err
}
