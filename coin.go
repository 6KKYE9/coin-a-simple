package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// 区块结构
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// 生成区块哈希
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

// 创建新区块
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

// 主函数：启动区块链
func main() {
	// 创世区块（第一个块）
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block - YM coin",
		PrevHash:  "0",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)

	// 区块链列表
	blockchain := []Block{genesisBlock}

	// 造3个新区块
	for i := 1; i <= 3; i++ {
		newBlock := generateBlock(blockchain[len(blockchain)-1],
			fmt.Sprintf("第 %d 笔数据：转账 10 枚 Token", i))
		blockchain = append(blockchain, newBlock)
	}

	// 打印整条链
	fmt.Println("===== YM coin =====")
	for _, block := range blockchain {
		fmt.Printf("索引: %d\n数据: %s\n前哈希: %s\n当前哈希: %s\n\n",
			block.Index, block.Data, block.PrevHash, block.Hash)
	}
}