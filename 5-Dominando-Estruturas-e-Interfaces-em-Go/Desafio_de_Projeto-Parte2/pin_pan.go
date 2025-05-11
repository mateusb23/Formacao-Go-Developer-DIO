package main

import "fmt"

func main() {
	fmt.Printf(`Para os números de 1 a 100, imprima: 
	 Pin se o número for divisível por 3 
	 Pan se o número for divisível por 5 
	 PinPan se o número for divisível por 3 e 5 
	 O próprio número caso contrário.`)
	pinPan()
}

func pinPan() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
