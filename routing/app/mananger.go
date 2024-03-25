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
	password := reqParams.Password
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
			err := nft.InsertNewNftInfo(database.DB)
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
	nl, err := nft.GetAllNftInfoByFlag(1, database.DB)
	nl2, err := nft.GetAllNftInfoByFlag(2, database.DB)
	nl3, err := nft.GetAllNftInfoByFlag(0, database.DB)
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
	for _, nft := range nl2 {
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
	for _, nft := range nl3 {
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
func TokenIdFrom(c *fiber.Ctx) error {
	data := types.TokenIdFromResp{}
	nftInfo := model.NftInfo{}
	id, err := nftInfo.GetMaxTokenId(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get token id error", ""))
	}
	data.TokenId = id
	return c.JSON(pkg.SuccessResponse(data))
}
