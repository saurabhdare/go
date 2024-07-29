package main

import "fmt"

func zero(x *int) {
	*x = 10
}

func one(xPtr *int) {
	*xPtr = 51
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}
