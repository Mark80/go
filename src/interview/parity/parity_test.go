package parity

import "testing"
import (
  "test"
  "fmt"
 _  "strconv"
)
//
//func TestParity(t *testing.T) {
//
//  var a  int64= 0XAC
//
//
//  fmt.Println(strconv.FormatInt(^a, 2))
//
//  fmt.Println(strconv.FormatInt(a, 2))
//
//  a = a | 0xF0
//
//  fmt.Println(strconv.FormatInt(0xF0, 2))
//  fmt.Println(strconv.FormatInt(a, 2))
//
//  test.AssertTrue(true ,t)
//
//}

func TestPointer(t *testing.T) {

  var a int
  var ptr *int

  fmt.Println(a , ptr)


  a = 300
  ptr = &a

  var prtOfPtr **int

  prtOfPtr = &ptr

  fmt.Println(a , ptr)

  fmt.Println(prtOfPtr)
  fmt.Println(**prtOfPtr)


  test.AssertTrue(true ,t)

}

