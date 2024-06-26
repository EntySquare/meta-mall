package routing

import (
	"github.com/gofiber/fiber/v2"
	"meta-mall/routing/app"
	intcpt "meta-mall/routing/intercept"
)

func Setup(f *fiber.App) {
	appApi := f.Group("/app")
	AppSetUp(appApi)

}

func AppSetUp(appApi fiber.Router) {
	appApi.Post("/login", app.LoginAndRegister)                      //login //login
	appApi.Post("/myInvestment", intcpt.AuthApp(), app.MyInvestment) //login
	appApi.Post("/getNftList", app.GetNftList)
	appApi.Post("/purchaseNft", intcpt.AuthApp(), app.PurchaseNft)
	appApi.Post("/purchaseCheck", intcpt.AuthApp(), app.PurchaseCheck)
	appApi.Post("/cancelCheck", intcpt.AuthApp(), app.CancelCheck)
	appApi.Post("/myNftList", intcpt.AuthApp(), app.GetMyNftList)
	appApi.Post("/getContractList", intcpt.AuthApp(), app.GetMyContractList)
	appApi.Post("/getMiningIncome", intcpt.AuthApp(), app.GetMiningIncome)
	appApi.Post("/getMyPromotionList", intcpt.AuthApp(), app.MyInvestment)
	appApi.Post("/getMyPromotion", intcpt.AuthApp(), app.MyPromotion)
	appApi.Post("/getAvailableBenefit", intcpt.AuthApp(), app.GetAvailableBenefit)
	appApi.Post("/applyForMember", intcpt.AuthApp(), app.ApplyForMember)

	appApi.Post("/manager/insertNft", intcpt.AuthManagerApp(), app.InsertNft)
	appApi.Post("/manager/tokenIdFrom", intcpt.AuthManagerApp(), app.TokenIdFrom)
	appApi.Post("/manager/login", app.LoginManager)
	appApi.Post("/manager/getNftlist", intcpt.AuthManagerApp(), app.GetManagerNftList)
	appApi.Post("/manager/offNft", intcpt.AuthManagerApp(), app.OffNft)
	appApi.Post("/manager/uploadIMG", intcpt.AuthManagerApp(), app.UploadIMG) //上传
	appApi.Post("/manager/setIncome", intcpt.AuthManagerApp(), app.SetIncome)
	appApi.Post("/manager/getIncome", intcpt.AuthManagerApp(), app.GetIncome)
	appApi.Post("/manager/getMemberList", intcpt.AuthManagerApp(), app.GetMemberList)
	appApi.Post("/manager/createNewMember", intcpt.AuthManagerApp(), app.CreateNewMember)
	appApi.Post("/manager/offMember", intcpt.AuthManagerApp(), app.ChangeMemberStatus)

	//appApi.Post("/deposit", intcpt.AuthApp(), app.Deposit)
	//appApi.Post("/withdraw", intcpt.AuthApp(), app.Withdraw)
	//appApi.Post("/checkHash", intcpt.AuthApp(), app.CheckHashApi)
	//appApi.Post("/cancelContract", intcpt.AuthApp(), app.CancelContract)
	//appApi.Post("/getAllNftInfo", intcpt.AuthApp(), app.GetAllNftInfo)
	//appApi.Post("/updateNftInfo", intcpt.AuthManagerApp(), app.UpdateNftInfo)
	//appApi.Post("/pledgeNft", intcpt.AuthApp(), app.pledgeNft)

}
