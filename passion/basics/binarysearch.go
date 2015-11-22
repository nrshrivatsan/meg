package main

import (
  "math/rand"
  // "fmt"
  "time"
  "sort"
)

var values []int
var limit int

func p()  {
  for _,v := range(values){
    println(v)
  }
}

/*
  Generates random numbers and assigns them to
*/
func randr(r int) {

  for i := 0; i< r; i++ {
    // Commented code is the simple way
    s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)
    values = append(values,r.Intn(limit))
    /*Leveraging Golang Channels to generate random numbers*/
    // values = append(values, <-gen())

  }

}

/*
Generates random number and channels it out
*/
func gen() <-chan int {
  out:= make(chan int)
  go func(){
    /*
      a new pseudo-random Source seeded with Unix nano second as the seed
    */
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    /*Chanelling out the random generated number*/
    out <- r.Intn(limit)

    close(out)
  }()
  return out
}

func binsrch(search int)  {
   var found bool = false
   var index int = -1
  // TODO write sorting code yourself
  // TODO try quick sort
  // TODO try trim sort
  // TODO try Dual-Pivot Quicksort
  // SORTING ARRAY VALUES
  sort.Ints(values)
  lo := 0
  hi := len(values)-1
  for ;lo <= hi; {
   mid := lo + (hi - lo)/2
   if search < values[mid] {
     hi = mid -1
   }else if ( search > values[mid]){
     lo = mid+1
   }else{
     found = true
     index = mid
     break
   }
 }

 if found {
   println("Found",search,"@",index)
 }

  p()
}

func main() {
  //  println(<-gen())
    // values = make([]int, limit)
    limit = 10
    randr(limit)
    binsrch(7)
}
