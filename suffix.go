package main

import "sort"

//Type definition for constructing the suffix array
type node struct {
	char string;
	sa []node;
}

//Constructs the suffix array given the SuffixMap & the sorted set of unique symbols in the given string
func constructSA(suffixMap map[string]int , set []string) []node {
	sa := make([]node, 0) 
	counter := 0
	for k,v := range(suffixMap){			 
		n := node{ k, make([]node,0)}
		for i:=counter ; i<counter+v ; i++{
			// println("appending ",set[counter+i]," into ",n.char)
			n.sa = append(n.sa, node{set[counter+i], nil})
		}
		sa = append(sa,n)
		// println(k,v)
		counter++
	}
	return sa
}

// Constructs Suffix Map 
// suffixMap[character] -> number of suffixes
// NOTE This function needs to be optimized
func getSuffixMap(set []string) map[string]int {
	
	// println("\nSymbols from input")
	// println("-----------------")

	length := 0
	alpha := set[0]
	suffixMap := make(map[string]int)
	suffixMap[alpha] = length
	for _,x := range(set) {
		char := string(x[0])
		if suffixMap[char] >=0 {
			suffixMap[char] = suffixMap[char]+1
		}else{
			suffixMap[char] = 1
		}
		// suffixArray.append(node)
		// println(x)
	}
	return suffixMap
}


type suffixArray interface {	
	print() 
	construct(word string) []node
	getSet(word string) []string	
}


func (n node) print() {
	println(n.char)
	for i:=0 ; i< len(n.sa) ; i++ {
		println(" ",n.sa[i].char)
	}
	println("----------")	
}

func (n node) construct(word string) []node{
	set := n.getSet(word)
	suffixMap := getSuffixMap(set)
	sa := constructSA(suffixMap,set)
	return sa
}

func (n node) getSet(word string) []string {

	var set []string = make([]string, len(word))
	for i:=len(word)-1 ; i>=0 ; i-- {
		set[i]= word[i:]
	}
	sort.Strings(set)
	return set
}

func main() {
	word := "banana"
	
	// suffixMap := getSuffixMap(set)
	
	// sa := constructSA(suffixMap,set)	
	var n node;
	sa := n.construct(word)
	set := n.getSet(word)
	for _,x := range set {
		println(x)
	}
	for _,n := range(sa){
		n.print()
	}
	
}
