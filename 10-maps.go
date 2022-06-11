package main

import "fmt"

func main(){
	cart := make(map[string]int)

	cart["lamp"] = 5
	cart["bowl"] = 1
	cart["laptop"] += 1

	delete(cart, "laptop")

	fmt.Println(cart, cart["lamp"], cart["stephen"])

	lamp, found := cart["book"]

	fmt.Println(lamp, found)

}