package main

import (
	"fmt"
	"os"
	"strings"
)

//define map of prefixes we can convert, as these are known.
var (
	__VERSION__ = "0.1"
	prefixMap   = map[string]string{
		"cosmosaccaddr": "cosmos",
		"cosmosaccpub":  "cosmospub",
		"cosmosvaladdr": "cosmosvaloper",
		"cosmosvalpub":  "cosmosvaloperpub",
	}
)

func welcome() {
	//Print a nice little welcome message
	fmt.Printf("\nBech32 convertor v%v\nBy Joshua Opaka.\n", __VERSION__)
}

func start(key string) {
	//Main application logic
	welcome()
	//Basic error checking
	if key == " " {
		usage()
	}
	if func() (elts []bool) {
		for _, x := range prefixMap {
			elts = append(elts, strings.HasPrefix(key, x))
		}
		return
	}() == nil {
		invalid(key)
	}
	//Decode existing key
	hrp, decodedPubkey := bech32decode(key)

	//Error if parsing legacy key failed
	if decodedPubkey == false {
		invalid(key)
	}
	//Determine prefix from map
	newPrefix := prefixMap[hrp]

	//Convert, and display result
	x, _ := bech32encode(newPrefix, decodedPubkey.([]int))
	fmt.Printf("Old Key: %s\nNew Key: %s\n", key, x)
}

func invalid(key string) {
	//Display warning for invalid format key
	fmt.Printf("ERROR: %v is an invalid legacy cosmos-sdk bech32 key\n", key)
	usage()
}

func usage() {
	//Print usage instructions
	fmt.Println("Usage:\n\nChange string variable in main file")
	os.Exit(255)
}

func bech32encode(hrp string, data []int) (string, error) {
	//Encode a bech32 string given prefix and bytearray
	converted, _ := convertbits(data, 8, 5, true)
	return Encode(hrp, converted)
}

func bech32decode(bech string) (string, interface{}) {
	//Decode bech32 string into prefix and bytearray of data
	hrp, data, _ := Decode(bech)
	//Return early if we couldn't parse the input strin
	if data == nil {
		return hrp, false
	}
	converted, _ := convertbits(data, 5, 8, false)
	return hrp, converted
}
