package main

import "fmt"

func main(){
	type Animal struct{
		class	string
		age		int	
		gender	bool	
	}

	var teddy = Animal{
		class: "bear",
		age: 24,
		gender: true,
	}

	fmt.Println(teddy.age, teddy.class)

	teddy.age = teddy.age + 1

	fmt.Println(teddy.age, teddy.class)

	var leo = Animal{"lion", 2, false}

	var lalo = Animal{}

	var tuco struct {
		class	string
		age		int	
		gender	bool	
	}

	tuco.class = "gorilla"

	nacho := struct {
		class	string
		age		int	
		gender	bool	
	} {
		class: "Penguin",
		age: 1,
		gender: false,
	}
	fmt.Println(teddy, leo, lalo, tuco, nacho)
}