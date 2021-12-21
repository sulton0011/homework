package main

import (
	"fmt"
)

func main () {
	fmt.Println(Farqi("asdfs asf", "gfva"))
}

func Farqi(s, a string) (lst string, count int) {
	for i := range s{
		if Bormi(a, string(s[i])){
			if ! Bormi(lst, string(s[i])){
				lst = lst + string(s[i])
			}
		}
	}

	return lst, len(lst)
}

func Bormi(soz, harf string) bool{
	for i := range soz {
		if string(soz[i]) == harf {
			return true
		}
	}
	return false
}