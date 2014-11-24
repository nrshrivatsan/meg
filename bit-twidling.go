package main

import (
	"fmt"
	"math/big"
)

func main() {
	 
	var bitVector big.Int	
	
	var number int64 = 15
	var bitValue uint = 0
	var bitIndex int = 3
	
	//Setting the 3rd bit of 15 to zero 
	bitVector.SetBit(big.NewInt(number), bitIndex, bitValue)

	fmt.Printf("\nSetting the bit at %d of %b to %b \n",bitIndex, number, bitValue)

	fmt.Printf("\nThe new value %d => %b \n\n",bitVector.Uint64(), bitVector.Uint64())
}
