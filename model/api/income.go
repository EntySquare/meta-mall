package api

import (
	"errors"
	"gorm.io/gorm"
	"meta-mall/model"
	"time"
)

var maps map[string]string

// IncomeRunP 收益跑p
func IncomeRunP(db *gorm.DB) {
	//更新质押记录数据

	//跑p查询全部质押记录
	err := db.Transaction(func(tx *gorm.DB) error {
		dateStr := time.Now().Format("2006-01-02")
		err := tx.Create(&model.PLog{DateStr: dateStr}).Error
		if err != nil {
			return errors.New("ok")
		}
		//contract := model.Contract{}
		//ap, err := contract.GetAllPowers(db)
		//if err != nil {
		//	return err
		//}
		list, err := model.SelectContractByFlag(tx, "2")
		if err != nil {
			panic(err)
		}
		for _, v := range list {
			mp := v.Power
			var rate float64
			if mp > 5 {
				rate = 1.5
			} else if mp < 50 {
				rate = 2
			} else if mp > 100 {
				rate = 3
			} else {
				rate = 1
			}
			println(rate)
			//区块未确认 跳过
			//if v.ChainUnix == 0 {
			//	continue
			//}
			////reward := 1.001 //奖励数量
			//nftId, err := strconv.Atoi(v.PledgeId)
			//if err != nil {
			//	return err
			//}
			//_, reward, _ := GetInterestRate(nftId, tx)
			////AccumulatedBenefit float64    //累计收益
			////PledgeFee          float64    //质押费用
			////ReleaseFee         float64    //释放费用
			//err = tx.Model(&model.Covenant{}).
			//	Where("id = ?", v.ID).
			//	Updates(map[string]interface{}{"accumulated_benefit": gorm.Expr("accumulated_benefit + ?", reward),
			//		"release_fee": gorm.Expr("release_fee + ?", reward)}).Error
			//if err != nil {
			//	return err
			//}
			//
			//err = tx.Model(&model.Account{}).
			//	Where("id = ?", v.OwnerId).
			//	Update("balance", gorm.Expr("balance + ?", reward)).Error
			//if err != nil {
			//	return err
			//}
			//
			//if err = tx.Create(&model.CovenantFlow{
			//	AccountId:   v.OwnerId,
			//	CovenantId:  v.ID,
			//	Num:         strconv.FormatInt(reward, 10),
			//	ReleaseDate: time.Now().Unix(),
			//	Flag:        "1",
			//}).Error; err != nil {
			//	return err
			//}
		}
		return nil
	})
	if err != nil {
		println("error")
	}
}
