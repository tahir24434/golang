package main

import (
  "os"
  "fmt"
)

func main() {
  var s, sep string    // var VARIABLE-NAME VARIABLE-TYPE
  for i:=1; i < len(os.Args); i++ {
     s += sep + os.Args[i]      // s = hello,world,something
     sep = ","
  }
  fmt.Println(s)
}

// ./echo "hello world something"
// hello,world,something
// ["echo", "hello", "world", "something"]
