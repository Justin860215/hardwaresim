package main

/*
extern int helloFromC();
*/
import "C"

func main() {
	C.helloFromC()
}
