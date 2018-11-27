package main

import (
	"fmt"
	"github.com/bruteforce1414/study-go/request/request"
)
func main()  {

	firstUSer:=request.CreateUserRequest{Name:"Ivanov Ivan",

	                                     Email:"ivan.ivanov@gmail.com",
	                                     Password:"qwerty"}
		err:=firstUSer.Validate()
		if err==nil {
			fmt.Println("Валидация прошла успешна")
		}else{
			fmt.Println("Ошибка валидации")
		}
}