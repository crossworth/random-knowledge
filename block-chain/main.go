package main

import (
	"fmt"
	"strconv"
)

// https://jeiwan.net/posts/building-blockchain-in-go-part-1/
// https://jeiwan.net/posts/building-blockchain-in-go-part-2/
// https://jeiwan.net/posts/building-blockchain-in-go-part-3/
func main() {
	bc := NewBlockChain()
	defer bc.db.Close()

	bc.AddBlock("User registered")
	bc.AddBlock("User viewed email")

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("PrevHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
