package main

import "fmt"

func main(){
	purchases := [5]float32{19.99, 20.99, 5.99, 1.99, 14.55};

	mySlice := purchases[:]
	// myOtherSlice = purchases[0:3]
	myOtherSlice := purchases[:3]
	myThirdSlice := purchases[2:]

	fmt.Println(mySlice, myOtherSlice, myThirdSlice)

	mySlice = append(mySlice, 1, 2, 3)

	fmt.Println(mySlice)

	combine := append(myOtherSlice, myThirdSlice...)

	fmt.Print(combine)
}
