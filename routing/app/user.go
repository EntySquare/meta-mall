package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"meta-mall/config"
	"meta-mall/database"
	"meta-mall/model"
	"meta-mall/model/api"
	"meta-mall/pkg"
	"meta-mall/routing/types"
	"strconv"
	"strings"
	"time"
)

// LoginAndRegister 登录注册
func LoginAndRegister(c *fiber.Ctx) error {
	fmt.Println("/LoginAndRegister api...")
	reqParams := types.LoginAndRegisterReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	user := model.User{
		Flag: "1",
	}
	user.WalletAddress = reqParams.WalletAddress
	returnT := ""
	err = user.GetByWalletAddress(database.DB)
	//if err != nil {
	//	if !errors.Is(err, gorm.ErrRecordNotFound) {
	//		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "ErrRecordNotFound", ""))
	//	}
	//}
	recommendUser := model.User{}
	fmt.Println("推荐码：", reqParams.Code)
	database.DB.Model(&model.User{}).Where("uid = ?", reqParams.Code).Find(&recommendUser)

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		if err != nil {
			if !strings.Contains(err.Error(), "record not found") {
				return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "user get by addresss error", ""))
			}
			user.Level = 0
			user.PledgeCount = 0
			user.UID = pkg.RandomCodes(6) + user.WalletAddress[6:9]
			user.InvestmentAddress = "https://metagalaxylands.com/" + user.UID
			//user.InvestmentAddress = "http://localhost:4001/" + user.UID
			returnT = pkg.RandomString(64)
			user.Token = returnT + ":" + strconv.FormatInt(time.Now().Unix(), 10)
			user.RecommendId = recommendUser.ID
			api.AddNewBranch(user.RecommendId, user.ID)
			err = user.InsertNewUser(tx)
			if err != nil {
				return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "注册失败"))
			}

			err := user.GetByWalletAddress(tx)
			if err != nil {
				return err
			}
			var acc = model.Account{
				UserId:        user.ID,
				Balance:       0,
				FrozenBalance: 0,
				Flag:          "1",
			}
			err = acc.InsertNewAccount(tx)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, err.Error(), "注册失败"))
	}
	returnT = strings.Split(user.Token, ":")[0]
	c.Locals(config.LOCAL_TOKEN, returnT)
	return c.JSON(pkg.SuccessResponse(returnT))
}
func MyInvestment(c *fiber.Ctx) error {
	//userId := c.Locals("user_id")
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	user := model.User{}
	user.ID = userId
	err := user.GetById(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询用户失败"))
	}
	data := types.MyInvestmentResp{
		InvestmentUsers: make([]types.InvestmentUserInfo, 0),
	}
	data.UID = user.UID
	data.InvestmentAddress = config.Config("WEB_URL") + "/&code=" + user.UID
	data.Level = user.Level
	data.InvestmentCount = int64(len(api.UserTree[userId].Branch))
	data.AccumulatedPledgeCount = api.GetBranchAccumulatedPledgeCount(userId)
	for _, branch := range api.UserTree[userId].Branch {
		in := types.InvestmentUserInfo{}
		in.UID = api.UserTree[branch].UID
		in.Level = api.UserTree[branch].Level
		in.PledgeCount = api.UserTree[branch].PledgeCount
		data.InvestmentUsers = append(data.InvestmentUsers, in)

	}
	return c.JSON(pkg.SuccessResponse(data))
}
func MyNgt(c *fiber.Ctx) error {
	//userId := c.Locals("user_id")
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	acc := model.Account{}
	acc.UserId = userId
	err := acc.GetByUserId(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询账户失败"))
	}
	benefit, err := getBenefit(acc)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询收益失败"))
	}
	data := types.MyNgtResp{
		BenefitInfo:  benefit,
		Transactions: make([]types.TransactionInfo, 0),
	}
	af := model.AccountFlow{}
	af.AccountId = acc.ID
	afs, err := af.GetByAccountId(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询账户交易失败"))
	}
	for _, accflow := range afs {
		txs := model.Transactions{}
		txs.Hash = accflow.Hash
		err := txs.GetByHash(database.DB)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询账户交易失败"))
		}
		//AchieveTime := nil
		in := types.TransactionInfo{
			Num:             accflow.Num,
			Chain:           accflow.Chain,
			Address:         accflow.Address,
			Hash:            accflow.Hash,
			AskForTime:      accflow.AskForTime.Unix(),
			TransactionType: accflow.TransactionType,
			Status:          txs.Status,
		}
		if accflow.AchieveTime != nil {
			in.AchieveTime = accflow.AchieveTime.Unix()
		}
		data.Transactions = append(data.Transactions, in)

	}
	return c.JSON(pkg.SuccessResponse(data))
}
func MyCovenantFlow(c *fiber.Ctx) error {
	//userId := c.Locals("user_id")
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	acc := model.Account{}
	acc.UserId = userId
	err := acc.GetByUserId(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询账户失败"))
	}
	benefit, err := getBenefit(acc)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询收益失败"))
	}
	data := types.MyCovenantFlowResp{
		BenefitInfo: benefit,
		Covenants:   make([]types.CovenantInfo, 0),
	}
	co := model.Covenant{}
	co.OwnerId = acc.UserId
	cos, err := co.SelectMyCovenant(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询账户交易失败"))
	}
	for _, coi := range cos {
		in := types.CovenantInfo{
			CovenantId:         coi.ID,
			NFTName:            coi.NFTName,
			PledgeId:           coi.PledgeId,
			ChainName:          coi.ChainName,
			Duration:           coi.Duration,
			Hash:               coi.Hash,
			InterestRate:       coi.InterestRate,
			AccumulatedBenefit: coi.AccumulatedBenefit,
			PledgeFee:          coi.PledgeFee,
			ReleaseFee:         coi.ReleaseFee,
			StartTime:          coi.StartTime.Unix(),
			ExpireTime:         coi.ExpireTime.Unix(),
			NFTReleaseTime:     coi.NFTReleaseTime.Unix(),
			Flag:               coi.Flag,
		}
		data.Covenants = append(data.Covenants, in)

	}
	return c.JSON(pkg.SuccessResponse(data))
}
func getLastDay() int64 {
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -1)
	y, m, d := oldTime.Date()
	date := int64(y*10000 + int(m)*100 + d)
	return date
}
func getBenefit(acc model.Account) (types.BenefitInfo, error) {
	data := types.BenefitInfo{
		LastDayBenefit: 0.0,
	}
	data.Balance = acc.Balance
	co := model.Covenant{}
	co.OwnerId = acc.UserId
	benefits, err := co.GetUserAccumulatedBenefit(database.DB)
	if err != nil {
		return data, err
	}
	data.AccumulatedBenefit = benefits
	cf := model.CovenantFlow{}
	cf.AccountId = acc.ID
	cf.ReleaseDate = getLastDay()
	//cfs, err := cf.GetByAccountIdAndReleaseDate(database.DB)
	if err != nil {
		return data, err
	}
	//for _, flow := range cfs {
	//	data.LastDayBenefit += flow.Num //TODO
	//}
	return data, nil
}
func GetInviteeInfo(c *fiber.Ctx) error {
	reqParams := types.InviteeInfoReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	var user = model.User{}
	user.UID = reqParams.Uid
	err = user.GetByUid(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "find uid error", ""))
	}
	data := types.InviteeInfoResp{
		Uid:         reqParams.Uid,
		Level:       user.Level,
		PledgeCount: api.UserTree[user.ID].PledgeCount,
		CreateTime:  user.CreatedAt.Unix(),
		Covenants:   make([]types.CovenantInfo, 0),
	}
	co := model.Covenant{}
	co.OwnerId = user.ID
	cos, err := co.SelectMyCovenant(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询账户交易失败"))
	}
	for _, coi := range cos {
		in := types.CovenantInfo{
			NFTName:            coi.NFTName,
			PledgeId:           coi.PledgeId,
			ChainName:          coi.ChainName,
			Duration:           coi.Duration,
			Hash:               coi.Hash,
			InterestRate:       coi.InterestRate,
			AccumulatedBenefit: coi.AccumulatedBenefit,
			PledgeFee:          coi.PledgeFee,
			ReleaseFee:         coi.ReleaseFee,
			StartTime:          coi.StartTime.Unix(),
			ExpireTime:         coi.ExpireTime.Unix(),
			NFTReleaseTime:     coi.NFTReleaseTime.Unix(),
			Flag:               coi.Flag,
		}
		data.Covenants = append(data.Covenants, in)

	}
	return c.JSON(pkg.SuccessResponse(data))
}
func GetCovenantDetail(c *fiber.Ctx) error {
	reqParams := types.CovenantDetailReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	co := model.Covenant{Hash: reqParams.Hash}
	err = co.GetByHash(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "wrong hash", ""))
	}
	cof := model.CovenantFlow{CovenantId: co.ID}
	bfs, err := cof.GetByCovenantId(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "query benefit flows error", ""))
	}
	data := types.CovenantDetailResp{}
	for _, bf := range bfs {
		in := types.CovenantDetail{
			Time: bf.CreatedAt.Unix(),
			Num:  bf.Num,
			Flag: bf.Flag,
		}
		data.List = append(data.List, in)
	}

	return c.JSON(pkg.SuccessResponse(data))
}
