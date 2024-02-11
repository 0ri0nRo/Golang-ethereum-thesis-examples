package main
import (
	"fmt"
	"log"	
	"context"
    "log"
	
	"github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
)

var(
	url = "https://goerli.infura.io/v3/7f0d5493f09941789897ae6ea75788a0"
)

func main(){
	
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer cliend.Close()
	a1 := common.HexToAddress("c4e5a877a603d740d1baf8d35206707160dc8cce")
	a2 := common.HexToAddress("ada64c03bb8ecb118f71585df729b4ae1b1eb2ef")
	

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance 1: ", b1)
	fmt.Println("Balance 2: ", b2)

}