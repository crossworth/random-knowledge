package main

import (
	"log"

	bolt "go.etcd.io/bbolt"
)

const (
	dbFile       = "blockchain.db"
	blocksBucket = "blocks"
)

type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

func NewBlockChain() *BlockChain {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var tip []byte

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				return err
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				return err
			}

			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	bc := &BlockChain{tip, db}
	return bc
}

func (b *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = b.db.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(blocksBucket))
		err := bk.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			return err
		}

		err = bk.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			return err
		}

		b.tip = newBlock.Hash

		return nil
	})
}

type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	bci := &BlockChainIterator{bc.tip, bc.db}

	return bci
}

func (i *BlockChainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}
