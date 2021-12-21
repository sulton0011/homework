package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Saralash("is2 Thi1s T4est 3a"))
}

func Saralash(s string) (lst string) {
	sList := strings.Split(s, " ")
	count := 1
	for count < len(sList) + 1{
		for i := range sList {
			for j := range sList[i] {
				a, _ := strconv.Atoi(string(sList[i][j]))
				if a == count {
					lst = lst + " " + string(sList[i])
					lst = strings.ReplaceAll(lst, string(sList[i][j]), "")
					count ++
				}
			}
		}
	}

	return
}