package main

import "fmt"

func main(){
	purchases := [5]float32{19.99, 20.99, 5.99, 1.99, 14.55};

	income := [4]float32{20535.4, 4242.4, 52325.5}

	sales := [...]float32{424.44, 423.423}

	income[3] = 324234.5
	largestIncome := income[3]
	// income[4] = 423423
	fmt.Println(purchases, income, sales, largestIncome)

	for i := 0; i < len(purchases); i++ {
		fmt.Println(i)
	}
}