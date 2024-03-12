package app

import (
	"errors"
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
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "", ""))
		}
	}

	recommendUser := model.User{}
	fmt.Println("推荐码：", reqParams.Code)
	database.DB.Model(&model.User{}).Where("uid = ?", reqParams.Code).Find(&recommendUser)

	err = database.DB.Transaction(func(tx *gorm.DB) error {
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
	data.Address = user.WalletAddress
	data.InvestmentAddress = config.Config("WEB_URL") + "/&code=" + user.UID
	data.Level = user.Level
	data.Powers = api.UserTree[userId].Power
	//data.AccumulatedPledgeCount = api.GetBranchAccumulatedPower(userId)
	for _, branch := range api.UserTree[userId].Branch {
		in := types.InvestmentUserInfo{}
		in.Address = api.UserTree[branch].Address
		in.Level = api.UserTree[branch].Level
		in.Powers = api.UserTree[branch].Power
		user.WalletAddress = api.UserTree[branch].Address
		err = user.GetByWalletAddress(database.DB)
		if err != nil {
			return err
		}
		in.Time = user.Model.CreatedAt.Unix()
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
	flow.Flag = "2"
	flow.TokenName = "unc"
	mrb, err := flow.GetUserReleaseBenefitByTokenName(database.DB)
	if err != nil {
		return err
	}
	mb := 2100.0
	data := types.MyPromotionResp{
		AllPromotionPower:           api.GetBranchAccumulatedPower(0),
		MyPromotionPower:            api.UserTree[userId].Power,
		MyPromotionBenefit:          mb,
		MyAvailablePromotionBenefit: mb - mrb,
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
