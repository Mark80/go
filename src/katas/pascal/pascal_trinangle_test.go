1package pascal

import "test"
import (
  "testing"
)

func TestDimension(t *testing.T) {

  var triangle = triangleLevel(1)

  test.AssertEquals(1 , len(triangle) ,t)
  test.AssertEquals(1 , len(triangle[0]) ,t)

}

func TestDimension2(t *testing.T) {

  var triangle = triangleLevel(3)

  test.AssertEquals(3 , len(triangle) ,t)
  test.AssertEquals(3 , len(triangle[0]) ,t)

}




func triangleLevel(level int) [][]int {
    var result = [][] int { [] int {1}}

	return result
}


