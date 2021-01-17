package main

import (
  "fmt"
)


func incr(p int) {
  p++ // increments what p points to; does not change p
}

func incrWithPointer(p *int) {
  *p++ // increments what p points to; does not change p
}

func main() {
  u := 1
  incr(u)
  fmt.Println(u)

  v := 1
  incrWithPointer(&v)              // side effect: v is now 2
  fmt.Println(v) 
}
