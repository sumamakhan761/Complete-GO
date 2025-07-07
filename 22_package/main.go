package main

import (
	"fmt"

	"github.com/sumamakhan761/Complete-GO/auth"
	"github.com/sumamakhan761/Complete-GO/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("sumama", "secret")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}

	// fmt.Println(user.Email, user.Name)
	color.Green(user.Email)
	color.Green(user.Name)

}