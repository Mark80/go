package main

import (
  "workqueues"
  "fmt"
  "net/http"
)


func main() {

  fmt.Println("starting dispatcheer")
  workqueues.StartDispatcher(1)


  fmt.Println("starting server")
  http.Handle("/work" , new(workqueues.HandlerWork))
  err := http.ListenAndServe(":8080" , nil)


  if err != nil {
    fmt.Println(err.Error())
  }
}
