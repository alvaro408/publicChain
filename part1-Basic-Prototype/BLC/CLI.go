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
	fmt.Println("\tcreateblockchain -data - 交易数据")
	fmt.Println("\taddblock -data DATA - 交易数据")
	fmt.Println("\tprintchain --输出区块信息")
}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	if !DBExists() {
		fmt.Println("数据库不存在。。。")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()
	blockchain.AddBlockToBlcokchain(data)
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

func (cli *CLI) createGenesisBlockchain(data string) {
	CreateBlockchainWithGenesisBlock(data)
}

func (cli *CLI) Run() {
	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "www.google.com", "交易数据。。。")
	flagCreateBlockChainWithData := createBlockChainCmd.String("data", "Genesis block data...", "创世区块交易数据...")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
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

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		//fmt.Println(*flagAddBlockData)
		cli.addBlock(*flagAddBlockData)
	}

	if printChainCmd.Parsed() {
		//fmt.Println("输出所有区块的数据。。。")
		cli.printchain()
	}
	if createBlockChainCmd.Parsed() {
		if *flagCreateBlockChainWithData == "" {
			fmt.Println("交易数据不能为空...")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain(*flagCreateBlockChainWithData)
	}
}
