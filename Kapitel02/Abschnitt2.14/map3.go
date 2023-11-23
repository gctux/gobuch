package main

import "fmt"

func main() {
	m := map[int]string{
		1: "Eintrag 1",
		2: "Eintrag 2",
	}
	fmt.Println(m)
}

// Output:
// map[1:Eintrag 1 2:Eintrag 2]
