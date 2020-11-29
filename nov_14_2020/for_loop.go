package main

import (
  "fmt"
  "os"
)

func main() {
  // os.Args = [for_loop, a, b, c, d]
  for i:=1; i<len(os.Args); i++ {
    fmt.Println(i, os.Args[i]);
  }

  fmt.Println("------");

  // os.Args = [for_loop, a, b, c, d]
  for index, val := range os.Args[2:] {
    fmt.Println(index, val);
  }
}
