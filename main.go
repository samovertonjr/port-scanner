package main

import (
	"fmt"

	"port-scanner/port"
)


func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
