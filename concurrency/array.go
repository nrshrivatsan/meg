//Youtube URL : https://www.youtube.com/watch?v=XW1r1-4JI1s

package main

import (
	"fmt"
	"net/http"
)

func main() {

	//An Array of URLS from 500px.com
	arr := [20]string{"http://ppcdn.500px.org/53472962/016dc8008e0e1b0a71437a18ec4ef502c0e69cf6/5.jpg", "http://ppcdn.500px.org/53498178/0d13aa906538a8021568b948edfe3a3b3343dcf2/5.jpg", "http://ppcdn.500px.org/53495988/69d5a868beb9a4610e06ad1d64efa10204f8f45c/5.jpg", "http://ppcdn.500px.org/53488670/be88ee3404909ffb5e0f058133c8cb3cb3c61582/5.jpg", "http://ppcdn.500px.org/53472200/d23ece8bf9ab2f405b396edaf38a6ef77aab0c1e/5.jpg", "http://ppcdn.500px.org/53450310/977ea22beefa8c7148b3155bd829209a3ccf3330/5.jpg", "http://ppcdn.500px.org/53503408/5ed549003f8e919a3bf4b55f1e47dfa76ff77218/5.jpg", "http://ppcdn.500px.org/53496530/4a724f69bb4c1cb5e890594967572584f0c405cc/5.jpg", "http://ppcdn.500px.org/53493274/d55f2a1c0e639888d2ac7f20c04d360a18a75511/5.jpg", "http://ppcdn.500px.org/53478682/60a97a27221b471a00d2ea4c8a080cdf05350865/5.jpg", "http://ppcdn.500px.org/53472206/c99f0f1329f9ba908080cfb20def5fae433163da/5.jpg", "http://ppcdn.500px.org/53470904/5b99c63c3f1d1d19e01e73f49e88d72dc5776dca/5.jpg", "http://ppcdn.500px.org/53487620/514a8bb11650e53a25295e05f13e76069ffe982f/5.jpg", "http://ppcdn.500px.org/53471946/eb4244d152527e5f28199e0f822ffb6f9e7307a5/5.jpg", "http://ppcdn.500px.org/53512606/b46b81f293c18870bc537eb29af41176c364b177/5.jpg", "http://ppcdn.500px.org/53508262/b3f1845372b29b7721d45e16e8696e061a229e1c/5.jpg", "http://ppcdn.500px.org/53502718/3b1ab7097f8f1c6e32034f363680c1e08a37ab73/5.jpg", "http://ppcdn.500px.org/53499444/af5315c17755f7cfcc2c58023ecb553984882ce6/5.jpg", "http://ppcdn.500px.org/53497032/eb94e8f8780d3aca740d5e53ad0585419bb792d9/5.jpg", "http://ppcdn.500px.org/53481666/58ab7f6db578f75f2b9e1da485b180acebd8b8a9/5.jpg"}

	i := 0

	//Iterating over the array
	for ; i < len(arr); i++ {

		url := arr[i]

		go hit(url)
		//hit(url)

	}

	//Reading input from end user so as to see the result of execution of the go routines
	var input string
	fmt.Scanln(&input)
}

func hit(url string) {

	resp, err := http.Get(url)

	if err == nil {
		fmt.Println("Content Length ", (resp.ContentLength / 1024), " kb")

	}
}
