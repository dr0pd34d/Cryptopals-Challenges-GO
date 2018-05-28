package main

import (
	"fmt"
	"encoding/hex"
	"log"
	"strconv"
	"strings"
)

// HEX string
var sourceString string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
var common = [12]string {"E","T","A","O","I","N","S","H","R","D","L","U"}

func max(arr []int) (int){
	max := arr[0] // assume first value is the smallest

	for _, value := range arr {
		if value > max {
			max = value // found another smaller value, replace previous value in max
		}
	}
	return max
}

func findKey(s1decode []byte) (){
	var probability []int
	var countCommon int

	for x:=0; x < 256; x++ {
		var final string = decrypt(x,s1decode)
		countCommon = 0
		for z:=0; z < len(common);z++ {
			countCommon += strings.Count(final, common[z])
		}
		probability = append(probability, countCommon)
	}

	maxvalue := max(probability)
	fmt.Println("Maximum probability is " + strconv.Itoa(maxvalue))
	for i := range probability {
		if probability[i] ==  maxvalue{
			// Found!
			fmt.Println("Key for decoding should be byte " + strconv.Itoa(i))
			fmt.Println(decrypt(i,s1decode))
		}
	}
}

func decrypt(key int, source []byte, ) (string){
	var xored []byte
	for i:=0; i < len(source);i++ {
		byte1 := source[i]
		byte2 := byte(key)
		xored = append(xored, byte1 ^ byte2)
	}
	return string(xored)
}

func main() {
	s1decode, err := hex.DecodeString(sourceString)
	if err != nil {
		log.Fatal(err)
	}

	findKey(s1decode)
}
