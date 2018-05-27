package main

import (
	"fmt"
	"encoding/hex"
	"encoding/base64"
)

var source string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

func main() {
	collection, err := hex.DecodeString(source)
	if err != nil {
	}
	fmt.Println(collection)
	var dest string = base64.StdEncoding.EncodeToString(collection)
	fmt.Println(dest)
}
