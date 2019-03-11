package main

import (
	"fmt"
	"publicChain/part1-Basic-Prototype/BLC"
)

func main() {

	genesisBlock := BLC.CreateGenesisBlock("Genesis Block")

	fmt.Println(genesisBlock)
}
