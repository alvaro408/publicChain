package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	//1. 区块高度
	Height int64
	//2. 上一个区块的hash
	PreBlockHash []byte
	//3. 交易数据
	Data []byte
	//4. 时间戳
	Timestamp int64
	//5. hash
	Hash []byte
}

//1. 创建新的区块
func NewBlock(data string, height int64, preBlockHash []byte) *Block {
	//1. 生成区块
	block := &Block{height, preBlockHash, []byte(data), time.Now().Unix(), nil}
	//2. 设置hash
	block.SetHash()

	return block
}

//2. 生成创世区块
func CreateGenesisBlock(data string) *Block {

	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func (block *Block) SetHash() {
	//1. Height转换为字节数组
	heightBytes := IntToHex(block.Height)
	//2. 时间戳转换为字节数组
	timeString := strconv.FormatInt(block.Timestamp, 2)
	timeByte := []byte(timeString)
	//3. 拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes, block.PreBlockHash, block.Data, timeByte, block.Hash}, []byte{})
	//4. 生成Hash
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}
