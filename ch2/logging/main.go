package main

import "github.com/tahir24434/go-packages/logging"

func main() {
    logging.Debug(true)

    logging.Log("This is a debug statement...")
    logging.Log("This is another one ...")
}
