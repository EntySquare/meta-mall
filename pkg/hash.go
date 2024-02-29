package pkg

import (
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
)

func HashStr(sStr string, hStr string) string {
	salt := []byte(sStr) // 自定义盐
	// 对支付密码进行hash并使用自定义盐
	hashedPassword := pbkdf2.Key([]byte(hStr), salt, 4096, 32, sha256.New)
	// 将二进制哈希密码转换为Base64编码，便于存储和传输
	hashedPasswordBase64 := base64.StdEncoding.EncodeToString(hashedPassword)
	return hashedPasswordBase64
}
