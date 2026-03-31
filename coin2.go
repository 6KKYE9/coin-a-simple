package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

func calculateHash(block Block) string {
	record := fmt.Sprintf(
		"%d%s%s%s",
		block.Index,
		block.Timestamp,
		block.Data,
		block.PrevHash,
	)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func generateBlock(oldBlock Block, data string) Block {
	newBlock := Block{
		Index:     oldBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  oldBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func main() {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "创世区块 - 江博的区块链",
		PrevHash:  "0",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)

	blockchain := []Block{genesisBlock}

	fmt.Println("🚀 开始自动挖矿，按 Ctrl+C 停止\n")

	// 无限循环自动挖矿
	for {
		lastBlock := blockchain[len(blockchain)-1]
		data := fmt.Sprintf("自动挖矿 - 第 %d 块", lastBlock.Index+1)
		newBlock := generateBlock(lastBlock, data)
		blockchain = append(blockchain, newBlock)

		// 打印一下刚挖出来的块
		fmt.Printf("✅ 已挖到区块 #%d | Hash: %.10s...\n", newBlock.Index, newBlock.Hash)

		// 每隔 1 秒挖一个，太快刷屏太猛
		time.Sleep(1 * time.Second)
	}
}