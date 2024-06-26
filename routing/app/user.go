package app

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math"
	"math/big"
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
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "", ""))
		}
	}
	recommendUser := model.User{}
	fmt.Println("推荐码：", reqParams.Code)
	database.DB.Model(&model.User{}).Where("uid = ?", reqParams.Code).Find(&recommendUser)

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		if !pkg.CheckTokenValidityTime(&user.Token) {
			returnT = pkg.RandomString(64)
			user.Token = returnT + ":" + strconv.FormatInt(time.Now().Unix(), 10)
			err := user.UpdateUserToken(tx)
			if err != nil {
				return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "update token error", ""))
			}
		}
		if err != nil {
			if !strings.Contains(err.Error(), "record not found") {
				return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "user get by addresss error", ""))
			}
			user.Level = 0
			user.Power = 0
			user.UID = pkg.RandomCodes(6) + user.WalletAddress[6:9]
			//user.InvestmentAddress = "https://metagalaxylands.com/" + user.UID
			user.InvestmentAddress = "http://localhost:4001/" + user.UID
			returnT = pkg.RandomString(64)
			user.Token = returnT + ":" + strconv.FormatInt(time.Now().Unix(), 10)
			user.RecommendId = recommendUser.ID

			newId, err := user.InsertNewUser(tx)
			if err != nil {
				return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "注册失败"))
			}
			api.AddNewBranch(user.RecommendId, reqParams.WalletAddress, newId)
			err = user.GetByWalletAddress(tx)
			if err != nil {
				return err
			}
			var acc = model.Account{
				UserId:      user.ID,
				USDTBalance: 0,
				UNCBalance:  0,
				METABalance: 0,
				Flag:        "1",
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
	data.Address = user.WalletAddress
	data.InvestmentAddress = config.Config("WEB_URL") + "/?code=" + user.UID
	data.Level = user.Level
	data.Powers = api.UserTree[userId].Power
	//data.AccumulatedPledgeCount = api.GetBranchAccumulatedPower(userId)
	for _, branch := range api.UserTree[userId].Branch {
		in := types.InvestmentUserInfo{}
		in.Address = api.UserTree[branch].Address
		in.Level = api.UserTree[branch].Level
		in.Powers = api.UserTree[branch].Power
		branchUser := model.User{}
		branchUser.WalletAddress = api.UserTree[branch].Address
		err = branchUser.GetByWalletAddress(database.DB)
		if err != nil {
			return err
		}
		in.Time = branchUser.Model.CreatedAt.Unix()
		data.InvestmentUsers = append(data.InvestmentUsers, in)
	}
	return c.JSON(pkg.SuccessResponse(data))
}
func MyPromotion(c *fiber.Ctx) error {
	//userId := c.Locals("user_id")
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	user := model.User{}
	user.ID = userId
	err := user.GetById(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询用户失败"))
	}
	flow := model.ContractFlow{}
	flow.UserId = userId
	flow.Flag = "4"
	flow.TokenName = "unc"
	mrb, err := flow.GetUserReleaseBenefitByTokenName(database.DB)
	if err != nil {
		return err
	}
	flow.Flag = "3"
	mb, err := flow.GetUserReleaseBenefitByTokenName(database.DB)
	if err != nil {
		return err
	}
	data := types.MyPromotionResp{
		AllPromotionPower:           api.GetBranchAccumulatedPower(0),
		MyPromotionPower:            api.GetUserPromotionPower(userId),
		MyPromotionBenefit:          mb + mrb,
		MyAvailablePromotionBenefit: mb,
	}
	return c.JSON(pkg.SuccessResponse(data))
}
func GetAvailableBenefit(c *fiber.Ctx) error {
	fmt.Println("/GetAvailableBenefit api...")
	reqParams := types.GetAvailableBenefitReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	user := model.User{}
	user.ID = userId
	err = user.GetById(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	myFlows := model.ContractFlow{}
	myFlows.UserId = userId
	contractFlows := make([]model.ContractFlow, 0)
	if reqParams.Type == "1" {
		myFlows.TokenName = reqParams.TokenName
		myFlows.Flag = "1"
		contractFlows, err = myFlows.GetByContractFlowByUserId(database.DB)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get contract flow error", ""))
		}
	} else if reqParams.Type == "2" {
		myFlows.Flag = "3"
		myFlows.TokenName = "unc"
		contractFlows, err = myFlows.GetByContractFlowByUserId(database.DB)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get contract flow error", ""))
		}
	} else {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "type error", ""))
	}
	mab := 0.0
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		for _, flow := range contractFlows {
			mab += flow.ReleaseNum
			if reqParams.Type == "1" {
				flow.Flag = "2"
			} else if reqParams.Type == "2" {
				flow.Flag = "4"
			} else {
				return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "type error", ""))
			}
			err := flow.UpdateContractFlow(tx)
			if err != nil {
				return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "update contract flow error", ""))
			}
		}
		account := model.Account{}
		account.UserId = userId
		err = account.GetByUserId(tx)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get account  error", ""))
		}
		if reqParams.TokenName == "usdt" {
			account.USDTBalance = account.USDTBalance - mab
		} else if reqParams.TokenName == "unc" {
			account.UNCBalance = account.UNCBalance - mab
		} else {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "token name error", ""))
		}
		hash, err := api.ReleaseERC20(user.WalletAddress, benefitCalculate(mab, 18), reqParams.TokenName)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "transfer token error", ""))
		}
		tt := time.Now()
		if mab != 0 {
			flow := model.AccountFlow{
				AccountId:       account.ID,
				Account:         account,
				Num:             mab,
				Chain:           reqParams.TokenName,
				Address:         user.WalletAddress,
				Hash:            hash,
				TokenName:       reqParams.TokenName,
				AskForTime:      &tt,
				AchieveTime:     &tt,
				TransactionType: "2",
				Flag:            "1",
			}
			err = flow.InsertNewAccountFlow(tx)
			if err != nil {
				return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "insert account flow error", ""))
			}
		}

		err = account.UpdateAccount(tx)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "update account error", ""))
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "transaction error", ""))
	}
	return c.JSON(pkg.SuccessResponse(""))
}
func ApplyForMember(c *fiber.Ctx) error {
	//userId := c.Locals("user_id")
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	user := model.User{}
	user.ID = userId
	err := user.GetById(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询用户失败"))
	}
	reqParams := types.ApplyForMemberReq{}
	err = c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	manager := model.Manager{}
	manager.UserName = user.WalletAddress
	err = manager.GetByUsername(database.DB)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "", ""))
		} else {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "you already apply for membership", ""))
		}
	}
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		newManager := model.Manager{
			UserName: user.WalletAddress,
			Password: reqParams.Password,
			Flag:     "0",
		}
		_, err := newManager.InsertNewManager(tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "insert error", ""))
	}
	return c.JSON(pkg.SuccessResponse(""))
}
func getLastDay() int64 {
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -1)
	y, m, d := oldTime.Date()
	date := int64(y*10000 + int(m)*100 + d)
	return date
}
func benefitCalculate(value float64, exponent int) *big.Int {
	// 将10的指数转换为整数幂
	multiplier := math.Pow10(exponent - 6)
	var result *big.Int
	bv := big.NewInt(int64(value * math.Pow10(6)))
	result = bv.Mul(bv, big.NewInt(int64(multiplier)))
	return result

}
