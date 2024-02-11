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
}