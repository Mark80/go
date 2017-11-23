package circuitsimulator

func NewInverter(input *Wire, output *Wire) *Inverter {
  output.value = !input.value
  inverter := &Inverter{input, output}
  input.gates = []*Inverter{inverter}
  return inverter
}

type Inverter struct {
  wireIn  *Wire
  wireOut *Wire
}

func (inv *Inverter) receive(value bool) {
  inv.wireOut.Set(!value)
}

