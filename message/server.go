package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("--- message ---")
		time.Sleep(10 * time.Second)
	}
}
