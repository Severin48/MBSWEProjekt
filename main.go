package main

import "fmt"

// Windows: go run $(Get-ChildItem -Path *.go -Name)
func main() {
	fmt.Printf("\nStarting task 1\n\n")
	Task1()
	fmt.Printf("\nStarting task 2\n\n")
	Task2()
	fmt.Printf("\nStarting task 3\n\n")
	Task3()
	fmt.Printf("\nStarting task 4\n\n")
	Task4()
}
