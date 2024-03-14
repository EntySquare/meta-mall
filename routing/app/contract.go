package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"meta-mall/config"
	"meta-mall/database"
	"meta-mall/model"
	"meta-mall/pkg"
	"meta-mall/routing/types"
)

func GetMyContractList(c *fiber.Ctx) error {
	fmt.Println("/GetMyContractList api...")
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	user := model.User{}
	user.ID = userId
	err := user.GetById(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	contract := model.Contract{}
	contract.OwnerId = userId
	cl, err := contract.SelectMyContract(database.DB)
	if err != nil {
		return err
	}

	data := types.GetMyContractListResp{
		List: make([]types.ContractDetail, 0),
	}
	for _, contract := range cl {
		in := types.ContractDetail{
			Id:                 contract.ID,
			AccumulatedBenefit: contract.AccumulatedBenefit,
			Power:              contract.Power,
			TokenName:          contract.TokenName,
			Flag:               contract.Flag,
		}
		if contract.Flag == "2" {
			in.StartTime = contract.StartTime.Unix()
		} else if contract.Flag == "1" {
			in.StartTime = 0
		}
		data.List = append(data.List, in)
	}

	return c.JSON(pkg.SuccessResponse(data))
}
func GetMiningIncome(c *fiber.Ctx) error {
	fmt.Println("/GetMiningIncome api...")
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	user := model.User{}
	user.ID = userId
	err := user.GetById(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	myFlows := model.ContractFlow{}
	myFlows.UserId = userId
	myFlows.TokenName = "usdt"
	myFlows.Flag = "1"
	usdtmab, err := myFlows.GetUserReleaseBenefitByTokenName(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user available benefit  error", ""))
	}
	myFlows.TokenName = "unc"
	uncmab, err := myFlows.GetUserReleaseBenefitByTokenName(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user available benefit  error", ""))
	}
	contract := model.Contract{}
	contract.TokenName = "usdt"
	contract.Flag = "2"
	usdtab, err := contract.GetAllBenefitLimitByTokenName(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get  accumulated benefit  error", ""))
	}
	contract.OwnerId = userId
	usdtmb, err := contract.GetUserAccumulatedBenefitByTokenName(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get  user accumulated benefit  error", ""))
	}
	contract.TokenName = "unc"
	unctab, err := contract.GetAllBenefitLimitByTokenName(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get  accumulated benefit  error", ""))
	}
	uncmb, err := contract.GetUserAccumulatedBenefitByTokenName(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get  user accumulated benefit  error", ""))
	}
	ap, err := contract.GetAllPowers(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get   accumulated power  error", ""))
	}
	data := types.GetMiningIncomeResultResp{
		AllPowers:                 ap,
		AllAccumulatedUSDTBenefit: usdtab,
		MyAccumulatedUSDTBenefit:  usdtmb,
		MyAvailableUSDTBenefit:    usdtmab,
		AllAccumulatedUNCBenefit:  unctab,
		MyAccumulatedUNCBenefit:   uncmb,
		MyAvailableUNCBenefit:     uncmab,
	}
	return c.JSON(pkg.SuccessResponse(data))
}
