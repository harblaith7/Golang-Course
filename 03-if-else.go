package main

import "fmt"

func main(){
	age := 34

	if age >= 23 {
		fmt.Println("You are allowed in")
	} else {
		fmt.Println("Frigg off Ricky")
	}

	if age >= 23 {
		fmt.Println("You are allowed in")
	} else if age >= 20 {
		fmt.Println("Almost there")
	} else {
		fmt.Println("Frigg off Ricky")
	}
}