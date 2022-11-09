package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	blockNum := big.NewInt(15928325)
	blockNumber := (blockNum)

	clientUrl := getEnvVar("CLIENT_URL")

	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Block data: %d\n", block)
	fmt.Printf("Block header: %v\n", header.Number.String())
	fmt.Printf("Tx count: %v\n", len(block.Transactions()))
	fmt.Printf("Block number: %v\n", block.Number().Uint64())
	fmt.Printf("Block difficulty: %v\n", block.Difficulty().Uint64())
	fmt.Printf("Block hash: %v\n", block.Hash().Hex())

}
