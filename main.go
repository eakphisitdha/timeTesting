package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Measuring time in Go")
	start := time.Now()

	TimeSleep()

	fmt.Printf("The function took %s", time.Since(start))
}

func TimeSleep() {

	// Calling Sleep method
	time.Sleep(3 * time.Second)
}
