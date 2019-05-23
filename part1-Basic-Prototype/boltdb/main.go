package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//创建表
	/*	err = db.Update(func(tx *bolt.Tx) error {
			//创建BlockBucket
			b := tx.Bucket([]byte("BlockBucket"))

			if b != nil {
				err := b.Put([]byte("l"), []byte("Send 10 BTC To 哥"))
				if err != nil {
					log.Panic("数据库存储失败......")
				}
			}

			return nil
		})
		//更新失败
		if err != nil {
			log.Panic(err)
		}
	*/
	//创建表
	err = db.View(func(tx *bolt.Tx) error {
		//创建BlockBucket
		b := tx.Bucket([]byte("BlockBucket"))
		if b != nil {
			data := b.Get([]byte("l"))
			fmt.Printf("%s\n", data)
			data = b.Get([]byte("ll"))
			fmt.Printf("%s\n", data)
		}

		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}
}
