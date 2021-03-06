package main

import "fmt"

func main() {
  // ****************** Declaration *********************
  // Declaration:    var name[size]type
  // Initialization: var name[size]type = [size]type{values}
  // ShortHand:      name := [...]type{value}

  var a[3]int
  fmt.Println(a[1])

  var b[3]int = [3]int{1, 2, 3}
  fmt.Println(b[2])

  c := [...]int{1,2,3}
  fmt.Println(c[2])
  // fmt.Println(c[3])      // ERROR: Out of bound.
  d := [3]int{1,2,3}
  fmt.Println(d[2])
  
  // ****************** Length *********************
  fmt.Println(len(d))
  fmt.Println(d[len(d)-1])
 
  // ****************** Iteration *********************
  for i, v := range a {
    fmt.Printf("%d %d\n", i, v)
  }
  for _, v := range a {
    fmt.Printf("%d \n", v)
  }

  // ***************** Comparison *********************
  e := [2]int{1,2}
  f := [...]int{1,2}
  g := [2]int{1,3}
  fmt.Println(e==f, e==g, f==g)   // "true", "false", "false"
  h := [3]int{1,2}
  fmt.Println(e==h)               // mismatched types [2]int and [3]int

  
}
