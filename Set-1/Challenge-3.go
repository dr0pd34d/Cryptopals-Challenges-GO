// WIP!!!

package main

import (
	"fmt"
	"encoding/hex"
	"log"
	"strconv"
	"reflect"
)

// HEX string
var s1 string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
var common = [12]string {"E","T","A","O","I","N","S","H","R","D","L","U"}

func findkey(s1decode []byte) (){
	var xored []byte
	
	for x:=0; x < 256;x++ {
		xored = nil
		for i:=0; i < len(s1decode);i++ {
			byte1 := s1decode[i]
			byte2 := byte(0xx
			//fmt.Println(reflect.TypeOf(byte2))
			xored = append(xored, byte1 ^ byte2)
		}
		fmt.Println("Result with byte " + strconv.Itoa(x) + "...")
		var final string = hex.EncodeToString(xored)
		fmt.Println(final)
	}
}

func main() {
	s1decode, err := hex.DecodeString(s1)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(s1decode)
	findkey(s1decode)
}
