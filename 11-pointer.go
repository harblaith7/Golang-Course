package main

import "fmt"

func main(){
	var a = 1

	var b = a

	b++

	var c = 1

	var d *int

	d = &c

	*d += 1

	fmt.Println(a, c)

	
}