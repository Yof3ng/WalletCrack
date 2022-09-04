package main

import (
	"bufio"
	"container/list"
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
)

var client, ___ = ethclient.Dial(Web3API)
var Web3API = ""
var ThreadNum = 10
var TargetAccount = ""

func main() {
	flag.StringVar(&Web3API, "api", "", "API string, [https://mainnet.infura.io/v3/key] or [Local Node].")
	flag.IntVar(&ThreadNum, "t", 10, "Thread number, default 10.")
	flag.StringVar(&TargetAccount, "st", "", "Wallet address, query balance of target account.")
	flag.Parse()

	if Web3API == "" {
		fmt.Println("[!] usage: program -api apistring -t threadnum\n[!] usage: program -api apistring -st 0x00000000000-0")
		os.Exit(0)
	}
	client, ___ = ethclient.Dial(Web3API)

	if TargetAccount != "" {
		stValue, err := getBalance(common.HexToAddress(TargetAccount))
		if err != nil {
			log.Println(err)
		}
		fmt.Println("[" + TargetAccount + "] => " + "Balance:" + stValue)
		os.Exit(0)
	}

	fmt.Println("start...")
	for i := 1; i < ThreadNum; i++ {
		go crack()
	}
	crack()
}

func crack() {
	for {
		all := 0.0
		mnemonic, testTen := genWallet()
		fmt.Println("mnemonic: " + mnemonic)
		for i := testTen.Front(); i != nil; i = i.Next() {
			addr := fmt.Sprintf("%v", i.Value)
			account := common.HexToAddress(addr)
			value, err := getBalance(account)
			for err != nil {
				value, err = getBalance(account)
			}
			fmt.Println("[" + addr + "] => " + "Balance:" + value)
			s, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Println("Convert string to float failed!", err)
			} else {
				all = all + s
			}
		}
		if all > 0.0 {
			fmt.Println("Congratulations!!!!")
			filePath := "bonus.txt"
			file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
			if err != nil {
				fmt.Println("Open file failed!", err)
			}
			defer file.Close()
			write := bufio.NewWriter(file)
			write.WriteString(mnemonic + "\n")
			write.Flush()

		}
	}
}

func genWallet() (string, *list.List) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		log.Fatal(err)
	}
	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed := bip39.NewSeed(mnemonic, "")
	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}
	accountList := list.New()
	for i := 0; i < 3; i++ {
		flag := strconv.Itoa(i)
		path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + flag)
		account, err := wallet.Derive(path, false)
		if err != nil {
			log.Fatal(err)
		}
		address := account.Address.Hex()
		accountList.PushBack(address)
	}
	return mnemonic, accountList
}

func getBalance(account common.Address) (string, error) {
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	value := ethValue.String()
	return value, nil
}
