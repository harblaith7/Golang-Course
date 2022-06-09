package main

import "fmt"

func main(){
	// declaring and assigning a variable
	var favBook = "Harry Potter"

	fmt.Println(favBook)

	// reassign
	favBook = "Power of Habits" 

	fmt.Println(favBook)

	// cannot do
	// favBook = 12

	// explicted defining the type
	var otherFavBook string = "Bad Blood"
	// var otherFavBook int = "Bad Blood"

	fmt.Println(otherFavBook)

	var thirdFavBook string
	var myAge int
	var amICool bool

	fmt.Println(thirdFavBook, myAge, amICool)
	
	thirdFavBook = "Diary of Wimpy Kid"

	fmt.Println(thirdFavBook)

	// COMPOUND CREATION

	// var favNumber, favChocolate, favTeam = 27, "KitKat", "Knicks";

	// BLOCK CREATION
	var (
		favNumber = 27
		favChocolate = "KitKat" 
		favTeam = "Knicks"
	)

	fmt.Println(favNumber, favChocolate, favTeam)

	// shortcut to declare and assign

	favAnimal := "Tiger"

	favAnimal = "Lion"

	fmt.Println(favAnimal)

	pet1, pet2, pet3 := "cat", "dog", "rat"

	pet2 = "parrot"

	// pet2 := "monkey"

	pet4, pet2 := "monkey", "bird"

	fmt.Println(pet1, pet2, pet3, pet4)

	const MyName = "Laith Harb"
	// MyName = "asdasda"
}

