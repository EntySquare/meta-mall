package tron

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/mr-tron/base58"
	"github.com/savsgio/gotils/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/sha3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/big"
	"strconv"
)

// SendUsdt 发usdt 交易
//
//	@Description: 发usdt 交易
//	@param toAddress 收钱地址
//	@param usdt 发送的usdt数量 1usdt = "1"
//	@return err

// SendUsdt 发usdt 交易
//
//	@Description:
//	@param fromPrivateKey 发起者钱包私钥
//	@param toAddress 收钱地址
//	@param usdt 发送的usdt数量 1usdt = "1000000"
//	@return from 交易详细 from
//	@return to 交易详细 to
//	@return valueStr 交易详细 value 交易金额 1usdt = "1000000" 链的单位
//	@return txId 交易id
//	@return err
func SendUsdt(fromPrivateKey, toAddress, usdt string) (from string, to string, valueStr string, txId string, err error) {
	//mnemonicStr := "mimic link field aisle nut tail endorse witness business garlic mean carry churn yard narrow owner oyster fix dash canyon position nurse pond police"
	//
	//seed := bip39.NewSeed(mnemonicStr, "")
	//
	//// 使用种子生成 BIP32 主密钥
	//masterKey, err := bip32.NewMasterKey(seed)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 生成 BIP32 子密钥
	//childKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild)
	//if err != nil {
	//	panic(err)
	//}
	//
	privateKeyBytes, err := hex.DecodeString(fromPrivateKey)
	if err != nil {
		return "", "", "", "", err
	}

	// 将子密钥转换为 ECDSA 私钥
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	// 获取 ECDSA 公钥
	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	// 使用公钥生成波场地址
	tronAddress := publicKeyToTronAddress(publicKey)

	//privateKeyBytes := crypto.FromECDSA(privateKey)
	//privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)
	//fmt.Println("Private Key:", privateKeyHex)
	fmt.Println(tronAddress)

	// 创建客户端连接到波场节点
	grpcClient := client.NewGrpcClient("grpc.trongrid.io:50051")
	defer grpcClient.Stop()
	if grpcClient == nil {
		log.Fatal("Failed to create grpcClient:", err)
	}

	defer grpcClient.Stop()
	err = grpcClient.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("grpcClient.Start(grpc.WithTransportCredentials(insecure.NewCredentials())) err:", err)
	}
	// 获取当前账户余额
	account, err := grpcClient.GetAccount(tronAddress)
	if err != nil {
		log.Fatal("Failed to get account:", err)
	}
	fmt.Printf("Account balance: %d\n", account.Balance)

	// USDT 的合约地址
	usdtAddress := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"

	// 获取 USDT 代币余额信息
	usdtBalance, err := grpcClient.TRC20ContractBalance(tronAddress, usdtAddress)
	if err != nil {
		log.Fatal("Failed to create USDT token:", err)
	}
	fmt.Printf("USDT balance: %s\n", usdtBalance.String())

	// 构建 USDT 转账交易
	//value, ok := decimals.FromString(usdt)
	//if !ok {
	//	log.Fatal("Failed to parse amount")
	//}
	//tokenDecimals, err := grpcClient.TRC20GetDecimals(usdtAddress)
	//if err != nil {
	//	log.Fatal("Failed to get USDT decimals:", err)
	//}
	//amount, _ := decimals.ApplyDecimals(value, tokenDecimals.Int64())
	amount, _ := new(big.Int).SetString(usdt, 10)
	txe, err := grpcClient.TRC20Send(tronAddress, toAddress, usdtAddress, amount, 100000000) //feeLimit 100TRX
	if err != nil {
		log.Fatal("Failed to send USDT :", err)
	}

	// 签名交易
	signature, err := crypto.Sign(txe.Txid, privateKey)
	if err != nil {
		log.Fatal("Failed to sign transaction:", err)
	}
	txe.Transaction.Signature = append(txe.Transaction.Signature, signature)

	// 广播交易
	result, err := grpcClient.Broadcast(txe.Transaction)
	if err != nil {
		log.Fatal("Failed to broadcast transaction:", err)
	}
	fmt.Printf("Transaction %x result: %v\n", txe.Txid, result)
	txeStr := hex.EncodeToString(txe.Txid)
	return tronAddress, toAddress, usdt, txeStr, nil
}

// 地址生成，注意校验和的计算
func publicKeyToTronAddress(publicKey *ecdsa.PublicKey) string {
	// 获取公钥的字节表示形式
	pubBytes := crypto.FromECDSAPub(publicKey)

	// 使用 Keccak-256（SHA-3）哈希算法计算公钥哈希
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubBytes[1:])
	address := hash.Sum(nil)[12:]

	// 将波场地址字节前添加 0x41（十进制 65）
	tronAddressBytes := append([]byte{0x41}, address...)

	// 计算波场地址的校验和
	hash256 := sha256.New()
	hash256.Write(tronAddressBytes)
	firstHash := hash256.Sum(nil)
	hash256.Reset()
	hash256.Write(firstHash)
	secondHash := hash256.Sum(nil)
	checksum := secondHash[:4]

	// 将波场地址字节与校验和合并，生成原始地址
	rawAddress := append(tronAddressBytes, checksum...)

	// 使用 Base58 编码生成波场地址字符串表示形式
	tronAddress := base58.Encode(rawAddress)
	return tronAddress
}

