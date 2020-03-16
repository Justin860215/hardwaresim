package main

import "C"
import "fmt"

//export HelloFromGo
func HelloFromGo() {
	fmt.Println("Hello from Go.")
}
