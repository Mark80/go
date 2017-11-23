package main

import "fmt"


type Mutable struct {
  A string
  B string
}


func (m Mutable) StayTheSame(){
  m.A = "same"
  m.B = "same"
}


func (m *Mutable) Mutate(){
  m.A = "mutate"
  m.B = "mutate"
}

func main() {

  var m = &Mutable{"1","2"}
  m.StayTheSame()
  fmt.Println(m)

  m.Mutate()
  fmt.Println(m)



}