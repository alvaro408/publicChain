package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
)

type Transaction struct {
	//1.交易hash
	TxHash []byte
	//2.输入
	Vins []*TXInput
	//3.输出
	Vouts []*TXOutput
}

//1.创建创世区块时的transaction
func NewCoinbaseTransaction(address string) *Transaction {
	//代表消费
	txInput := &TXInput{[]byte{}, -1, "Genesis Block"}

	txOutput := &TXOutput{10, address}

	txCoinbase := &Transaction{[]byte{}, []*TXInput{txInput}, []*TXOutput{txOutput}}
	//设置哈希值
	txCoinbase.HashTransaction()

	return txCoinbase
}

//将
func (tx *Transaction) HashTransaction() {
	//创建缓冲区
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())

	tx.TxHash = hash[:]
}

//2.转账时的transaction
func NewSimpleTransaciton(from string, to string, amount int) *Transaction {

	var txInputs []*TXInput
	var txOutputs []*TXOutput

	//代表消费
	bytes, _ := hex.DecodeString("a4362fed91252cbfc9c98b1320c7253d0d19505a1deb9b9354780c2908127f0f")
	txInput := &TXInput{bytes, 0, from}

	//消费
	txInputs = append(txInputs, txInput)

	//转账
	txOutput := &TXOutput{int64(amount), to}
	txOutputs = append(txOutputs, txOutput)

	//找零钱
	txOutput = &TXOutput{10 - int64(amount), from}
	txOutputs = append(txOutputs, txOutput)

	tx := &Transaction{[]byte{}, txInputs, []*TXOutput{txOutput}}
	//设置哈希值
	tx.HashTransaction()

	return tx

}
