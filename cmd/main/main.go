package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/robfig/cron/v3"
	"meta-mall/database"
	"meta-mall/model/api"
	"meta-mall/routing"
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
	//	go poly.ScanPoly(database.DB)
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
		//api.IncomeRunP(database.DB)
		api.CovenantCycle(database.DB)
	})

	if err != nil {
		panic(err)
		return
	}
	c.Start()
}
