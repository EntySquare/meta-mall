package test

import (
	"fmt"
	"meta-mall/database"
	"meta-mall/model"
	"meta-mall/model/api"
	"testing"
)

func TestInsertUser(t *testing.T) {
	database.ConnectDB()

	err := database.DB.Create(&model.User{
		WalletAddress:     "",
		RecommendId:       4,
		UID:               "150v1d391",
		Level:             0,
		Power:             100,
		InvestmentAddress: "http://localhost:4001/150v1d391",
		Token:             "UEWhAkykkIRuKhD39arTw1rv4ClL1D391CY5dO8JEvdJMrNRRjuOntAuqCH2jz4:1710249311",
		Account:           model.Account{},
		Flag:              "",
	}).Error
	fmt.Println(err)
}
func TestRunPassword(t *testing.T) {
	input := "123456"
	new := api.Get256Pw(input)
	println(new)
}
func TestInsertManager(t *testing.T) {
	database.ConnectDB()

	err := database.DB.Create(&model.Manager{
		UserName: "testUser",
		Password: "150v1d391",
		Token:    "UTGhAkykkIRuD78tY43Tw1rv4ClL1D391CIK3H5EvdJMrNRRjuOmgSdsCR414M:1710249311",
		Flag:     "1",
	}).Error
	fmt.Println(err)
}
