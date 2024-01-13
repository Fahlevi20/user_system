package main

import "fmt"

func main() {

	var number = 10
	//pakai pointer
	var pointer *int = &number
	fmt.Println("Nilai pointer:", pointer)
	*pointer = 200

	fmt.Println("ini hasil variabel number", number)
}
