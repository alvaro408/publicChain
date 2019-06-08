package BLC

import "fmt"

//查询余额
func (cli *CLI) getBalance(address string) {

	fmt.Println("地址：" + address)

	blockchain := BlockchainObject()
	defer blockchain.DB.Close()

	//txOutputs := blockchain.UnUTXOs(address)
	amount := blockchain.GetBalance(address)

	fmt.Printf("%s一共有%d个Token\n", address, amount)

	//fmt.Println("=======================")
	//for _, out := range txOutputs {
	//	fmt.Println(out)
	//}

}
