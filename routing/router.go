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
	appApi.Post("/login", app.LoginAndRegister)                                //login
	appApi.Post("/myCovenantFlow", intcpt.AuthApp(), app.MyCovenantFlow)       //login
	appApi.Post("/myNgt", intcpt.AuthApp(), app.MyNgt)                         //login
	appApi.Post("/myInvestment", intcpt.AuthApp(), app.MyInvestment)           //login
	appApi.Post("/getInviteeInfo", intcpt.AuthApp(), app.GetInviteeInfo)       //login
	appApi.Post("/getCovenantDetail", intcpt.AuthApp(), app.GetCovenantDetail) //login
	appApi.Post("/getNftList", app.GetNftList)
	//appApi.Post("/pledgeNft", intcpt.AuthApp(), app.PledgeNft)
	//appApi.Post("/deposit", intcpt.AuthApp(), app.Deposit)
	//appApi.Post("/withdraw", intcpt.AuthApp(), app.Withdraw)
	//appApi.Post("/checkHash", intcpt.AuthApp(), app.CheckHashApi)
	//appApi.Post("/cancelCovenant", intcpt.AuthApp(), app.CancelCovenant)
	//appApi.Post("/getAllNftInfo", intcpt.AuthApp(), app.GetAllNftInfo)
	//appApi.Post("/updateNftInfo", intcpt.AuthManagerApp(), app.UpdateNftInfo)
	//appApi.Post("/pledgeNft", intcpt.AuthApp(), app.pledgeNft)

}
