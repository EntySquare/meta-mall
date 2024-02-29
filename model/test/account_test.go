package test

import (
	"fmt"
	"gorm.io/gorm"
	"meta-mall/database"
	"meta-mall/model"
	"testing"
)

func TestInsertAccount(t *testing.T) {
	database.ConnectDB()
	user := model.User{
		Model: gorm.Model{ID: 28},
	}
	err := user.GetById(database.DB)
	if err != nil {
		return
	}
	err = database.DB.Create(&model.Account{
		Model:         gorm.Model{},
		UserId:        user.ID,
		Balance:       200,
		FrozenBalance: 0,
		Flag:          "1",
	}).Error
	fmt.Println(err)
}
