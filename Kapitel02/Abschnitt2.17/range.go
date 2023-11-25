package main

import "fmt"

func main() {

	s := []string{"null", "eins"}
	for i, v := range s {
		fmt.Println(i, v)
	}

	m := make(map[int]string)
	m[10] = "zehn"
	m[100] = "hundert"
	for k, v := range m {
		fmt.Println(k, v)
	}
}
