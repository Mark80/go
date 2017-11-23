package circuitsimulator

type Wire struct {
  value bool
  gates []*Inverter
}

func NewWire() *Wire {
  return &Wire{false, nil}
}

func (w *Wire) Set(value bool) {
  if len(w.gates) > 0 {
	for _ , gate :=  range w.gates{
	  gate.receive(value)
	}
  }
  w.value = value
}

func (w *Wire) Get() bool {
  return w.value
}

