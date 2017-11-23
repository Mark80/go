package remotecommand


type Command struct {
  Name string
}

type Response struct {
  Value string
}

func receive(command Command) Response {
  var t = command.Name
  var result Response
  switch t {
  case "say":
	result = Say()
  }
  return result
}


func Say() Response {
  return Response{"Hello"}
}

