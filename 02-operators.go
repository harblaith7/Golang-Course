package main

import "fmt"

func main(){
	// ARITHMETIC OPERATORS
	var num1 = 4;
	var num2 = 3;

	var sum = num1 + num2
	var difference = num1 - num1
	var product = num1 * num2
	var quotient = num1 / num2
	var remainder = num1 % 2

	// sum = sum + 1
	// sum++
	sum+= 5
	fmt.Println(sum, difference, product, quotient, remainder)


	// RELATIONAL OPERATORS

	var result = num1 > num2
	result = num1 >= num2 
	result = num1 < num2
	result = num1 <= num2
	result = num1 == num2 
	result = num1 != num2 

	fmt.Println(result)

	// LOGICAL OPERATORS

	const name = "Laith"
	var age = 25

	var inviteToParty = name == "Laith" && age > 23
	inviteToParty = name == "Laith" || age >= 21
	inviteToParty = name == "Laith" || (age >= 21 && age < 90)
	inviteToParty = !(name == "Laith") && (age >= 21 && age < 90)
	
	fmt.Println(inviteToParty)
}

