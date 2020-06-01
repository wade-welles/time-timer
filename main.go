package main

import (
	"fmt"

	time "./time"
)


func main() {
	fmt.Println("time tester")
	fmt.Println("===========")

	ticker := time.NewTicker(10 * time.Second) 
	ticker.Stop()
}
