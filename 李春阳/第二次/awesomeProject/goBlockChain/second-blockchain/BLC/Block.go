package BLC

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type Block struct{
	Height int64   			//1. 区块高度
	PrevBlockHash []byte	//2. 上一个区块HASH
	Data []byte				//3. 交易数据
	Timestamp int64			//4. 时间戳
	//Nonce int64				//5. Nonce
	Hash []byte				//6. Hash
}

func (block *Block) SetHash()  {
	//	1. Height []byte
	heightBytes := IntToHex(block.Height)
	//fmt.Println(heightBytes)

	//  2. 将时间戳转[]byte
	timeString := strconv.FormatInt(block.Timestamp,2)

	//fmt.Println(timeString)
	timeBytes := []byte(timeString)

	//fmt.Println(timeBytes)
	//  3. 拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes,block.PrevBlockHash,block.Data,timeBytes},[]byte{})

	// 4. 生成Hash
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]

}

func NewBlock(data string,height int64,prevBlockHash []byte) *Block {
	//创建区块
	block := &Block{height,prevBlockHash,[]byte(data),time.Now().Unix(),nil}
	//设置Hash
	block.SetHash()

	return block
}

//2. 单独写一个方法，生成创世区块

func CreateGenesisBlock(data string) *Block {
	return NewBlock(data,1, []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}