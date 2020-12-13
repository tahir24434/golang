package greet

import "fmt"

var Greeting = "Hello"
var farewell string = "bye"

func Hello(name string) string {
  return fmt.Sprintf("%s %s", Greeting, name)
}
