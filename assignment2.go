package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var noOfBlocks = 0
var noOfTransactions = 0
var newHash string
var root string
var tail string

type Transaction struct {
	transactionID              string
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

type Block struct {
	no            int
	nonce         int
	previous_hash string
	hash          string
	transactions  []Transaction
}

func noncefinder(hash string) int {
	var nonce int
	check := 0
	nonce = 0
	for check == 0 {
		if hash[0] == '0' && hash[1] == '0' && hash[2] == '0' {
			check = 1
			newHash = hash
			break
		} else {
			nonce++
			hash = CalculateHash(hash + strconv.Itoa(nonce))
		}

	}
	fmt.Println("Block mined! with nonce : ", nonce)
	return nonce
}

func NewBlock(blockchain []*Block, transactions []*Transaction) *Block {
	block := new(Block)
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(10000)
	block.nonce = v

	no := 3

	for i := 0; i < no; i++ {
		block.transactions = append(block.transactions, *transactions[i])
	}

	hashString := blocktoString(block)
	block.hash = CalculateHash(hashString)

	tail = block.hash
	if len(blockchain) == 0 {
		block.previous_hash = ""
		root = block.hash
	} else {
		print("Length of blockchain : ", len(blockchain))
		block.previous_hash = blockchain[len(blockchain)-1].hash

	}
	fmt.Println("Mining block...")
	block.nonce = noncefinder(block.hash)

	block.hash = newHash

	return block
}

func blocktoString(block *Block) string {

	var hashString string
	hashString = strconv.Itoa(block.nonce) + block.hash + block.previous_hash

	for j := 0; j < len(block.transactions); j++ {

		hashString += block.transactions[j].transactionID + block.transactions[j].senderBlockchainAddress
		hashString += block.transactions[j].recipientBlockchainAddress + fmt.Sprintf("%v", block.transactions[j].value)
	}

	return hashString

}

func CalculateHash(stringToHash string) string {

	sum := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", sum)
}

func ListBlocks(blockchain []*Block) {
	var length int
	length = len(blockchain)
	if length == 0 {
		fmt.Println("\n\nBlockchain is empty!")
		return
	}

	for i := 0; i < length; i++ {
		fmt.Println("\n\n\t", strings.Repeat("=", 25), " Block ", i+1, strings.Repeat("=", 25), "\n")
		fmt.Println("\n\tNonce : ", blockchain[i].nonce)
		fmt.Println("\n\t Hash : ", blockchain[i].hash)
		fmt.Println("\n\tPrevious Hash : ", blockchain[i].previous_hash)
		printBlockTransactions(blockchain[i].transactions)

	}
}

func newTransaction(senderBlockchainAddress string, receiverBlockchainAddress string, value float32) *Transaction {
	var transaction Transaction
	transaction.transactionID = CalculateHash(senderBlockchainAddress + receiverBlockchainAddress + fmt.Sprintf("%v", value))
	transaction.senderBlockchainAddress = senderBlockchainAddress
	transaction.recipientBlockchainAddress = receiverBlockchainAddress
	transaction.value = value

	return &transaction
}

func (b *Blockchain) AddTransaction(senderBlockchainAddress string, receiverBlockchainAddress string, value float32) {
	transaction := newTransaction(senderBlockchainAddress, receiverBlockchainAddress, value)
	b.transactionPool = append(b.transactionPool, transaction)
}

type Blockchain struct {
	chain           []*Block
	transactionPool []*Transaction
}

func printBlockTransactions(transactions []Transaction) {

	if len(transactions) != 0 {
		fmt.Println("\n\tTransactions : ")
		for i := 0; i < len(transactions); i++ {
			fmt.Printf("\t\tNo.	%d\t", i+1)
			fmt.Println("\n\t\t\nTransaction ID : ", transactions[i].transactionID)
			fmt.Println("\t\tSender Blockchain Address : ", transactions[i].senderBlockchainAddress)
			fmt.Println("\t\tReceiver Blockchain Address : ", transactions[i].recipientBlockchainAddress)
			fmt.Println("\t\tAmount Transfered : ", transactions[i].value)
		}
	} else {
		fmt.Println("\n\t\tNo transactions in this block!")
	}
}
func printTransactionPool(transactions []*Transaction) {
	if len(transactions) != 0 {
		fmt.Println("\n\tTransaction Pool : ")
		for i := 0; i < len(transactions); i++ {
			fmt.Printf("\t\tNo.	%d\t", i+1)
			fmt.Println("\n\t\tTransaction ID : ", transactions[i].transactionID)
			fmt.Println("\t\tSender Blockchain Address : ", transactions[i].senderBlockchainAddress)
			fmt.Println("\t\tReceiver Blockchain Address : ", transactions[i].recipientBlockchainAddress)
			fmt.Println("\t\tAmount Transfered : ", transactions[i].value)
		}
	} else {
		fmt.Println("\n\t\tNo transactions in the pool!")
	}
}

func jsonEncode(block *Block) {

	blockToEncode := map[string]interface{}{
		"No":            block.no,
		"Nonce":         block.nonce,
		"Hash":          block.hash,
		"Previous Hash": block.previous_hash,
	}
	jsonBlock, err := json.Marshal(blockToEncode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBlock))
	fmt.Println("Transactions: ")
	for i := 0; i < len(block.transactions); i++ {
		transactionToEncode := map[string]interface{}{
			"Transaction ID":               block.transactions[i].transactionID,
			"Sender Blockchain Address":    block.transactions[i].senderBlockchainAddress,
			"Recipient Blockchain Address": block.transactions[i].recipientBlockchainAddress,
			"Value":                        block.transactions[i].value,
		}
		jsonTransaction, err := json.Marshal(transactionToEncode)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(jsonTransaction))
	}
}

func menu() {
	fmt.Println("\n\n\t\t", strings.Repeat("=", 25), "Menu", strings.Repeat("=", 25))
	fmt.Println("Enter 1 to add transactions.")
	fmt.Println("Enter 2 to print the Blockchain.")
	fmt.Println("Enter 3 to print transaction pool.")
	fmt.Println("Enter 4 to print a block with json encoding.")
	fmt.Println("Enter 0 to Quit.")
}

func main() {

	blockchain := new(Blockchain)
	var choice int
	check := 1

	for check >= 0 {
		menu()
		fmt.Print("Enter your choice : ")
		fmt.Scan(&choice)

		switch choice {
		case 0:
			check = -1
			break
		case 2:
			fmt.Println("Printing the Blockchain...")
			ListBlocks(blockchain.chain)
			break
		case 1:
			var senderBlockchainAddress string
			var receiverBlockchainAddress string
			var value float32

			fmt.Print("Enter sender's blockchain address : ")
			fmt.Scan(&senderBlockchainAddress)

			fmt.Print("Enter recipient's blockchain address : ")
			fmt.Scan(&receiverBlockchainAddress)

			fmt.Print("Enter value : ")
			fmt.Scan(&value)

			blockchain.AddTransaction(senderBlockchainAddress, receiverBlockchainAddress, value)

			if len(blockchain.transactionPool) == 3 {
				blockchain.chain = append(blockchain.chain, NewBlock(blockchain.chain, blockchain.transactionPool))
				blockchain.transactionPool = nil
			}
			noOfBlocks++
			break
		case 3:
			printTransactionPool(blockchain.transactionPool)
			break
		case 4:
			var blockNo int
			fmt.Print("Enter block number : ")
			fmt.Scan(&blockNo)
			jsonEncode(blockchain.chain[blockNo-1])
			break
		default:
			fmt.Println("Invalid input!")

		}
	}
}
