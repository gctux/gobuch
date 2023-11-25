package main

import "fmt"

func main() {

	for i := 0; i <= 4; i++ {
		fmt.Println(i)
	}

	var j int
	for j <= 4 {
		fmt.Println(j)
		j++
	}

	var k int
	for {
		if k > 4 {
			break
		}
		fmt.Println(k)
		k++
	}
}
