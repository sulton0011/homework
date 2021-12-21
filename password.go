package main

import (
	"fmt"
	"unicode/utf8"
)
func main() {
	fmt.Println(PasswordVerification("asdfamps"))
}

func PasswordVerification(s string) string{
	if len(s) == utf8.RuneCountInString(s) && utf8.RuneCountInString(s) >= 5{
		return "Ok"
	}else {
		return "Wrong password"
	}
}