package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n   hex <hex>\nConverts hex to a string. Hex can include spaces.\n")
	}
	args := ""
	for _, v := range os.Args[1:] {
		args += v
	}
	hexData := strings.ReplaceAll(args, " ", "")
	if str, err := hex.DecodeString(hexData); err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Printf("%s\n", str)
	}

}
