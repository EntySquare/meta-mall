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

func LoginManager(c *fiber.Ctx) error {
	fmt.Println("/LoginManager api...")
	reqParams := types.LoginMangerReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	username := reqParams.UserName
	inputPassword := reqParams.Password
	password := api.Get256Pw(inputPassword)
	manager := model.Manager{}
	manager.UserName = username
	err = manager.GetByUsername(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	data := types.LoginManagerResp{}
	if password == manager.Password {
		returnT := strings.Split(manager.Token, ":")[0]
		c.Locals(config.LOCAL_TOKEN, returnT)
		data.Token = returnT
	} else {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "wrong password", ""))
	}
	returnT := ""
	if !pkg.CheckTokenValidityTime(&manager.Token) {
		returnT = pkg.RandomString(64)
		manager.Token = returnT + ":" + strconv.FormatInt(time.Now().Unix(), 10)
		err := manager.UpdateManagerToken(database.DB)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "update token error", ""))
		}
		c.Locals(config.LOCAL_TOKEN, returnT)
		data.Token = returnT
	}
	data.Flag = manager.Flag
	return c.JSON(pkg.SuccessResponse(data))
}
func InsertNft(c *fiber.Ctx) error {
	fmt.Println("/InsertNft api...")
	reqParams := types.InsertNftReq{}
	userName := c.Locals(config.LOCAL_MANAGERNAME_STRING).(string)
	manger := model.Manager{}
	manger.UserName = userName
	err := manger.GetByUsername(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	err = c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		for x, id := range reqParams.TokenId {
			nft := model.NftInfo{

				Name:            reqParams.Name,
				NftNumber:       int64(x),
				DayRate:         "",
				TokenId:         id,
				ContractAddress: config.Config("CONTRACT_ADDRESS"),
				OwnerAddress:    reqParams.OwnerAddress,
				ChainName:       reqParams.TokenName,
				Price:           reqParams.Price,
				Power:           reqParams.Power,
				TypeNum:         reqParams.TypeNum,
				ImgUrl:          reqParams.ImgUrl,
				Flag:            "1",
			}
			err := nft.InsertNewNftInfo(tx)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "insert nft info error", ""))
	}
	return c.JSON(pkg.SuccessResponse(""))
}
func GetManagerNftList(c *fiber.Ctx) error {
	//userId := c.Locals("user_id")
	nft := model.NftInfo{}
	userName := c.Locals(config.LOCAL_MANAGERNAME_STRING).(string)
	manger := model.Manager{}
	manger.UserName = userName
	err := manger.GetByUsername(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	data := types.NftListResp{
		List: make([]types.NftDetail, 0),
	}
	if manger.Flag == "1" {
		nl, err := nft.GetAllNftInfo(database.DB)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询nft失败"))
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
	} else if manger.Flag == "2" {
		nft.OwnerAddress = manger.UserName
		nl, err := nft.GetAllNftInfoByOwner(database.DB)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询nft失败"))
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
	}

	return c.JSON(pkg.SuccessResponse(data))
}
func TokenIdFrom(c *fiber.Ctx) error {
	data := types.TokenIdFromResp{}
	nftInfo := model.NftInfo{}
	userName := c.Locals(config.LOCAL_MANAGERNAME_STRING).(string)
	manger := model.Manager{}
	manger.UserName = userName
	err := manger.GetByUsername(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	if manger.Flag != "1" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "you don't have permission", ""))
	}
	id, err := nftInfo.GetMaxTokenId(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get token id error", ""))
	}
	data.TokenId = id
	return c.JSON(pkg.SuccessResponse(data))
}
func SetIncome(c *fiber.Ctx) error {
	fmt.Println("/SetIncome api...")
	reqParams := types.SetIncomeReq{}
	userName := c.Locals(config.LOCAL_MANAGERNAME_STRING).(string)
	manger := model.Manager{}
	manger.UserName = userName
	err := manger.GetByUsername(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	if manger.Flag != "1" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "you don't have permission", ""))
	}
	err = c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	api.UncAmount = reqParams.UncAmount
	api.MetaAmount = reqParams.MetaAmount

	return c.JSON(pkg.SuccessResponse(""))
}
func GetIncome(c *fiber.Ctx) error {
	fmt.Println("/GetIncome api...")
	manger := model.Manager{}
	userName := c.Locals(config.LOCAL_MANAGERNAME_STRING).(string)
	manger.UserName = userName
	err := manger.GetByUsername(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	if manger.Flag != "1" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "you don't have permission", ""))
	}
	data := types.GetIncomeResp{}
	data.UncAmount = api.UncAmount
	data.MetaAmount = api.MetaAmount
	return c.JSON(pkg.SuccessResponse(data))
}
func OffNft(c *fiber.Ctx) error {
	fmt.Println("/OffNft api...")
	reqParams := types.OffNftReq{}
	userName := c.Locals(config.LOCAL_MANAGERNAME_STRING).(string)
	manger := model.Manager{}
	manger.UserName = userName
	err := manger.GetByUsername(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	err = c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		for _, id := range reqParams.OffIdList {
			nft := model.NftInfo{}
			nft.ID = id
			err := nft.GetById(tx)
			if err != nil {
				return err
			}
			nft.Flag = "4"
			err = nft.UpdateNftInfo(tx)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "update status error", ""))
	}
	return c.JSON(pkg.SuccessResponse(""))
}
func GetMemberList(c *fiber.Ctx) error {
	fmt.Println("/CheckApplyMember api...")
	userName := c.Locals(config.LOCAL_MANAGERNAME_STRING).(string)
	manger := model.Manager{}
	manger.UserName = userName
	err := manger.GetByUsername(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	if manger.Flag != "1" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "you don't have permission", ""))
	}
	apply := model.Manager{}
	list, err := apply.SelectApplyList(database.DB)
	if err != nil {
		return err
	}
	data := types.GetMemberList{}
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		for _, a := range list {
			data.List = append(data.List, types.Member{
				Id:            a.ID,
				WalletAddress: a.UserName,
				Time:          a.CreatedAt.Unix(),
				Flag:          a.Flag,
			})
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get apply error", ""))
	}
	return c.JSON(pkg.SuccessResponse(data))
}
func CreateNewMember(c *fiber.Ctx) error {
	fmt.Println("/CreateNewMember api...")
	userName := c.Locals(config.LOCAL_MANAGERNAME_STRING).(string)
	manger := model.Manager{}
	manger.UserName = userName
	err := manger.GetByUsername(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	if manger.Flag != "1" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "you don't have permission", ""))
	}
	reqParams := types.CreateNewMemberReq{}
	err = c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	username := reqParams.UserName
	inputPassword := reqParams.Password
	password := api.Get256Pw(inputPassword)
	ma := model.Manager{}
	ma.UserName = username
	err = ma.GetByUsername(database.DB)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "", ""))
		} else {
			err = database.DB.Transaction(func(tx *gorm.DB) error {
				token := pkg.RandomString(64) + ":" + strconv.FormatInt(time.Now().Unix(), 10)
				newManager := model.Manager{
					UserName: username,
					Password: password,
					Token:    token,
					Flag:     "2",
				}
				_, err := newManager.InsertNewManager(tx)
				if err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "", ""))
			}
		}
	} else {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "username has been used before", ""))
	}
	return c.JSON(pkg.SuccessResponse(""))
}
func ChangeMemberStatus(c *fiber.Ctx) error {
	fmt.Println("/OffMember api...")
	reqParams := types.OffMemberReq{}
	userName := c.Locals(config.LOCAL_MANAGERNAME_STRING).(string)
	manger := model.Manager{}
	manger.UserName = userName
	err := manger.GetByUsername(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get user error", ""))
	}
	if manger.Flag != "1" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "you don't have permission", ""))
	}
	err = c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		for _, id := range reqParams.OffIdList {
			ma := model.Manager{}
			ma.ID = id
			err := ma.GetById(tx)
			if err != nil {
				return err
			}
			if ma.Flag == "2" {
				ma.Flag = "0"
			} else if ma.Flag == "0" {
				ma.Flag = "2"
			}
			err = ma.UpdateManager(tx)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "update status error", ""))
	}
	return c.JSON(pkg.SuccessResponse(""))
}
