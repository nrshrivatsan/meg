package main

import "fmt"

func bits() {

	for i:= 0; i<=15; i++ {
	
	  //Prints first 4 bits representing the given integer
		fmt.Printf("%.4b\n",i)
		
	}
	
}

func main() {
	bits()	
}
