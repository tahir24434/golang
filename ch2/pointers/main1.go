package main

import (
  "fmt"
)

func main() {
  var x int
  var p *int
  x = 1
  p = &x
  fmt.Println(p, *p, x)
  *p = 2
  fmt.Println(p, *p, x)
  x = 3
  fmt.Println(p, *p, x)
}
