package main

import "fmt"

func main() {
  
  var PRICE int = 500
  const (
    PRODUCT string = "Shampoo"
  )
  PRICE = 600
  fmt.Println(PRODUCT)
  fmt.Println(PRICE)
}
