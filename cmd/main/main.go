package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	contract "meta-mall/contracts"
	"meta-mall/database"
	"meta-mall/model"
	"meta-mall/model/api"
	"meta-mall/routing"
	"time"
)

func main() {
	database.ConnectDB()
	fiberApp := fiber.New()
	// 创建一个速率限制器，每秒最多只允许10个请求
	//fiberApp.Use(limiter.New(limiter.Config{
	//	Max:        10,
	//	Expiration: 2 * time.Second,
	//	KeyGenerator: func(c *fiber.Ctx) string {
	//		return c.IP() // 使用客户端IP作为限流key
	//	},
	//}))
	// 添加 CORS 中间件
	fiberApp.Use(cors.New())
	// 将速率限制器添加到路由中间件中
	err := api.InitUserTree(database.DB)
	if err != nil {
		fmt.Println(err.Error())
	}

	InitTask()
	go scanContract(database.DB)
	//  go eth.ScanEth(database.DB)
	routing.Setup(fiberApp)

	err = fiberApp.Listen(":4001")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func InitTask() {
	var (
		c = cron.New(cron.WithSeconds())
		//db  = database.DB
		err error
	)
	_, err = c.AddFunc("0 0 1 * * ?", func() {
		api.IncomeRunP(database.DB)
		//api.ContractCycle(database.DB)
	})

	if err != nil {
		panic(err)
		return
	}
	c.Start()
}
func scanContract(db *gorm.DB) {
	for {
		err := database.DB.Transaction(func(tx *gorm.DB) error {
			tt := time.Now()
			list, err := model.SelectContractByFlag(tx, "1")
			if err != nil {
				panic(err)
			}
			for _, v := range list {
				hash := v.Hash
				checkFlag := false
				txs, err := contract.GetTransactionByHash(hash)
				if err != nil {
					println(err.Error())
				} else {
					if txs["blockNumber"] != "" {
						fmt.Println(txs["blockNumber"])
						checkFlag = true
					}

				}
				if checkFlag == true {
					v.Flag = "2"
					v.StartTime = &tt
					nft := model.NftInfo{}
					nft.ID = v.NftId
					err = nft.GetById(db)
					if err != nil {
						return err
					}
					nft.Flag = "0"
					err = nft.UpdateNftInfo(tx)
					if err != nil {
						return err
					}
					order := model.Order{}
					order.NftId = v.NftId
					err = order.GetByNftId(db)
					if err != nil {
						return err
					}
					order.Flag = "2"
					err := order.UpdateOrder(tx)
					if err != nil {
						return err
					}
					err = v.UpdateContract(tx)
					if err != nil {
						return err
					}
					api.UserPurchase(v.OwnerId, v.Power)
					user := model.User{}
					user.ID = v.OwnerId
					err = user.GetById(db)
					if err != nil {
						return err
					}
					if user.RecommendId != 0 {
						err = updateTreeUserLevel(tx, user)
						if err != nil {
							return err
						}
					}

				}
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		time.Sleep(time.Second * 60)
	}
}
func updateTreeUserLevel(tx *gorm.DB, user model.User) error {
	newLevel := api.UserLevelCheck(user.RecommendId)
	user.Level = newLevel
	err := user.UpdateUserLevel(tx)
	if err != nil {
		return err
	}
	if user.RecommendId != 0 {
		recommender := model.User{}
		recommender.ID = user.RecommendId
		err := updateTreeUserLevel(tx, recommender)
		if err != nil {
			return err
		}
	}
	return nil
}
