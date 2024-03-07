package test

import (
	"fmt"
	"gorm.io/gorm"
	"meta-mall/database"
	"meta-mall/model"
	"testing"
)

func TestInsertNft(t *testing.T) {
	database.ConnectDB()

	err := database.DB.Create(&model.NftInfo{
		Model:           gorm.Model{},
		Name:            "testNft",
		NftNumber:       1,
		DayRate:         "0",
		TokenId:         0,
		ContractAddress: "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		OwnerAddress:    "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		ChainName:       "bsc",
		Price:           1688,
		TypeNum:         1,
		ImgUrl:          "测试地址",
		Flag:            "1",
	}).Error
	err = database.DB.Create(&model.NftInfo{
		Model:           gorm.Model{},
		Name:            "testNft",
		NftNumber:       3,
		DayRate:         "0",
		TokenId:         1,
		ContractAddress: "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		OwnerAddress:    "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		ChainName:       "bsc",
		Price:           1688,
		TypeNum:         1,
		ImgUrl:          "测试地址",
		Flag:            "1",
	}).Error
	err = database.DB.Create(&model.NftInfo{
		Model:           gorm.Model{},
		Name:            "testNft",
		NftNumber:       2,
		DayRate:         "0",
		TokenId:         2,
		ContractAddress: "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		OwnerAddress:    "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		ChainName:       "bsc",
		Price:           1688,
		TypeNum:         1,
		ImgUrl:          "测试地址",
		Flag:            "1",
	}).Error
	err = database.DB.Create(&model.NftInfo{
		Model:           gorm.Model{},
		Name:            "testNft1",
		NftNumber:       1,
		DayRate:         "0",
		TokenId:         0,
		ContractAddress: "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		OwnerAddress:    "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		ChainName:       "bsc",
		Price:           1688,
		TypeNum:         1,
		ImgUrl:          "测试地址",
		Flag:            "1",
	}).Error
	err = database.DB.Create(&model.NftInfo{
		Model:           gorm.Model{},
		Name:            "testNft1",
		NftNumber:       3,
		DayRate:         "0",
		TokenId:         1,
		ContractAddress: "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		OwnerAddress:    "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		ChainName:       "bsc",
		Price:           1688,
		TypeNum:         1,
		ImgUrl:          "测试地址",
		Flag:            "1",
	}).Error
	err = database.DB.Create(&model.NftInfo{
		Model:           gorm.Model{},
		Name:            "testNft1",
		NftNumber:       2,
		DayRate:         "0",
		TokenId:         2,
		ContractAddress: "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		OwnerAddress:    "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		ChainName:       "bsc",
		Price:           88,
		TypeNum:         1,
		ImgUrl:          "测试地址",
		Flag:            "1",
	}).Error
	err = database.DB.Create(&model.NftInfo{
		Model:           gorm.Model{},
		Name:            "testNft2",
		NftNumber:       1,
		DayRate:         "0",
		TokenId:         0,
		ContractAddress: "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		OwnerAddress:    "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		ChainName:       "bsc",
		Price:           188,
		TypeNum:         1,
		ImgUrl:          "测试地址",
		Flag:            "1",
	}).Error
	err = database.DB.Create(&model.NftInfo{
		Model:           gorm.Model{},
		Name:            "testNft2",
		NftNumber:       3,
		DayRate:         "0",
		TokenId:         1,
		ContractAddress: "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		OwnerAddress:    "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		ChainName:       "bsc",
		Price:           188,
		TypeNum:         1,
		ImgUrl:          "测试地址",
		Flag:            "1",
	}).Error
	err = database.DB.Create(&model.NftInfo{
		Model:           gorm.Model{},
		Name:            "testNft2",
		NftNumber:       2,
		DayRate:         "0",
		TokenId:         2,
		ContractAddress: "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		OwnerAddress:    "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		ChainName:       "bsc",
		Price:           188,
		TypeNum:         1,
		ImgUrl:          "测试地址",
		Flag:            "1",
	}).Error
	fmt.Println(err)
}
