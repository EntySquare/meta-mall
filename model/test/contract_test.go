package test

import (
	"fmt"
	"gorm.io/gorm"
	"meta-mall/database"
	"meta-mall/model"
	"testing"
	"time"
)

func TestInsertContract(t *testing.T) {
	database.ConnectDB()
	tt := time.Now()
	user := model.User{}
	user.ID = 2
	err := user.GetById(database.DB)
	if err != nil {
		return
	}
	err = database.DB.Create(&model.Contract{
		Model:              gorm.Model{},
		NFTName:            "testNft",
		NftId:              1,
		TokenName:          "usdt",
		Hash:               "0x4513bsda24h2r1574123512",
		AccumulatedBenefit: 0,
		Power:              0,
		BenefitLimit:       0,
		StartTime:          &tt,
		ExpireTime:         &tt,
		Flag:               "2",
		OwnerId:            2,
		Owner:              user,
	}).Error
	err = database.DB.Create(&model.Contract{
		Model:              gorm.Model{},
		NFTName:            "testNft",
		NftId:              3,
		TokenName:          "usdt",
		Hash:               "0x4513bsda24h2r1574123512",
		AccumulatedBenefit: 0,
		Power:              0,
		BenefitLimit:       0,
		StartTime:          &tt,
		ExpireTime:         &tt,
		Flag:               "2",
		OwnerId:            2,
		Owner:              user,
	}).Error
	user.ID = 4
	err = user.GetById(database.DB)
	if err != nil {
		return
	}
	err = database.DB.Create(&model.Contract{
		Model:              gorm.Model{},
		NFTName:            "testNft",
		NftId:              2,
		TokenName:          "usdt",
		Hash:               "0x4513bsda24h2r1574123512",
		AccumulatedBenefit: 0,
		Power:              0,
		BenefitLimit:       0,
		StartTime:          &tt,
		ExpireTime:         &tt,
		Flag:               "2",
		OwnerId:            4,
		Owner:              user,
	}).Error

	fmt.Println(err)
}
func TestInsertContractFlow(t *testing.T) {
	database.ConnectDB()
	tt := time.Now()
	c := model.Contract{}
	c.ID = 11
	err := c.GetById(database.DB)
	if err != nil {
		return
	}
	err = database.DB.Create(&model.ContractFlow{
		UserId:      2,
		ContractId:  1,
		Contract:    c,
		ReleaseNum:  10,
		ReleaseDate: &tt,
		TokenName:   "usdt",
		Flag:        "1",
	}).Error
	err = database.DB.Create(&model.ContractFlow{
		UserId:      2,
		ContractId:  1,
		Contract:    c,
		ReleaseNum:  10,
		ReleaseDate: &tt,
		TokenName:   "usdt",
		Flag:        "2",
	}).Error
	fmt.Println(err)
}
