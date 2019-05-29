package main

import (
	"flag"
	"fmt"
)

func main() {

	flagString := flag.String("printchain", "", "输出所有区块")

	flagInt := flag.Int("number", 6, "输出一个整数。。。")

	flagBool := flag.Bool("open", false, "判断真假。。。")

	flag.Parse()

	fmt.Printf("%s\n", *flagString)
	fmt.Printf("%d\n", *flagInt)
	fmt.Printf("%v\n", *flagBool)
}
