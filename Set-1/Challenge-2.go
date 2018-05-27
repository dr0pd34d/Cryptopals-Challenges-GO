package main

import (
	"fmt"
	"encoding/hex"
	"log"
)

// Both HEX strings
var s1 string = "1c0111001f010100061a024b53535009181c"
var s2 string = "686974207468652062756c6c277320657965"

func xor(s1decode,s2decode []byte) ([]byte){
	var xored []byte
	
	if(len(s1decode) == len(s2decode)){
		for i:=0; i < len(s1decode);i++ {
			byte1 := s1decode[i]
			byte2 := s2decode[i]
			xored = append(xored, byte1 ^ byte2)
			//fmt.Println(byte1 ^ byte2)
		}
	} else {
		log.Fatal("Buffer length does not match!")
	}
	return xored
}

func main() {
	s1decode, err := hex.DecodeString(s1)
	if err != nil {
		log.Fatal(err)
	}
	s2decode, err := hex.DecodeString(s2)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(s1decode)
	fmt.Println(s2decode)
	
	xored := xor(s1decode, s2decode)
	
	fmt.Println(xored)
	var final string = hex.EncodeToString(xored)
	fmt.Println(final)
}

