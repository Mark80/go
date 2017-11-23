package test

import (
  "testing"
  "strings"
  "strconv"
  "fmt"
)

func TestStringNil(t *testing.T) {

  var rangeStr string = "100"

  if strings.Contains(rangeStr, "-") {
	rangeStr = strings.Split(rangeStr, "-")[1]
  }

  var newrangeStr = strings.Replace(rangeStr, "%", "", -1)

  maxCollateralFactorRange, _ := strconv.ParseFloat(newrangeStr, 64)

  if maxCollateralFactorRange / 100 != 1 {

	t.Errorf("found ", rangeStr)

  }

}
