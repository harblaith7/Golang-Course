package main

import "fmt"

func countdown(num int) {
	for i := num; i > 0; i-- {
		fmt.Println(i)
	}
}

func main(){
	countdown(10)
}