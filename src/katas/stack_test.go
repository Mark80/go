package katas

import "test"
import (
  "testing"
  "errors"
)

type Stack struct {
  values []int
}

func (stack Stack) Count() int {
  return len(stack.values)
}

func (stack *Stack) Push(value int) {

  stack.values = append(stack.values, value)
  return
}

func (stack *Stack) Pop() (int, error) {
  if stack.Count() == 0 {
	return -1, errors.New("empty stack")
  } else {
	value := stack.values[0]
	stack.values = stack.values[1:]
	return value, nil
  }
}

func TestEmptyStack(t *testing.T) {

  var stack = Stack{}

  test.AssertEquals(stack.Count(), 0, t)

}

func TestPushOneElement(t *testing.T) {

  var stack = Stack{}
  stack.Push(1)

  test.AssertEquals(stack.Count(), 1, t)

}

func TestPushTwoElement(t *testing.T) {

  var stack = Stack{}
  stack.Push(1)
  stack.Push(2)

  test.AssertEquals(stack.Count(), 2, t)

}

func TestPopValueFromEmptyStack(t *testing.T) {

  var stack = Stack{}

  _, err := stack.Pop()

  test.AssertEquals(errors.New("empty stack").Error(), err.Error(), t)

}

func TestPopValue(t *testing.T) {

  var stack = Stack{}
  stack.Push(25)
  value, err := stack.Pop()

  test.AssertEquals(nil, err, t)
  test.AssertEquals(25, value, t)
  test.AssertEquals(0, stack.Count(), t)

}

func TestPopValue2ElementStack(t *testing.T) {

  var stack = Stack{}
  stack.Push(25)
  stack.Push(5)

  value, err := stack.Pop()

  test.AssertEquals(nil, err, t)
  test.AssertEquals(5, value, t)
  test.AssertEquals(1, stack.Count(), t)

}
