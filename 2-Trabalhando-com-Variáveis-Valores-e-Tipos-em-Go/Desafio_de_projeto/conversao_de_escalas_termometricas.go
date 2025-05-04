package main

import "fmt"

const tempC float64 = 100.0

func main() {
	tempK := tempC + 273.15
	fmt.Printf("Ponto de ebulição da água em Celsius: %g\nPonto de ebulição da água em Kelvin: %g", tempC, tempK)
}
