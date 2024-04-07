package model

import (
	"errors"
	"gorm.io/gorm"
)

// Manager struct
type Manager struct {
	gorm.Model
	UserName string
	Password string
	Token    string
	Flag     string // 启用标志(1-启用 0-停用)
}

type APIManager struct {
	Phone string
}

type ManagerBranch struct {
	ID uint
}

func (u *Manager) GetById(db *gorm.DB) error {
	return db.First(&u, u.ID).Error
}
func (u *Manager) GetByUsername(db *gorm.DB) error {
	return db.Model(&u).Where("user_name = ? ", u.UserName).Take(&u).Error
}

// InsertNewManager 新增用户
func (u *Manager) InsertNewManager(db *gorm.DB) (id uint, err error) {
	result := db.Create(u)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return u.ID, nil
	}
}
func (m *Manager) UpdateManager(db *gorm.DB) error {
	return db.Model(&m).Updates(m).Error
}

// ManagerrSelectIdByToken token查询用户数据 token = "HASH"
func ManagerSelectIdByToken(db *gorm.DB, token string) (username string, tokenData string, err error) {
	err = db.Table("managers").
		Select("user_name", "token").
		Where("token LIKE ?", token+":%").
		Row().Scan(&username, &tokenData)
	return
}

// UpdateManagerToken 更新用户Token
func (u *Manager) UpdateManagerToken(db *gorm.DB) error {
	res := db.Model(&u).Where("id = ?", u.ID).Update("token", u.Token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}
func ManagerRefreshToken(db *gorm.DB, username string, token string) (err error) {
	res := db.Model(&Manager{}).Where("user_name = ?", username).Update("token", token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}
func (m *Manager) SelectApplyList(db *gorm.DB) (as []Manager, err error) {
	as = make([]Manager, 0)
	err = db.Model(&Manager{}).Where("flag in ('0','2')").Find(&as).Error
	return as, err
}
