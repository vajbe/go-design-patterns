package main

import "fmt"

func main() {
	for i := 0; i < 30; i++ {
		go getInstanceSyncOnce()
		// go getInstance()
	}

	fmt.Scanln()
}
