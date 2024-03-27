package test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	contract "meta-mall/contracts"
	"testing"
)

func TestTransfer(t *testing.T) {
	//rpcURL := "https://bsc-dataseed.binance.org/"
	rpcURL := "https://public.stackup.sh/api/v1/node/bsc-testnet/"
	// 转账的私钥
	privateKey := "0xdaba574f36fdc73322e3f0eb30687b06b7eaf0f222b69efa0d6f19a659cf3c76"
	// USDT合约地址
	usdtContractAddress := "0x9EDd0f2B153660469E69fd0436cab3CBaA3CAADE" //test
	// 转账金额
	transferAmount := big.NewInt(100000000000000000) // 0.1 USDT
	// 转账目标地址
	toAddress := "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"
	// 初始化客户端连接
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	// 解析私钥
	privateKeyBytes, err := hexutil.Decode(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	key, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	// 创建交易
	auth := bind.NewKeyedTransactor(key)
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // 转账的ETH数量，如果是转账ETH则需要设置
	auth.GasLimit = uint64(1000000) // 转账的Gas上限
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(usdtContractAddress)
	instance, err := contract.NewContracts(address, client)
	if err != nil {
		log.Fatal(err)
	}

	to := common.HexToAddress(toAddress)
	tx, err := instance.Transfer(auth, to, transferAmount)
	if err != nil {
		fmt.Println("Transferring err...")
	}
	fmt.Println("Transferring USDT...")
	// 等待交易被打包
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction receipt: %+v\n", receipt)
}

func TestCal(t *testing.T) {
	var value float64 = 123.456
	var exponent int = 18

	// 将10的指数转换为整数幂
	multiplier := math.Pow10(exponent - 6)
	var result *big.Int
	bv := big.NewInt(int64(value * math.Pow10(6)))
	result = bv.Mul(bv, big.NewInt(int64(multiplier)))
	println(result.String())
}

//func TestTransfer(t *testing.T) {
//	// 你的 BSC 钱包地址和私钥
//	address := "YOUR_ADDRESS"
//	privateKey := "YOUR_PRIVATE_KEY"
//
//	// 创建 BSC 钱包
//	keyManager, err := keys.NewPrivateKeyManager(privateKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 设置 BSC 网络参数
//	client := tx.NewClient("https://bsc-dataseed.binance.org", types.TestNetwork)
//
//	// 获取 USDT 代币合约地址
//	usdtContractAddress := "0x55d398326f99059ff775485246999027b3197955" // USDT 在 BSC 上的合约地址
//
//	// 创建一个转账交易
//	sendMsg := msg.NewTokenTransferMsg(
//		address,
//		types.AccAddress(usdtContractAddress),
//		big.NewInt(100000000), // 以最小单位为单位（18位小数），即发送 1 USDT
//	)
//
//	// 创建交易编码器
//	txEncoder := txcodec.NewTxEncoder(client)
//
//	// 构建交易
//	txBuilder := tx.NewBuilder(txEncoder, keyManager.GetSigner())
//
//	// 签名交易
//	signedTx, err := txBuilder.BuildAndSignTx(context.Background(), usecase.NewTransferTxBuilder(client, txBuilder, []msg.Msg{sendMsg}))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 广播交易
//	res, err := client.Broadcast(context.Background(), signedTx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("Transaction Hash:", res.Hash)
//}
