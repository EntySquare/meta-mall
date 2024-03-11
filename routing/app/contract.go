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
	fmt.Println("/GetMyNftList api...")
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
			StartTime:          contract.StartTime.Unix(),
			Flag:               contract.Flag,
		}
		data.List = append(data.List, in)
	}

	return c.JSON(pkg.SuccessResponse(data))
}
