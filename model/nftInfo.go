package model

import (
	"gorm.io/gorm"
	"strconv"
)

// NftInfo struct
type NftInfo struct {
	gorm.Model
	Name            string
	NftNumber       int64
	DayRate         string
	TokenId         int64
	ContractAddress string
	OwnerAddress    string
	ChainName       string //公链名 //Polygon
	Price           float64
	Power           float64 //T
	TypeNum         int64   //nft种类
	ImgUrl          string
	Flag            string // // 启用标志(1-可购买 2-预定中 0-已售出)
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
func (ni *NftInfo) GetAllNftInfoByFlag(flag int, db *gorm.DB) (ns []NftInfo, err error) {
	ns = make([]NftInfo, 0)
	err = db.Model(&NftInfo{}).Where("flag = ? or flag = ?", strconv.Itoa(flag), strconv.Itoa(flag+1)).Find(&ns).Error
	return ns, err
}
