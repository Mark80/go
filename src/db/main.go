package main

import (
  "crypto/sha256"

  "fmt"
)

func main() {

  input := []byte("ASDFGHJKUGCCDSXCD")
  result := sha256.Sum256(input)

  fmt.Println(sha256.Size)

  fmt.Println(result)


}
