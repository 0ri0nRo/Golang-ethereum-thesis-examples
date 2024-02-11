package main
import (
	"fmt"
	"log"	
	"context"
    "log"
	
	"github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
//    "github.com/ethereum/go-ethereum/crypto"
)

var(
	url = "https://sepolia.infura.io/v3/7f0d5493f09941789897ae6ea75788a0"
)

func main(){
	
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
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
	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatal(err)
	}
	// 1 ether = 1000000000000000000 wei
	amount := big.NewInt(100000000000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, a2, amount, 21000, gasPrice, nil)
	
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("wallet/UTC--2024-02-11T11-18-03.210022925Z--c4e5a877a603d740d1baf8d35206707160dc8cce")
	if err != nil {
		log.Fatal(err)
	}

	keystore.DecryptKey(b, "password")
	tx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}