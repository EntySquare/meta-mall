package app

import (
	"github.com/gofiber/fiber/v2"
	"meta-mall/config"
	"meta-mall/pkg"
	"path/filepath"
)

// UploadIMG 上传视频文件
func UploadIMG(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "from filer err", ""))
	}
	open, err := file.Open()
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Open filer err", ""))
	}
	// 上传文件
	url, err := pkg.SetFileOss(open, filepath.Ext(file.Filename))
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "setPngOss err", ""))
	}

	return c.JSON(pkg.SuccessResponse(url))
}
