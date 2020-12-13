package main

import (
  "greet"
  "fmt"
)

func main() {
  greet.Greeting = "AoA"
  fmt.Println(greet.Hello("Areesha"))
  fmt.Println(greet.Farewell("areesha"))
}
