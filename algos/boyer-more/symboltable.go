package main

import "fmt"

func getSymbolTable(){

	s := "UnitedStatesOfAmerica"
	symbolTable := make([]string,0)
	for start:=0 ; start<len(s); start++ {
	  for end:=start; end <=len(s); end++{
			if start!=end {
				symbolTable = append(symbolTable,string(s[start:end]))
			}
	  }
	}

	for i,sub := range(symbolTable){
	   fmt.Println(i,sub)
	}
	//fmt.Println("Hello, playground")
}

func main() {
	getSymbolTable();
}
