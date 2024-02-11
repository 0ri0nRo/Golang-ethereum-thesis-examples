package main

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"

)

func main(){
	ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	_, err := ks.NewAccount("password")
	if err != nil{
		log.Fatal(err)
	}

	_, err = ks.NewAccount("password")
	if err != nil{
		log.Fatal(err)
	}
	//"c4e5a877a603d740d1baf8d35206707160dc8cce"
	//"ada64c03bb8ecb118f71585df729b4ae1b1eb2ef"
}

