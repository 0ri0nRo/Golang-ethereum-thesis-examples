import "github.com/ethereum/go-ethereum/ethclient"

var(
	url = "https://goerli.infura.io/v3/7f0d5493f09941789897ae6ea75788a0"
)

package main
func main(){
	
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	defer cliend.Close()
	
}