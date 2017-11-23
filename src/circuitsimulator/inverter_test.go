package circuitsimulator

import "testing"
import "test"

func TestInverterTrue(t *testing.T) {

  var wireInput = NewWire()
  var wireOutPut = NewWire()
  var inverter = NewInverter(wireInput, wireOutPut)

  wireInput.Set(true)

  test.AssertEquals(false, inverter.wireOut.value, t)

}

func TestInverterFalse(t *testing.T) {

  var wireInput = NewWire()
  var wireOutPut = NewWire()
  var inverter = NewInverter(wireInput, wireOutPut)

  wireInput.Set(false)

  test.AssertEquals(true, inverter.wireOut.value, t)

}
