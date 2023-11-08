package main

import "fmt"

type rechteck struct {
	laenge int
	breite int
}

func (r *rechteck) setBreite(b int) {
	r.breite = b
}

// Funktion
//func flaeche(r rechteck) int {
//	return r.laenge * r.breite
//}

func (r rechteck) flaeche() int {
	return r.laenge * r.breite
}

func main() {
	r := rechteck{laenge: 10, breite: 5}
	r.setBreite(100)
	fmt.Println(r.flaeche())
}

// Output: 1000
