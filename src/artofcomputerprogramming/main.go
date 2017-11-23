package main

func main() {

  mcd := algoritmoEuclideo(544,119)

  println(mcd)

}

func algoritmoEuclideo(a int, b int) int {

  resto := a % b

  for resto != 0 {
	a = b
	b = resto
	resto = a % b
  }

  return b

}