package test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"meta-mall/pkg"
	"testing"
)

func TestRsa1(t *testing.T) {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	// 明文消息
	message := "Hello, world!"

	// 加密明文消息
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		[]byte(message),
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后的消息：%x\n", ciphertext)

	// 解密密文消息
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		ciphertext,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后的消息：%s\n", plaintext)
}

func TestRsa2(t *testing.T) {
	hash := pkg.RSAEncrypt("dadssda", "")
	fmt.Println(hash)

	fmt.Println(pkg.RSADecrypt(hash, ""))

}
func TestRsa33(t *testing.T) {
	//pkg.RSADecrypt("Hello, world!")
}
