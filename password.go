package main

import (
	"fmt"
	"unicode/utf8"
)
func main() {
	fmt.Println(PasswordVerification("asdfampлs"))
}

func PasswordVerification(s string) string{
	var PassVer bool

	for i := range s {
		if s[i] <= 90 && s[i] >= 65 || s[i] >= 97 && s[i] <= 122 {
			fmt.Println("a")
		}
	}
	fmt.Println(len(s), utf8.RuneCountInString(s))
	if ! PassVer {		
		return "Ok"
	}else {
		return "Wrong password"
	}
}