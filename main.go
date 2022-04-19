package main

import "fmt"

func main() {

	a := 10
	fmt.Println(&a)

	var ponteiro *int = &a

	fmt.Println(ponteiro)
	fmt.Println(*ponteiro)
	fmt.Println(&ponteiro)

	*ponteiro = 50
	fmt.Println(*ponteiro)
	fmt.Println(a)

	b := &a
	fmt.Println(*b)
	*b = 60
	fmt.Println(*b)
	fmt.Println(*ponteiro)
	fmt.Println(a)

}
