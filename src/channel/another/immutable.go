package another


type Immutable struct {

  name string
  lastName string
}

func (imm Immutable) GetName () string {
  return imm.name
}

