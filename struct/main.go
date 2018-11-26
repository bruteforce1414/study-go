package main

import "fmt"

func main() {
	p := Person{}

	p.SetAge(21)
	p.SetName("John Doe")
	p.SetEmail("john.doe@example.com")

	fmt.Printf("%s %d : %s", p.GetName(), p.GetAge(), p.GetEmail())
}