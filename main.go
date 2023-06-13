package main

import (
	"bufio"
	"crypto/ed25519"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%20s: ", "Enter Seed")
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%20s: ", " Print key (yes/no)")
	adrPrint, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	printAdr := len(adrPrint) == 4 && adrPrint[:3] == "yes"

	fmt.Printf("%20s: ", "Number of keys")
	cntPrint, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	cnt, err := strconv.Atoi(cntPrint[:len(cntPrint)-1])
	if err != nil {
		fmt.Println(err)
		return
	}

	seed := sha512.Sum512([]byte(line))

	type key struct {
		public  []byte
		private []byte
		address []byte
	}
	var keys []*key

	for i := 0; i < cnt; i++ {
		k := new(key)
		keys = append(keys, k)
		key64 := ed25519.NewKeyFromSeed(seed[:32])
		k.private = key64[:32]
		k.public = key64[32:]
		address := sha256.Sum256(key64[32:])
		k.address = address[:]
		seed = sha512.Sum512(seed[:])
	}

	for _, k := range keys {
		if printAdr {
			fmt.Printf("private key:  %x \n", k.private)
		}
		fmt.Printf("public  key:  %x \n", k.public)
		fmt.Printf("hash of key:  %x\n\n", k.address)
	}
}
