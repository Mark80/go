package main

import (
  "fmt"
  "time"
)

func stuff() chan string {
  ch := make(chan string)
  go func(channel chan string)\ {
	for i := 0; i < 10; i++ {
	  channel <- fmt.Sprintf("eccomi %d", i)
	}
  }(ch)
  return ch
}

func main() {

  ch1 := stuff()

   for i := 0; i < 10; i++ {
	fmt.Printf("%s\n" , <- ch1)
  }


  time.Sleep(time.Duration(500000000))
}