package main

import (
	"publicChain/part1-Basic-Prototype/BLC"
)

func main() {

	cli := &BLC.CLI{}

	cli.Run()

	//blockchain.AddBlockToBlcokchain("Send 100RMB To zhangqiang")
	//
	//blockchain.AddBlockToBlcokchain("Send 100RMB To cangjingkong")
	//
	//blockchain.AddBlockToBlcokchain("Send 100RMB To juncheng")
	//
	//blockchain.AddBlockToBlcokchain("Send 100RMB To haolin")
	//
	//blockchain.Printchain()
	//fmt.Println(blockchain)
	//fmt.Println(blockchain.Blocks)

	//block := BLC.NewBlock("Test",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1})
	//fmt.Printf("%d\n",block.Nonce)
	//fmt.Printf("%x\n",block.Hash)

	//db, err := bolt.Open("my.db", 0600, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = db.Update(func(tx *bolt.Tx) error {
	//	//尝试取表对象
	//	b := tx.Bucket([]byte("blocks"))
	//	if b == nil {
	//		b,err = tx.CreateBucket([]byte("blocks"))
	//		if err != nil {
	//			log.Panic("created table failed")
	//		}
	//	}
	//
	//	err = b.Put([]byte("l"),block.Serialize())
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//	return nil
	//})
	//if err != nil {
	//	log.Panic(err)
	//}

	//查看数据
	//err = db.View(func(tx *bolt.Tx) error {
	//	//尝试取表对象
	//	b := tx.Bucket([]byte("blocks"))
	//	if b != nil {
	//		blockData := b.Get([]byte("l"))
	//		block := BLC.DeserializeBlock(blockData)
	//		fmt.Printf("%v\n",block)
	//	}
	//
	//
	//
	//	return nil
	//})
	//defer db.Close()
	//更新数据库
	//err = db.Update(func(tx *bolt.Tx) error {
	//	//取表对象
	//	b := tx.Bucket([]byte("blocks"))
	//	if b == nil {
	//		b, err = tx.CreateBucket([]byte("blocks"))
	//		if err != nil {
	//			log.Panic("Blocks table create failed")
	//		}
	//	}
	//
	//	err = b.Put([]byte("l"), []byte(block.Serialize()))
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//
	//	return nil
	//})
	//if err != nil {
	//	log.Panic(err)
	//}

	//查看数据
	//err = db.View(func(tx *bolt.Tx) error {
	//	//创建BlockBucket
	//	b := tx.Bucket([]byte("blocks"))
	//	if b != nil {
	//		data := b.Get([]byte("l"))
	//		block := BLC.DeserializeBlock(data)
	//		fmt.Printf("%v\n", block)
	//	}
	//
	//	return nil
	//})
	////更新失败
	//if err != nil {
	//	log.Panic(err)
	//}

}
