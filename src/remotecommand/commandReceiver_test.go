package remotecommand

import (
  "testing"
)



func TestReceiveCommand(t *testing.T) {

  var commandSay = Command{"say"}
  var result = Response{"Hello"}

  if receive(commandSay) != result {
	t.Fail()
  }

}