package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//256位Hash里面前面至少要有16个0
const targetBit = 20

type ProofOfWork struct {
	Block  *Block   //当前要验证的区块
	target *big.Int //大数据存储
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PreBlockHash,
			pow.Block.HashTransactions(),
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{},
	)
	return data
}

func (proofOfWork *ProofOfWork) IsValid() bool {
	//1. proofOfWork.Block.Hash
	//2. proofOfWork.Block.Target
	var hashInt big.Int
	hashInt.SetBytes(proofOfWork.Block.Hash)

	if proofOfWork.target.Cmp(&hashInt) == 1 {
		return true
	}

	return false
}

func (ProofOfWork *ProofOfWork) Run() ([]byte, int64) {
	//1. 将Block的属性拼接成字节数组

	//2. 生成Hash

	//3. 去判断Hash有效性，如果满足条件，跳出循环
	nonce := 0

	var hashInt big.Int //存储新生成的hash
	var hash [32]byte

	for {
		//1. 将Block的属性拼接成字节数组
		dataBytes := ProofOfWork.prepareData(nonce)
		//2. 生成Hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		//将hash存储到hashInt
		hashInt.SetBytes(hash[:])
		//判断hashInt是否小鱼Block里面的target
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if ProofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}

		nonce = nonce + 1
	}

	return hash[:], int64(nonce)
}

//创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {

	//1. 创建一个初始值为1的target
	target := big.NewInt(1)

	//2. 左移256 - targetBit
	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}
