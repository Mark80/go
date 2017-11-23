package katas

import "test"
import "testing"
import "container/list"


func TestList(t *testing.T) {

  var l = list.New()

  l.PushFront("marco")

  test.AssertEquals("marco" , l.Front().Value , t)


}


