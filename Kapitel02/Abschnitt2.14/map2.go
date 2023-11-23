package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "Eintrag 1"
	fmt.Println(m)
}

// Output:
// map[1:Eintrag 1]
