package main

import (
	"fmt"
	"index/suffixarray"
)

func main() {
	//A set of words delimited by space
	words := "a apple sphere atom atmosphere"

	//A suffix array created in golang by converting the given string into bytes 
	index := suffixarray.New([]byte(words))

	

	//Lookup Time complexity =  O(log(N)*len(s) + len(result))
	
	// N : the size of the indexed data
	// s : substring to be seached for
	// result : array containing integers which represent the index of the given substring 's' in the suffix array 
	
	//NOTE 
	// Let's take the following example 
	// var s string = "apple"
	// fmt.Println([]byte(s)) would print  = [97 112 112 108 101].
	// Golang Suffix array uses the byte representation of the sub-string "s" in order to perform most optimal comutation  


	offsets1 := index.Lookup([]byte("sphere"), -1) // the list of all indices where s occurs in data

	//Prints unsorted array of integers which are the indices of the given substring
	fmt.Println(offsets1)

}
