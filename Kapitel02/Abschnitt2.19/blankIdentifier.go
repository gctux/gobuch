package main

import "fmt"

func main() {

	a := []string{"eins", "zwei", "drei"}
	for _, v := range a {
		fmt.Println(v)
	}
}
