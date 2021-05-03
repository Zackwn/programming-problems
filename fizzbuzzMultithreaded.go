package main

import "fmt"

func fizz(c chan int, done chan bool) {
	for {
		i := <-c
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("fizz")
			done <- true
		}
	}
}

func buzz(c chan int, done chan bool) {
	for {
		i := <-c
		if i%5 == 0 && i%3 != 0 {
			fmt.Println("buzz")
			done <- true
		}
	}
}

func fizzbuzz(c chan int, done chan bool) {
	for {
		i := <-c
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
			done <- true
		}
	}
}

func number(c chan int, done chan bool) {
	for {
		i := <-c
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			done <- true
		}
	}
}

func spread(c, fc, bc, fbc, nc chan int) {
	for {
		x := <-c
		fc <- x
		bc <- x
		fbc <- x
		nc <- x
	}
}

func main() {
	fizzc := make(chan int)
	buzzc := make(chan int)
	fizzbuzzc := make(chan int)
	numberc := make(chan int)
	spreadc := make(chan int)
	done := make(chan bool)

	go fizz(fizzc, done)
	go buzz(buzzc, done)
	go fizzbuzz(fizzbuzzc, done)
	go number(numberc, done)

	go spread(spreadc, fizzc, buzzc, fizzbuzzc, numberc)

	x := 25
	for i := 0; i <= x; i++ {
		spreadc <- i
		<-done // wait done
	}
}
