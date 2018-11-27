package main

import
(
	"fmt"
	"github.com/bruteforce1414/study-go/struct/structs"
)

func main() {

	p := structs.Person{}
	p.SetAge(21)
	p.SetName("John Doe")
	p.SetEmail("john.doe@example.com")
	fmt.Printf("%s %d : %s", p.GetName(), p.GetAge(), p.GetEmail())
}