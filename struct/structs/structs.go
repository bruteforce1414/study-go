// This example demonstrate how to use private struct fields 
//
//
package structs

type Person struct {
	name string
	email string
	age int
}

func (P *Person) GetName() string {
	return P.name
}

func (P *Person) GetEmail() string {
	return P.email
}

func (P *Person) GetAge() int {
	return P.age
}

func (P *Person) SetName(name string) {
	P.name=name;
}

func (P *Person) SetEmail(email string) {
	P.email=email;
}

func (P *Person) SetAge(age int) {
	P.age=age;

}