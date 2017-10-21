package main

import (
	"fmt"
)

func digits(nb int, dch chan int) {
	for nb != 0 {
		digit := nb % 10
		dch <- digit
		nb /= 10
	}
	close(dch)
}
func computeSquares(nb int, sqch chan int) {
	sum := 0
	dch := make(chan int)
	go digits(nb, dch)
	for digit := range dch {
		sum += digit * digit
	}
	sqch <- sum
}

func main() {
	nb := 124
	sqch := make(chan int)
	go computeSquares(nb, sqch)
	squares := <-sqch
	fmt.Println("Sum of squares: ", squares)
}