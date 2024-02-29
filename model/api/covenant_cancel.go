package api

import (
	"gorm.io/gorm"

	"meta-mall/model"
	"strconv"
	"time"
)

func CovenantCycle(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		//dateStr := time.Now().Format("2006-01-02")
		//启用标志(1-质押中 2-已完成 0-取消中)
		//质押到期 释放nft
		list1, err := model.SelectMyCovenantByFlag(tx, "1")

		//取消到期 释放nft
		list2, err := model.SelectMyCovenantByFlag(tx, "0")
		if err != nil {
			panic(err)
		}
		list := append(list1, list2...)
		for _, v := range list {
			if v.ExpireTime.Unix() < time.Now().Unix() {
				err = tx.Model(&model.Covenant{}).
					Where("id = ?", v.ID).
					Updates(map[string]interface{}{"flag": gorm.Expr("?", "2")}).Error
				if err != nil {
					return err
				}
				_, err := strconv.ParseInt(v.PledgeId, 10, 64)
				if err != nil {
					return err
				}
				var u = model.User{}
				u.ID = v.OwnerId
				err = u.GetById(tx)
				if err != nil {
					return err
				}
				//address := common.HexToAddress(u.WalletAddress)
				//contracts.WithdrawNft(address, big.NewInt(parseInt), v.ChainName)
			}
		}

		return nil
	})
	if err != nil {
		if err.Error() != "ok" {
			panic(err)
		}
	}
}
