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
			StartTime:          contract.StartTime.Unix(),
			Flag:               contract.Flag,
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
		return err
	}
	myFlows.TokenName = "unc"
	uncmab, err := myFlows.GetUserReleaseBenefitByTokenName(database.DB)
	if err != nil {
		return err
	}
	contract := model.Contract{}
	contract.TokenName = "usdt"
	contract.Flag = "2"
	usdtab, err := contract.GetAllBenefitLimitByTokenName(database.DB)
	if err != nil {
		return err
	}

	contract.OwnerId = userId
	usdtmb, err := contract.GetUserAccumulatedBenefitByTokenName(database.DB)
	if err != nil {
		return err
	}
	contract.TokenName = "unc"
	unctab, err := contract.GetAllBenefitLimitByTokenName(database.DB)
	if err != nil {
		return err
	}
	uncmb, err := contract.GetUserAccumulatedBenefitByTokenName(database.DB)
	if err != nil {
		return err
	}
	data := types.GetMiningIncomeResultResp{
		AllAccumulatedUSDTBenefit: usdtab,
		MyAccumulatedUSDTBenefit:  usdtmb,
		MyAvailableUSDTBenefit:    usdtmab,
		AllAccumulatedUNCBenefit:  unctab,
		MyAccumulatedUNCBenefit:   uncmb,
		MyAvailableUNCBenefit:     uncmab,
	}
	return c.JSON(pkg.SuccessResponse(data))
}
func GetAvailableBenefit(c *fiber.Ctx) error {

	return nil
}
