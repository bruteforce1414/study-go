package request_test

import (
	"github.com/bruteforce1414/study-go/request/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUserRequest_Validate(t *testing.T) {
	assertUserRequest:=assert.New(t)
	testUser:=request.CreateUserRequest{Name:"Petrov Petr",

		Email:"petr.petrov@gmail.com",
		Password:"1234"}
	    err:=testUser.Validate()
	    assertUserRequest.NoError(err)

	    testUser2:=request.CreateUserRequest{}
		err2:=testUser2.Validate()
	    assertUserRequest.Error(err2)

	}
