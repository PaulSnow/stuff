package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"os"
)

func main() {
	if len(os.Args)< 2 {
		fmt.Printf("Usage:\n   sha256 <hex>\nreturns hash of the given hex string.\n")
	}
	args :=""
	for _,v := range os.Args[1:]{
		args+=v
	}
	hexData := strings.ReplaceAll(args, " ", "")
	bytes, err := hex.DecodeString(hexData)
	if err != nil {
		fmt.Print("arguments do not represent a valid hex string.\n")
		os.Exit(1)
	}
	hash := sha256.Sum256(bytes)
	fmt.Printf("%x\n", hash)
}
