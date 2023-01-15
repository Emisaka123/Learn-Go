package main

import "fmt"

type users struct {
	name     string
	password string
}

func (u users) checkPassword(password string) bool {
	return u.password == password
}
func (u *users) resetPassword(password string) {
	u.password = password
}
func main() {
	a := users{"wang", "1024"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048"))
}
