package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net"
	"os"
	"time"

	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

type Block struct {
	Index int
	Timestamp string
	BPM int
	Hash string
	PrevHash string
}

type Message struct {
	BPM int
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := fmt.Sprint(block.Index) + block.Timestamp + fmt.Sprint(block.BPM) + block.PrevHash
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)

}

func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index + 1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
}

var blockchainServer chan []Block

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	blockchainServer = make(chan []Block)

	// creating a genesis block
	t := time.Now()
	genesisBlock := Block{0, t.String(), 0, "", ""}
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)
	spew.Dump(Blockchain)

	// starting a TCP server
	server, err := net.Listen("tcp", ":" + os.Getenv("ADDR"))

	if err != nil {
		log.Fatal(err)
	}
	
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}