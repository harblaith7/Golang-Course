package main

import "fmt"

func main(){
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 1

	for j < 10 {
		fmt.Println(j)
		j++
	}

	k := 1
	for {
		fmt.Println(k)
		if k == 19 {
			break
		}
		k++
	}

	r := 1

	for r < 100 {
		if r % 2 == 0 {
			continue
		}
		fmt.Println(r)
		r++
	}
}