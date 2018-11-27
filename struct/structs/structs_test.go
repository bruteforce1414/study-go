package structs_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/bruteforce1414/study-go/struct/structs"
)

func TestPerson_GetAge(t *testing.T) {
	assertItem:=assert.New(t)
	testPerson:=structs.Person{}
	testPerson.SetAge(21)
	assertItem.Equal(21,testPerson.GetAge())
}

func TestPerson_GetName(t *testing.T) {
	assertItem:=assert.New(t)
	testPerson:=structs.Person{}
	testPerson.SetName("Ivan Ivanov")
	assertItem.Equal("Ivan Ivanov",testPerson.GetName())
}

func TestPerson_GetEmail(t *testing.T) {
	assertItem:=assert.New(t)
	testPerson:=structs.Person{}
	testPerson.SetEmail("ivan.ivanov@example.com")
	assertItem.Equal("ivan.ivanov@example.com",testPerson.GetEmail())
}