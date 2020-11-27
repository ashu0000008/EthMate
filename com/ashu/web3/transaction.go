package web3

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func Connect(host string) (*ethclient.Client, error) {
	ctx, err := rpc.Dial(host)
	if err != nil {
		return nil, err
	}
	conn := ethclient.NewClient(ctx)
	return conn, nil
}

func GetTxData(client *ethclient.Client, txHash string) string {
	hash := common.BytesToHash([]byte(txHash))
	tx, _, err := client.TransactionByHash(context.Background(), common.Hash(hash))
	if nil == err {
		return ""
	}
	return string(tx.Data())
}
