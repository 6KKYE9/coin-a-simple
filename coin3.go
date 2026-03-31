package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// 区块结构
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int // 挖矿用的随机数
}

// 难度：哈希前面要有几个 0
const difficulty = 6

// 计算哈希
func calculateHash(b Block) string {
	record := strconv.Itoa(b.Index) +
		b.Timestamp +
		b.Data +
		b.PrevHash +
		strconv.Itoa(b.Nonce)

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// 挖矿：找到满足难度的 Nonce
func mineBlock(oldBlock Block, data string) Block {
	newBlock := Block{
		Index:     oldBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  oldBlock.Hash,
		Nonce:     0,
	}

	// 暴力试 Nonce
	for {
		currentHash := calculateHash(newBlock)
		// 判断前缀是否有 difficulty 个 0
		prefixOK := true
		for i := 0; i < difficulty; i++ {
			if i >= len(currentHash) || currentHash[i] != '0' {
				prefixOK = false
				break
			}
		}
		if prefixOK {
			newBlock.Hash = currentHash
			break
		}
		newBlock.Nonce++
	}

	return newBlock
}

func main() {
	// 创世区块
	genesis := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "创世块 - 单机纯挖矿版",
		PrevHash:  "0",
	}
	genesis.Hash = calculateHash(genesis)

	chain := []Block{genesis}

	fmt.Printf("🚀 单机区块链启动，挖矿难度：%d\n", difficulty)
	fmt.Println("规则：哈希必须以 " + strconv.Itoa(difficulty) + " 个 0 开头")
	fmt.Println("======================================================")

	// 无限挖矿
	for {
		last := chain[len(chain)-1]
		data := fmt.Sprintf("第 %d 块：转账 100 枚 JBB 币", last.Index+1)

		// 开始挖
		newBlock := mineBlock(last, data)
		chain = append(chain, newBlock)

		// 输出结果
		fmt.Printf("✅ 成功挖出区块 #%d\n", newBlock.Index)
		fmt.Printf("Nonce: %d\n", newBlock.Nonce)
		fmt.Printf("Hash: %s\n\n", newBlock.Hash)
	}
}