package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  // Join concatenates the elements of its first argument to create a single string. 
  // The separator string sep is placed between elements in the resulting string.
  fmt.Println(strings.Join(os.Args[1:], "--"));

  fmt.Println(os.Args[1:]);
}
