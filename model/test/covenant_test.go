package test

import (
	"fmt"
	"gorm.io/gorm"
	"meta-mall/database"
	"meta-mall/model"
	"testing"
	"time"
)

func TestInsertCovenant(t *testing.T) {
	database.ConnectDB()
	tt := time.Now()
	err := database.DB.Create(&model.Covenant{
		Model:              gorm.Model{},
		NFTName:            "White Tiger",
		PledgeId:           "dadxz",
		ChainName:          "Polygon",
		Duration:           "7天",
		Hash:               "0x63c48fe0c1f4b60f6ae90b86ea91051b06fb8b371068db84c0ad68a54e9a466c",
		InterestRate:       0.6,
		AccumulatedBenefit: 2374,
		PledgeFee:          183,
		ReleaseFee:         324,
		StartTime:          &tt,
		ExpireTime:         &tt,
		NFTReleaseTime:     &tt,
		Flag:               "1",
		OwnerId:            1,
		Owner:              model.User{},
	}).Error
	err = database.DB.Create(&model.Covenant{
		Model:              gorm.Model{},
		NFTName:            "White Tiger",
		PledgeId:           "dadaxxz",
		ChainName:          "Polygon",
		Duration:           "7天",
		Hash:               "0x63c48fe0c1f4b60f6ae90b86ea91051b06fb8b371068db84c0ad68a54e9a466c",
		InterestRate:       0.6,
		AccumulatedBenefit: 2274,
		PledgeFee:          123,
		ReleaseFee:         324,
		StartTime:          &tt,
		ExpireTime:         &tt,
		NFTReleaseTime:     &tt,
		Flag:               "2",
		OwnerId:            1,
		Owner:              model.User{},
	}).Error
	err = database.DB.Create(&model.Covenant{
		Model:              gorm.Model{},
		NFTName:            "White Tiger",
		PledgeId:           "dadaxxz",
		ChainName:          "Polygon",
		Duration:           "7天",
		Hash:               "0x63c48fe0c1f4b60f6ae90b86ea91051b06fb8b371068db84c0ad68a54e9a466c",
		InterestRate:       0.6,
		AccumulatedBenefit: 2374,
		PledgeFee:          113,
		ReleaseFee:         314,
		StartTime:          &tt,
		ExpireTime:         &tt,
		NFTReleaseTime:     &tt,
		Flag:               "1",
		OwnerId:            1,
		Owner:              model.User{},
	}).Error
	fmt.Println(err)
}
func TestInsertCovenantFlow(t *testing.T) {
	database.ConnectDB()
	tt := time.Now()
	err := database.DB.Create(&model.CovenantFlow{
		AccountId:   2,
		CovenantId:  29,
		Num:         "100",
		ReleaseDate: tt.Unix(),
		Flag:        "1",
	}).Error
	fmt.Println(err)
}
