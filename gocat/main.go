package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hallo!")
	//fmt.Println(os.Args)
	if len(os.Args) == 1 {
		fmt.Println("Mindestens eine Datei als Parameter erwartet!")
		os.Exit(1)
	}
	for _, fname := range os.Args[1:] {
		fmt.Println(fname)
	}
}
