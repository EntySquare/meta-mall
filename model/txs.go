package model

import (
	"gorm.io/gorm"
)

// Transactions struct
type Transactions struct {
	gorm.Model
	Hash        string `gorm:"unique"` //交易哈希
	Status      string //交易状态  0 - 未处理  1 - 未确认 2 - 已确认
	ChainName   string //公链名 //poly
	Gas         string
	FromAddress string
	Flag        string // // 启用标志(1-充值 2-提现 0-取消中)
}

func NewTransactions(id int64) Transactions {
	return Transactions{Model: gorm.Model{ID: uint(id)}}
}

func (txs *Transactions) GetById(db *gorm.DB) error {
	return db.First(&txs, txs.ID).Error
}

func (txs *Transactions) UpdateTransactions(db *gorm.DB) error {
	return db.Model(&txs).Updates(txs).Error
}
func (txs *Transactions) InsertNewTransactions(db *gorm.DB) error {
	return db.Create(txs).Error
}
func (txs *Transactions) GetByHash(db *gorm.DB) error {
	return db.Model(&txs).Where("hash = ? ", txs.Hash).First(&txs).Error
}
func (txs *Transactions) GetUntreatedTxs(db *gorm.DB) ([]Transactions, error) {
	transactionList := make([]Transactions, 0)
	db.Model(&txs).Where("flag = 1 AND status in (0,1) AND chain_name = ?", txs.ChainName).Find(&transactionList)
	return transactionList, nil
}
