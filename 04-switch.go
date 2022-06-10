package main

import "fmt"

func main(){
	animal := "cat"

	switch animal {
		case "cat":
			fmt.Println("Meow")
		case "dog":
			fmt.Println("Woaf")
		case "frog":
			fmt.Println("Ribbit")
		case "horse":
			fmt.Println("Neigh")
		default:
			fmt.Println("Not Sure")
	}
}