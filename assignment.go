package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var noOfBlocks = 0
var root string
var tail string

type Block struct {
	no            int
	nonce         int
	previous_hash string
	hash          string
	transactions  []string
}

func NewBlock(blockchain []Block) *Block {
	block := new(Block)
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(10000)
	block.nonce = v

	var no int
	var transaction string

	fmt.Print("Enter the number of transactions you want to enter in block : ")
	fmt.Scan(&no)

	for i := 0; i < no; i++ {
		fmt.Printf("Enter transaction number %d : ", i+1)
		fmt.Scan(&transaction)

		block.transactions = append(block.transactions, transaction)
	}

	hashString := blocktoString(block)
	block.hash = CalculateHash(hashString)
	tail = block.hash
	if noOfBlocks == 0 {
		block.previous_hash = ""
		root = block.hash
	} else {
		block.previous_hash = blockchain[noOfBlocks-1].hash

	}

	return block
}

func blocktoString(block *Block) string {

	var hashString string
	hashString += strconv.Itoa(block.nonce) + block.hash + block.previous_hash

	for j := 0; j < len(block.transactions); j++ {
		hashString += block.transactions[j]
	}

	return hashString

}

func CalculateHash(stringToHash string) string {

	sum := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", sum)
}

func changeBlock(flag int, block *Block) Block {
	switch flag {
	case 2:
		var newString string
		var transaction string
		fmt.Print("Enter the new transaction : ")
		fmt.Scan(&transaction)
		block.transactions = append(block.transactions, transaction)
		newString = blocktoString(block)
		block.hash = CalculateHash(newString)
		fmt.Println("\nTransaction added!")
		break
	case 1:
		var newString string
		rand.Seed(time.Now().UnixNano())
		v := rand.Intn(10000)
		block.nonce = v
		newString = blocktoString(block)
		block.hash = CalculateHash(newString)
		fmt.Println("\nNonce changed!.")
		break
	default:
		fmt.Println("Invalid input!")
		break
	}

	return *(block)
}

func ListBlocks(blockchain []Block) {
	var length int
	length = len(blockchain)

	for i := 0; i < length; i++ {
		fmt.Println("\n\n\t", strings.Repeat("=", 25), " Block ", i+1, strings.Repeat("=", 25), "\n")
		fmt.Println("\n\tNonce : ", blockchain[i].nonce)
		fmt.Println("\n\t Hash : ", blockchain[i].hash)
		fmt.Println("\n\tPrevious Hash : ", blockchain[i].previous_hash)
		fmt.Println("\n\tTransactions : ")
		for j := 0; j < len(blockchain[i].transactions); j++ {

			fmt.Printf("\t\tNo.	%d\t", j+1)
			fmt.Println(blockchain[i].transactions[j])

		}

	}
}

func verifyBlockchain(blockchain []Block) int {
	length := len(blockchain)
	count := 0

	for i := length - 1; i >= 0; i-- {
		if i == 0 {
			if blockchain[i].hash == root {
				count++
			} else {
				fmt.Printf("Block %d changed!", i+1)
				return 0
			}
		} else if i == length-1 {
			if blockchain[i].hash == tail {
				count++
			} else {
				fmt.Printf("Block %d changed!", i+1)
				return 0
			}
		} else if blockchain[i].previous_hash == blockchain[i-1].hash {
			count++
		} else {
			fmt.Printf("Block %d changed!", i+1)
			return 0
		}
	}

	return 1
}

func menu() {
	fmt.Println("\n\n\t\t", strings.Repeat("=", 25), "Menu", strings.Repeat("=", 25))
	fmt.Println("Enter 1 to add Block.")
	fmt.Println("Enter 2 to print the Blockchain.")
	fmt.Println("Enter 3 to change a block.")
	fmt.Println("Enter 4 to verify blockchain.")
	fmt.Println("Enter 0 to Quit.")
}

func main() {

	blockchain := make([]Block, 0)

	var input int
	check := 1

	//var bl Block

	for check >= 0 {
		menu()
		fmt.Print("Enter a number : ")
		fmt.Scan(&input)

		switch input {
		case 0:
			fmt.Println("\nExiting...")
			fmt.Println("Done.")
			os.Exit(0)
			break
		case 1:
			bl := *(NewBlock(blockchain))
			blockchain = append(blockchain, bl)

			noOfBlocks++

			fmt.Println("Block added successfully.\n")
			break
		case 2:
			ListBlocks(blockchain)
			fmt.Println("\n\tEnd of Blockchain")
			break
		case 3:
			var flag int
			var bno int
			fmt.Print("Enter block number you want to change : ")
			fmt.Scan(&bno)
			for (bno > noOfBlocks) || (bno < 1) {
				fmt.Println("Invalid input!\n")
				fmt.Print("Enter block number you want to change : ")
				fmt.Scan(&bno)
			}
			fmt.Print("Enter the 1 to change nonce and 2 to add transactions : ")
			fmt.Scan(&flag)

			blockchain[bno-1] = changeBlock(flag, &blockchain[bno-1])

			break
		case 4:
			var check int
			check = verifyBlockchain(blockchain)
			if check == 1 {
				fmt.Println("Blockchain verified.")
			}
		default:
			fmt.Println("Invalid Input.\n")

		}

	}
}
