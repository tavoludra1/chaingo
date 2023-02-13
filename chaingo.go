package main

import (
	"fmt"
	"log"
	"time"
)

// Creacion de la estructura del bloque
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

/*
Funcion para la creacion del bloque
@return: puntero del bloque
*/
func NewBlock(nonce int, previousHash string) *Block {
	bloque := new(Block)                     // almacenamiento de la estructura del bloque
	bloque.timestamp = time.Now().UnixNano() // int64
	bloque.nonce = nonce
	bloque.previousHash = previousHash
	return bloque
}

/*
Funcion para la impresion de los valores del bloque
*/

func (bloque *Block) Print() {
	fmt.Printf("timestamp	%d\n", bloque.timestamp)
	fmt.Printf("nonce	%d\n", bloque.nonce)
	fmt.Printf("previous_hash	%s\n", bloque.previousHash)
	fmt.Printf("trasactions	%s\n", bloque.transactions)

}

// *********************** INICIO ********************
func init() {
	log.SetPrefix("cadena_de_Bloques: ")
}

func main() {
	bloque := NewBlock(0, "Hash inicial")
	bloque.Print()

}
