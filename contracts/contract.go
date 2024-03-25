package contract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

func GetTransactionByHash(hash string) (map[string]interface{}, error) {
	rpcUrl := "https://bsc-dataseed1.binance.org"
	// 创建一个 RPC 客户端
	client, err := rpc.DialContext(context.Background(), rpcUrl)
	defer client.Close()

	var tx map[string]interface{}
	err = client.CallContext(context.Background(), &tx, "eth_getTransactionByHash", hash)
	if err != nil {
		log.Fatal("查询交易失败:", err)
	}

	// 输出查询结果
	fmt.Println("交易信息:", tx)
	// 解析JSON响应
	return tx, nil

}
func WithdrawUSDT(amount float64, address string) error {
	return nil
}
func WithdrawUNC(amount float64, address string) error {
	return nil
}
