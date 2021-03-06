package main

import "fmt"

func inc(a [3]int) {
  for i, _ := range a {
    a[i] = a[i] + 1
  }

  for _, v := range a {
    fmt.Println(v)
  }
}

func inc_original_array(a *[3]int) {
  for i, _ := range a {
    a[i] = a[i] + 1
  }
}

func main() {
  var a[3]int = [3]int{1,2,3}
  inc(a)
  fmt.Println("*************** After inc **********")
  for _, v := range a {
    fmt.Println(v)
  }

  inc_original_array(&a)
  fmt.Println("*************** After inc_original_array **********")
  for _, v := range a {
    fmt.Println(v)
  }
}

//a      = 1, 2, 3
//a_copy = 2, 3, 4

// a = 1,2,3  --- p=0xabcdef
// p_copy = 0xabcdef


