package test

import (
	"fmt"
	"gorm.io/gorm"
	"meta-mall/database"
	"meta-mall/model"
	"testing"
	"time"
)

func TestInsertTransactions(t *testing.T) {
	database.ConnectDB()
	tt := time.Now()
	err := database.DB.Create(&model.AccountFlow{
		Model:           gorm.Model{},
		AccountId:       1,
		Num:             100,
		Chain:           "Polygon",
		Address:         "0xc0822561B310256Aef0032e09b149Ac7cD7b5D55",
		Hash:            "0x63c48fe0c1f4b60f6ae90b86ea91051b06fb8b37x068db84c0ad68a54e9a466c",
		AskForTime:      &tt,
		AchieveTime:     &tt,
		TransactionType: "1",
		Flag:            "1",
	}).Error
	err = database.DB.Create(&model.AccountFlow{
		Model:           gorm.Model{},
		AccountId:       1,
		Num:             200,
		Chain:           "Polygon",
		Address:         "0xc0822561B310256Aef0032e09b149Ac7cD7b5D55",
		Hash:            "0x63c48fe0c1f4b60f6ae90b86ea91051b06fb8b37x068db84c0ad68a54e9a466c",
		AskForTime:      &tt,
		AchieveTime:     &tt,
		TransactionType: "1",
		Flag:            "1",
	}).Error
	err = database.DB.Create(&model.AccountFlow{
		Model:           gorm.Model{},
		AccountId:       1,
		Num:             200,
		Chain:           "Polygon",
		Address:         "0xc0822561B310256Aef0032e09b149Ac7cD7b5D55",
		Hash:            "0x63c48fe0c1f4b60f6ae90b86ea91051b06fb8b371068db84c0ad68a54e9a466c",
		AskForTime:      &tt,
		AchieveTime:     &tt,
		TransactionType: "1",
		Flag:            "1",
	}).Error
	err = database.DB.Create(&model.Transactions{
		Model:     gorm.Model{},
		Hash:      "0x63c48fe0c1f4b60f6ae90b86ea91051b06fb8b37x068db84c0ad68a54e9a466c",
		Status:    "1",
		ChainName: "Polygon",
		Flag:      "1",
	}).Error
	err = database.DB.Create(&model.Transactions{
		Model:     gorm.Model{},
		Hash:      "0x63c48fe0c1f4b60f6ae90b86ea91051b06fb8b371068db84c0ad68a54e9a466c",
		Status:    "2",
		ChainName: "Polygon",
		Flag:      "1",
	}).Error
	err = database.DB.Create(&model.Transactions{
		Model:     gorm.Model{},
		Hash:      "0x63c48fe0c1f4b60f6ae90b86ea91051b06fb8b371068db84c0ad68a54e9a466c",
		Status:    "2",
		ChainName: "Polygon",
		Flag:      "2",
	}).Error
	fmt.Println(err)
}
