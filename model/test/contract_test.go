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
