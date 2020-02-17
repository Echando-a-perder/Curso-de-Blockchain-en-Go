package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

//Estructura del bloque
type Block struct {
	//Declaramos la variable de la fecha como int64
	Timestamp int64
	//Declaramos los datos del bloque en formato byte
	Data []byte
	//Declaramos el hash del bloque anterior para entrelazarlos
	PrevBlockHash []byte
	//Declaramos el hash del bloque actual
	Hash []byte
}

//Por ahora, solo tomaremos campos de bloque, los concatenaremos y calcularemos un hash SHA-256 en la combinación concatenada. Hagamos esto en el SetHashmétodo
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

//implementaremos una función que simplificará la creación de un bloque
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

//Declaramos la variable de la cadena de bloques
type Blockchain struct {
	blocks []*Block
}

//La funcion para agregar un bloque en la cadena
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

//Metodo para crear el primer bloque de la cadena llamado "Bloque Genesis"
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//Implementar una función que cree una cadena de bloques con el bloque génesis
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

//Programa principal
func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
