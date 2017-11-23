package main

import (
  "time"
  "log"
)

func expensiveFunction(num int , channel chan string){

  log.Println("start go routine number : " , num)
  time.Sleep(6000000000)
  channel <- "finito"
  log.Println("finish go routine number : " , num)

}

func main() {

  signal := make(chan string , 10)
  var start = time.Now()
  for i := 0 ; i<10 ; i++  {
	  go expensiveFunction(i, signal)
  }

  <- signal

  log.Println("finish in  ", time.Now().Second() - start.Second())


}
