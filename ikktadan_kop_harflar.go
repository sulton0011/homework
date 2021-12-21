package main

import (
	"fmt"
	"strings"
)
func main () {
	fmt.Println(CheckTheCharacters("assbgbd"))
}

func CheckTheCharacters(s string) (n string){
	for i := range s {
		if strings.Count(s, string(s[i])) == 1{
			n += string(s[i])
		}
	}

	return n
}