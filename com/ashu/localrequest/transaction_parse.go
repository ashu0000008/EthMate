package localrequest

import "ashu.com/com/ashu/web3"

func Parse(txHash string) string {
	client, err := web3.Connect(web3.URL_NODE_ETH_PUBLIC)
	if nil != err {
		return ""
	}

	data := web3.GetTxData(client, txHash)
	print(data)
	return data
}