// CreateTromAddress 创建钱包
func CreateTromAddress() (privateKeyStr string, tronAddressStr string, err error) {
	// 生成 256 位熵
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		panic(err)
	}

	// 使用熵生成助记词
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}
	fmt.Println("Mnemonic:", mnemonic)

	// 使用助记词和空密码生成种子
	seed := bip39.NewSeed(mnemonic, "")

	// 使用种子生成 BIP32 主密钥
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		panic(err)
	}

	// 生成 BIP32 子密钥
	childKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild)
	if err != nil {
		panic(err)
	}

	// 将子密钥转换为 ECDSA 私钥
	privateKey, err := crypto.ToECDSA(childKey.Key)
	if err != nil {
		panic(err)
	}

	// 获取 ECDSA 公钥
	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	// 使用公钥生成波场地址
	tronAddress := publicKeyToTronAddress(publicKey)

	// 输出私钥和波场地址
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)
	fmt.Println("Private Key:", privateKeyHex)
	fmt.Println("TRON Address:", tronAddress)
	return privateKeyHex, tronAddress, nil
}

// SendTxr
//
//	@Description: 发送TXR交易
//	@param fromPrivateKeyStr 发送者钱包私钥
//	@param toAddressStr 接收者钱包地址
//	@param trxSun  发送的trx数量 1 TRX = 1000000 sun 转1个TRX
//	@return from 发送者钱包地址
//	@return to 接收者钱包地址
//	@return valueStr 发送的trx数量
//	@return txId 交易ID
//	@return err 错误信息
func SendTxr(fromPrivateKeyStr, toAddressStr string, trxSun int64) (from string, to string, valueStr string, txId string, err error) {
	//// 从助记词恢复私钥
	//mnemonicStr := "mimic link field aisle nut tail endorse witness business garlic mean carry churn yard narrow owner oyster fix dash canyon position nurse pond police"
	//
	//seed := bip39.NewSeed(mnemonicStr, "")
	//
	//// 使用种子生成 BIP32 主密钥
	//masterKey, err := bip32.NewMasterKey(seed)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 生成 BIP32 子密钥
	//childKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 将子密钥转换为 ECDSA 私钥
	//privateKey, err := crypto.ToECDSA(childKey.Key)
	//if err != nil {
	//	panic(err)
	//}
	privateKeyBytes, err := hex.DecodeString(fromPrivateKeyStr)
	if err != nil {
		return "", "", "", "", err
	}

	// 将子密钥转换为 ECDSA 私钥
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	// 获取 ECDSA 公钥
	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	// 使用公钥生成波场地址
	tronAddress := publicKeyToTronAddress(publicKey)

	//privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)
	fmt.Println("Private Key:", privateKeyHex)
	fmt.Println(tronAddress)

	// 创建客户端连接到波场节点
	grpcClient := client.NewGrpcClient("grpc.trongrid.io:50051")
	if grpcClient == nil {
		log.Fatal("Failed to create grpcClient:", err)
	}

	defer grpcClient.Stop()
	err = grpcClient.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	// 获取当前账户余额
	account, err := grpcClient.GetAccount(tronAddress)
	if err != nil {
		log.Fatal("Failed to get account:", err)
	}
	fmt.Printf("Account balance: %d\n", account.Balance)

	// 构建交易 - 1 TRX = 1000000 sun 转1个TRX
	//toAddress, err := common.DecodeCheck("TQ7Zft6PKXJQTH3K6oFAQ5D1m5eGyDiJz8")
	//if err != nil {
	//	log.Fatal("Failed to decode destination address:", err, toAddress)
	//}
	//amount := int64(1000000) // 设置转账金额（单位：sun）
	//tx, err := grpcClient.Transfer(tronAddress, "TQ7Zft6PKXJQTH3K6oFAQ5D1m5eGyDiJz8", account.Balance)
	tx, err := grpcClient.Transfer(tronAddress, toAddressStr, trxSun)
	if err != nil {
		log.Fatal("Failed to create transaction:", err)
	}

	// 签名交易
	signature, err := crypto.Sign(tx.Txid, privateKey)
	if err != nil {
		log.Fatal("Failed to sign transaction:", err)
	}
	tx.Transaction.Signature = append(tx.Transaction.Signature, signature)

	// 广播交易
	result, err := grpcClient.Broadcast(tx.Transaction)
	if err != nil {
		log.Fatal("Failed to broadcast transaction:", err)
	}
	fmt.Printf("Transaction %x result: %v\n", tx.Txid, result)
	txeStr := hex.EncodeToString(tx.Txid)
	return tronAddress, toAddressStr, strconv.FormatInt(trxSun, 10), txeStr, nil
}

func SendTxrTest(fromPrivateKeyStr, toAddressStr string, trxSun int64) (from string, to string, valueStr string, txId string, err error) {
	return uuid.V4(), toAddressStr, strconv.FormatInt(trxSun, 10), uuid.V4(), nil
}
func SendUsdtTest(fromPrivateKey, toAddress, usdt string) (from string, to string, valueStr string, txId string, err error) {
	return uuid.V4(), toAddress, usdt, uuid.V4(), nil
}

// TODO 我调用你 实现需要扫的交易hash
func PushTxId(txId string) error {
	return nil
}

// TODO 需要扫的钱包地址
//func PushAddress(cAddress []string) error {
//	_, err := scan.RegisterAddress(cAddress)
//	if err != nil {
//		panic(err)
//	}
//	return nil
//}
