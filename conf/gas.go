package conf

import "C"
import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

var EthSuggestGas int64
var BscSuggestGas int64

func GasConfirm() {
	ethMainNet := new(Conf).GetConf().BSCManagerInfo.Dial
	ethClient, err := ethclient.Dial(ethMainNet)
	if err != nil {
		fmt.Println("get Client error")
	}
	bscMainNet := new(Conf).GetConf().BSCManagerInfo.Dial
	bscClient, err := ethclient.Dial(bscMainNet)
	if err != nil {
		fmt.Println("get Client error")
	}
	for {
		ethGasPrice, err := ethClient.SuggestGasPrice(context.Background())
		var ethGasPriceNew *big.Int
		ethGasPriceNew = big.NewInt(ethGasPrice.Int64() + 2*1000000000) // 额外加2gwei
		if err != nil {
			fmt.Println("get eth GasPrice error", err.Error())
		}
		EthSuggestGas = ethGasPriceNew.Int64()
		bscGasPrice, err := bscClient.SuggestGasPrice(context.Background())
		if err != nil {
			fmt.Println("get bsc GasPrice error", err.Error())
		}
		BscSuggestGas = bscGasPrice.Int64()
		fmt.Println("bsc gasPrice:", BscSuggestGas)
		fmt.Println("eth new gasPrice:", EthSuggestGas, "eth gasPrice:", ethGasPrice.Int64())
		time.Sleep(time.Second * 300)
	}

}
