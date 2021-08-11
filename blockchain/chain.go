package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/jihee102/explorer/db"
	"github.com/jihee102/explorer/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	utils.HandleErr(decoder.Decode(b))

}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			// Search for checkpoint on the db
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				//Restore b from bytes
				fmt.Println("Restoring")
				b.restore(checkpoint)
			}

		})
	}
	fmt.Printf("NewestHash: %s\nHeight:%d", b.NewestHash, b.Height)
	return b
}
