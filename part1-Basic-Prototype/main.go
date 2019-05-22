package main

import (
	"fmt"
	"publicChain/part1-Basic-Prototype/BLC"
)

func main() {
	//创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()

	//新区快
	blockchain.AddBlockToBlcokchain("Send 100RMB To zhangqiang", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].PreBlockHash)

	blockchain.AddBlockToBlcokchain("Send 100RMB To cangjingkong", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].PreBlockHash)

	blockchain.AddBlockToBlcokchain("Send 100RMB To juncheng", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].PreBlockHash)

	blockchain.AddBlockToBlcokchain("Send 100RMB To haolin", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].PreBlockHash)

	fmt.Println(blockchain)
	fmt.Println(blockchain.Blocks)
}
