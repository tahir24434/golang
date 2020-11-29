package main

import (
  "os"
  "strconv"
)

func main() {
  res := 0
  for i:=1; i < len(os.Args); i++ {
     res += strconv.Atoi(os.Args[i])      // s = hello,world,something
  }
  //fmt.Println(res)
}

