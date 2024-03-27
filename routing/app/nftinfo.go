package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"meta-mall/config"
	"meta-mall/database"
	"meta-mall/model"
	"meta-mall/pkg"
	"meta-mall/routing/types"
	"time"
)

func GetNftList(c *fiber.Ctx) error {
	//userId := c.Locals("user_id")
	nft := model.NftInfo{}
	nl, err := nft.GetAllNftInfoByFlag(1, database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询账户失败"))
	}
	data := types.NftListResp{
		List: make([]types.NftDetail, 0),
	}
	for _, nft := range nl {
		in := types.NftDetail{
			Id:              nft.ID,
			Name:            nft.Name,
			NftNumber:       nft.NftNumber,
			TokenId:         nft.TokenId,
			ContractAddress: nft.ContractAddress,
			OwnerAddress:    nft.OwnerAddress,
			Price:           nft.Price,
			TokenName:       nft.ChainName,
			Power:           nft.Power,
			TypeNum:         nft.TypeNum,
			ImgUrl:          nft.ImgUrl,
			Flag:            nft.Flag,
		}
		data.List = append(data.List, in)

	}
	return c.JSON(pkg.SuccessResponse(data))
}
func PurchaseCheck(c *fiber.Ctx) error {
	fmt.Println("/PurchaseCheck api...")
	reqParams := types.PurchaseCheackReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	nft := model.NftInfo{}
	nft.ID = reqParams.NftId
	err = nft.GetById(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get nft error", ""))
	}
	if nft.Flag == "3" {
		return c.JSON(pkg.SuccessResponse(-1))
	} else if nft.Flag == "1" {
		err = database.DB.Transaction(func(tx *gorm.DB) error {
			nft.Flag = "3"
			err := nft.UpdateNftInfo(tx)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "update nft status error", ""))
		}

	} else {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "nft status error", ""))
	}
	return c.JSON(pkg.SuccessResponse(1))
}
func CancelCheck(c *fiber.Ctx) error {
	fmt.Println("/CancelCheck api...")
	reqParams := types.CancelCheackReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	nft := model.NftInfo{}
	nft.ID = reqParams.NftId
	err = nft.GetById(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get nft error", ""))
	}
	if nft.Flag == "1" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "nft has not been locked", ""))
	} else if nft.Flag == "3" {
		err = database.DB.Transaction(func(tx *gorm.DB) error {
			nft.Flag = "1"
			err := nft.UpdateNftInfo(tx)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "update nft status error", ""))
		}

	} else {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "nft status error", ""))
	}
	return c.JSON(pkg.SuccessResponse(""))
}
func PurchaseNft(c *fiber.Ctx) error {
	fmt.Println("/PurchaseNft api...")
	reqParams := types.PurchaseNftReq{}
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	user := model.User{}
	user.ID = userId
	err := user.GetById(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	err = c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	nft := model.NftInfo{}
	nft.ID = reqParams.NftId
	err = nft.GetById(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get nft error", ""))
	}
	if nft.Flag != "3" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "nft is not saleable now error", ""))
	}
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		order := model.Order{
			NftId:        nft.ID,
			Price:        nft.Price,
			BuyerId:      userId,
			BuyerAddress: user.WalletAddress,
			Flag:         "1",
			Buyer:        user,
		}
		err = order.InsertNewOrder(tx)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "insert order error", ""))
		}
		nft.Flag = "2"
		err = nft.UpdateNftInfo(tx)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "update nft error", ""))
		}
		tt := time.Now()
		contract := model.Contract{
			NFTName:            nft.Name,
			NftId:              nft.ID,
			Hash:               reqParams.Hash,
			AccumulatedBenefit: 0,
			TokenName:          reqParams.TokenName,
			Power:              nft.Power,
			BenefitLimit:       nft.Power * 300,
			StartTime:          &tt,
			ExpireTime:         nil,
			Flag:               "1",
			OwnerId:            userId,
			Owner:              user,
		}
		err = contract.InsertNewContract(tx)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "insert contract error", ""))
		}
		return nil
	})
	if err != nil {
		return err
	}
	data := types.PurchaseNftResp{
		Price: nft.Price,
	}
	return c.JSON(pkg.SuccessResponse(data))
}
func GetMyNftList(c *fiber.Ctx) error {
	fmt.Println("/GetMyNftList api...")
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	user := model.User{}
	user.ID = userId
	err := user.GetById(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}

	order := model.Order{}
	order.BuyerId = userId
	ol, err := order.SelectMyOrder(database.DB)
	if err != nil {
		return err
	}
	data := types.GetMyNftListResp{
		List: make([]types.NftOrderDetail, 0),
	}
	nft := model.NftInfo{}
	for _, or := range ol {
		if or.Flag != "0" {
			nft.ID = or.NftId
			err := nft.GetById(database.DB)
			if err != nil {
				return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get nft error", ""))
			}
			in := types.NftOrderDetail{
				Id:              nft.ID,
				Name:            nft.Name,
				NftNumber:       nft.NftNumber,
				TokenId:         nft.TokenId,
				ContractAddress: nft.ContractAddress,
				OwnerAddress:    nft.OwnerAddress,
				Price:           nft.Price,
				TokenName:       nft.ChainName,
				Power:           nft.Power,
				TypeNum:         nft.TypeNum,
				ImgUrl:          nft.ImgUrl,
				OrderId:         or.ID,
				OrderFlag:       or.Flag,
			}
			data.List = append(data.List, in)
		}
	}
	return c.JSON(pkg.SuccessResponse(data))
}
