package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"meta-mall/database"
	"meta-mall/model"
	"meta-mall/pkg/chain"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type EH struct {
	ctx context.Context
	db  *gorm.DB
}

func ScanEth(db *gorm.DB) {
	for {
		err := database.DB.Transaction(func(tx *gorm.DB) error {
			var txs = model.Transactions{}
			txs.ChainName = "eth"
			untreatedTxsList, err := txs.GetUntreatedTxs(tx)
			if err != nil {
				return err
			}
			for _, transaction := range untreatedTxsList {
				if transaction.UpdatedAt.Add(time.Minute * 5).Before(time.Now()) {
					txsInChain, err := getTransactionByHash(transaction.Hash)
					if err != nil {
						return err
					}
					if txsInChain.Result.Value != "" {
						transaction.Status = "2"
					} else {
						transaction.Status = "1"
					}
					err = transaction.UpdateTransactions(tx)
					if err != nil {
						return err
					}

				}
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
}
func getTransactionByHash(hash string) (tron.Transaction, error) {
	url := "https://mainnet.infura.io/v3/a936bfa4553a4a95862326edddc46306"
	data := []byte(`{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":[` + hash + `],"id":1}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	transaction := tron.Transaction{}
	if err != nil {
		return transaction, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return transaction, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &transaction)
	if err != nil {
		return tron.Transaction{}, err
	}
	return transaction, nil
}
