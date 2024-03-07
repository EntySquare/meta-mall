package app

import (
	"github.com/gofiber/fiber/v2"
	"meta-mall/config"
	"meta-mall/database"
	"meta-mall/model"
	"meta-mall/pkg"
	"meta-mall/routing/types"
)

func GetNftList(c *fiber.Ctx) error {
	//userId := c.Locals("user_id")

	nft := model.NftInfo{}
	nl, err := nft.GetAllNftInfoByFlag(database.DB)
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
			TypeNum:         nft.TypeNum,
			ImgUrl:          nft.ImgUrl,
		}
		data.List = append(data.List, in)

	}
	return c.JSON(pkg.SuccessResponse(data))
}
