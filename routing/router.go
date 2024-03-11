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
	appApi.Post("/myNftList", intcpt.AuthApp(), app.GetMyNftList)
	appApi.Post("/getContractList", intcpt.AuthApp(), app.GetMyContractList) //login
	//appApi.Post("/deposit", intcpt.AuthApp(), app.Deposit)
	//appApi.Post("/withdraw", intcpt.AuthApp(), app.Withdraw)
	//appApi.Post("/checkHash", intcpt.AuthApp(), app.CheckHashApi)
	//appApi.Post("/cancelContract", intcpt.AuthApp(), app.CancelContract)
	//appApi.Post("/getAllNftInfo", intcpt.AuthApp(), app.GetAllNftInfo)
	//appApi.Post("/updateNftInfo", intcpt.AuthManagerApp(), app.UpdateNftInfo)
	//appApi.Post("/pledgeNft", intcpt.AuthApp(), app.pledgeNft)

}
