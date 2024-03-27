package api

import (
	"errors"
	"gorm.io/gorm"
	"math"
	"meta-mall/model"
	"time"
)

var maps map[string]string
var UncAmount float64
var MetaAmount float64

// IncomeRunP 收益跑p
func IncomeRunP(db *gorm.DB) error {
	//更新质押记录数据

	//跑p查询全部质押记录
	err := db.Transaction(func(tx *gorm.DB) error {
		dateStr := time.Now().Format("2006-01-02")
		err := tx.Create(&model.PLog{DateStr: dateStr}).Error
		if err != nil {
			return errors.New("ok")
		}
		var ap float64
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
			ap += mp * rate
		}
		for _, v := range list {
			tt := time.Now()
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
			mb := 0.0
			if ap != 0 {
				mb += round((UncAmount*mp*rate*6)/(ap*10), 4)
			}
			if GetBranchAccumulatedPower(0) != 0 {
				mb += round((GetBranchAccumulatedPower(v.OwnerId)*4)/(GetBranchAccumulatedPower(0)*10), 4)
			}

			cf := model.ContractFlow{
				UserId:      v.OwnerId,
				ContractId:  v.ID,
				Contract:    v,
				ReleaseNum:  mb,
				ReleaseDate: &tt,
				TokenName:   "unc",
				Flag:        "1",
			}
			mb2 := 0.0
			if ap != 0 {
				mb2 += round((MetaAmount*mp*rate*6)/(ap*10), 4)
			}
			if GetBranchAccumulatedPower(0) != 0 {
				mb2 += round((GetBranchAccumulatedPower(v.OwnerId)*4)/(GetBranchAccumulatedPower(0)*10), 4)
			}
			cf2 := model.ContractFlow{
				UserId:      v.OwnerId,
				ContractId:  v.ID,
				Contract:    v,
				ReleaseNum:  mb2,
				ReleaseDate: &tt,
				TokenName:   "meta",
				Flag:        "1",
			}
			err := cf.InsertNewContractFlow(tx)
			if err != nil {
				return err
			}
			err = cf2.InsertNewContractFlow(tx)
			if err != nil {
				return err
			}
			account := model.Account{}
			account.UserId = v.OwnerId
			err = account.GetByUserId(tx)
			if err != nil {
				return err
			}
			account.UNCBalance = account.UNCBalance + mb2
			account.METABalance = account.METABalance + mb2
			err = account.UpdateAccount(tx)
			if err != nil {
				return err
			}
			contract := model.Contract{}
			contract.ID = v.ID
			err = contract.GetById(tx)
			if err != nil {
				return err
			}
			contract.AccumulatedBenefit = contract.AccumulatedBenefit + mb
			err = contract.UpdateContract(tx)
			if err != nil {
				return err
			}
		}
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

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
func round(num float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(num*shift+.5) / shift
}
