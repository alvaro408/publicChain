package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{}

func printUsage() {

	fmt.Println("Usage:")
	fmt.Println("\tcreateblockchain -address - 交易数据")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT -- 交易明细")
	fmt.Println("\tprintchain --输出区块信息")
}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(txs []*Transaction) {
	if !DBExists() {
		fmt.Println("数据库不存在。。。")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.AddBlockToBlcokchain(txs)
}

func (cli *CLI) printchain() {
	if !DBExists() {
		fmt.Println("数据库不存在。。。")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.Printchain()
}

func (cli *CLI) createGenesisBlockchain(address string) {

	blockchain := CreateBlockchainWithGenesisBlock(address)
	defer blockchain.DB.Close()
}

func (cli *CLI) send(from []string, to []string, amount []string) {
	if !DBExists() {
		fmt.Println("数据库不存在。。。")
		os.Exit(1)
	}

	blockchain := BlockchainObject()
	defer blockchain.DB.Close()

	blockchain.MineNewBlock(from, to, amount)

}

func (cli *CLI) Run() {
	isValidArgs()

	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("", flag.ExitOnError)

	flagFrom := sendBlockCmd.String("from", "", "转账源地址")
	flagTo := sendBlockCmd.String("to", "", "转账目标地址")
	flagAmount := sendBlockCmd.String("amount", "", "转账金额")

	flagCreateBlockChainWithAddress := createBlockChainCmd.String("address", "", "创世区块地址...")

	switch os.Args[1] {
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}

	if sendBlockCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || *flagAmount == "" {
			printUsage()
			os.Exit(1)
		}
		//fmt.Println(*flagAddBlockData)
		//cli.addBlock([]*Transaction{})
		//fmt.Println(*flagFrom)
		//fmt.Println(*flagTo)
		//fmt.Println(*flagAmount)
		//
		//fmt.Println(JSONToArray(*flagFrom))
		//fmt.Println(JSONToArray(*flagTo))
		//fmt.Println(JSONToArray(*flagAmount))
		from := JSONToArray(*flagFrom)
		to := JSONToArray(*flagTo)
		amount := JSONToArray(*flagAmount)
		cli.send(from, to, amount)
	}

	if printChainCmd.Parsed() {
		//fmt.Println("输出所有区块的数据。。。")
		cli.printchain()
	}
	if createBlockChainCmd.Parsed() {
		if *flagCreateBlockChainWithAddress == "" {
			fmt.Println("地址不能为空...")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain(*flagCreateBlockChainWithAddress)
	}
}
