package intcpt

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"meta-mall/config"
	"meta-mall/database"
	"meta-mall/model"
	"meta-mall/pkg"
)

// AuthApp Protected protect routes
func AuthApp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//var token = c.Get("token")
		//fmt.Println(token)
		//c.Locals("guo", name)
		//c.Locals("guo2", "23123123123123123")
		var (
			userId int64
			token  = c.Get(config.LOCAL_TOKEN)
			db     = database.DB
			err    error
		)

		//开关 测试打开token
		//if true {
		//	c.Locals(config.LOCAL_USERID_UINT, 1)
		//	c.Locals(config.LOCAL_USERID_INT64, 1)
		//	_ = c.Next()
		//	return nil
		//}

		// 打印请求地址
		log.Info("Request URL: ", c.Path())
		log.Info("Request JSON: ", string(c.Body()))
		if token == "" || len(token) < 10 {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is null", ""))
		}

		userId, tokenData, err := model.UserSelectIdByToken(db, token)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		if !pkg.CheckSpecialCharacters(&token) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		//检查token 有效时间
		if !pkg.CheckTokenValidityTime(&tokenData) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is exceed", ""))
		}

		//刷新token有效时间
		if err = model.UserRefreshToken(db, userId, tokenData); err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "db UserRefreshAppToken err", ""))
		}
		//if true {
		//	err := c.JSON("token???")
		//	if err != nil {
		//		return err
		//	}
		//	return nil
		//}
		//
		c.Locals(config.LOCAL_USERID_UINT, uint(userId))
		c.Locals(config.LOCAL_USERID_INT64, userId)
		_ = c.Next()
		//c.JSON("231231231")
		return nil
	}
}

// AuthApp Protected protect routes
func AuthManagerApp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//var token = c.Get("token")
		//fmt.Println(token)
		//c.Locals("guo", name)
		//c.Locals("guo2", "23123123123123123")
		var (
			username string
			token    = c.Get(config.LOCAL_TOKEN)
			db       = database.DB
			err      error
		)

		//开关 测试打开token
		//if true {
		//	c.Locals(config.LOCAL_USERID_UINT, 1)
		//	c.Locals(config.LOCAL_USERID_INT64, 1)
		//	_ = c.Next()
		//	return nil
		//}

		// 打印请求地址
		log.Info("Request URL: ", c.Path())
		log.Info("Request JSON: ", string(c.Body()))
		if token == "" || len(token) < 10 {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is null", ""))
		}

		username, tokenData, err := model.ManagerSelectIdByToken(db, token)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		if !pkg.CheckSpecialCharacters(&token) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		//检查token 有效时间
		if !pkg.CheckTokenValidityTime(&tokenData) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is exceed", ""))
		}

		//刷新token有效时间
		if err = model.ManagerRefreshToken(db, username, tokenData); err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "db UserRefreshAppToken err", ""))
		}
		//if true {
		//	err := c.JSON("token???")
		//	if err != nil {
		//		return err
		//	}
		//	return nil
		//}
		//
		c.Locals(config.LOCAL_MANAGERNAME_STRING, username)
		_ = c.Next()
		//c.JSON("231231231")
		return nil
	}
}
