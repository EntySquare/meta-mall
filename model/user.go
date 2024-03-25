package model

import (
	"errors"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	WalletAddress     string
	RecommendId       uint
	UID               string
	Level             int64
	Power             float64
	InvestmentAddress string
	Token             string
	Account           Account `gorm:"foreignKey:UserId"`
	Flag              string  // 启用标志(1-启用 0-停用)
}

type APIUser struct {
	Phone string
}

type UserBranch struct {
	ID uint
}

func (u *User) GetById(db *gorm.DB) error {
	return db.First(&u, u.ID).Error
}

func QueryUserCount(db *gorm.DB) (uCount int64, err error) {
	if err := db.Model(&User{}).Count(&uCount).Error; err != nil {
		return 0, err
	}
	return uCount, nil
}
func (u *User) GetByWalletAddress(db *gorm.DB) error {
	return db.Model(&u).Where("wallet_address = ? ", u.WalletAddress).Take(&u).Error
}
func (u *User) GetByUid(db *gorm.DB) error {
	return db.Model(&u).Where("uid = ? ", u.UID).Take(&u).Error
}

// SelectAllUser 查询所有用户
func SelectAllUser(db *gorm.DB) (us []User, err error) {
	if err := db.Model(&User{}).Order("id").Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

// SelectAllUserID SelectAllUser 查询所有用户ID
func SelectAllUserID(db *gorm.DB) (us []uint, err error) {
	us = make([]uint, 0)
	if err := db.Model(&User{}).Select("id").Order("id").Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

// InsertNewUser 新增用户
func (u *User) InsertNewUser(db *gorm.DB) (id uint, err error) {
	result := db.Create(u)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return u.ID, nil
	}
}

// UpdateUserToken 更新用户Token
func (u *User) UpdateUserToken(db *gorm.DB) error {
	res := db.Model(&u).Where("id = ?", u.ID).Update("token", u.Token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// UserSelectIdByToken token查询用户数据 token = "HASH"
func UserSelectIdByToken(db *gorm.DB, token string) (userId int64, tokenData string, err error) {
	err = db.Table("users").
		Select("id", "token").
		Where("token LIKE ?", token+":%").
		Row().Scan(&userId, &tokenData)
	return
}

// UserRefreshToken
// @Description: 修改指定用户的token数据
// @param token 数据格式 <token_value:timestamp>
// @return err
func UserRefreshToken(db *gorm.DB, userId int64, token string) (err error) {
	res := db.Model(&User{}).Where("id = ?", userId).Update("token", token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}
