package naivechain

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"time"
)

type BlockChain struct {
	hash hash.Hash
	head *Block
	tail *Block
	len  uint64
}

type Block struct {
	index     uint64
	data      []byte
	timestamp time.Time
	hash      []byte
	prev      *Block
}

// New creates a new BlockChain and genesis Block with the data and Hash specified
func New(data []byte, hash hash.Hash) *BlockChain {
	// If no hash is set then default to SHA256
	if hash == nil {
		hash = sha256.New()
	}

	// Create a new BlockChain
	bc := BlockChain{
		hash: hash,
		len:  1,
	}

	// Create a new Block
	b := Block{
		index:     0,
		data:      data,
		timestamp: time.Now(),
		prev:      nil,
	}

	// Hash the Block and store the result on itself
	b.hash = bc.hashBlock(b)

	// Set the head and tail Blocks of the BlockChain to the new Block
	bc.head = &b
	bc.tail = &b

	return &bc
}

func (bc *BlockChain) hashBlock(block Block) []byte {
	data := append([]byte(string(block.timestamp.Unix())), block.data...)
	if block.prev != nil {
		data = append([]byte(block.prev.hash), data...)
	}
	data = append([]byte(string(block.index)), data...)
	bc.hash.Write(data)
	return bc.hash.Sum(nil)
}

// Len returns the current length of the BlockChain
func (bc *BlockChain) Len() uint64 {
	return bc.len
}

// Tail returns a copy of the last Block in the BlockChain
func (bc *BlockChain) Tail() Block {
	return *bc.tail
}

// Write creates a new Block with the provided data and adds it to the end of the BlockChain
func (bc *BlockChain) Write(data []byte) *BlockChain {
	// Create a new Block with the proper index, data, and prev Block pointer to the current tail
	b := Block{
		index:     bc.len,
		data:      data,
		timestamp: time.Now(),
		prev:      bc.tail,
	}

	// Hash the Block and store the result on itself
	b.hash = bc.hashBlock(b)

	// Set the tail of the BlockChain to the new Block
	bc.tail = &b

	// Increment the length of the BlockChain
	bc.len++

	return bc
}

// Print is a crude method to see how they are all getting chained together
func (bc *BlockChain) Print() {
	b := bc.tail
	for {
		if b == nil {
			break
		}
		fmt.Println("index: ", b.index)
		fmt.Println("data: ", string(b.data))
		fmt.Printf("hash: %x", b.hash)
		fmt.Println("")
		if b.prev != nil {
			fmt.Println("previous: ", string(b.prev.data))
		}
		fmt.Println("tail: ", string(bc.tail.data))
		fmt.Println("-----")
		b = b.prev
	}
}
