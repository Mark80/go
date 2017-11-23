package test

import "testing"

func AssertEquals(expected interface{}  , result interface{} , t *testing.T) {

  if expected != result {
	t.Errorf("expected \n%s , found \n%s", expected , result)
  }

}

func AssertNotEquals(expected interface{}  , result interface{} , t *testing.T) {

  if expected == result {
	t.Errorf("expected %s , found %s", expected , result)
  }

}

func AssertTrue(expected bool  ,  t *testing.T) {

  if !expected {
	t.Errorf("expected %s , found %s", expected , false)
  }

}

func AssertFalse(expected bool  ,  t *testing.T) {

  if expected {
	t.Errorf("expected %s , found %s", expected , true)
  }

}
