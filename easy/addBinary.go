package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}

func addBinary(a string, b string) string {
	aInt, bInt, sum := new(big.Int), new(big.Int), new(big.Int)
	aInt.SetString(a, 2)
	bInt.SetString(b, 2)
	sum.Add(aInt, bInt)
	return sum.Text(2)
}
