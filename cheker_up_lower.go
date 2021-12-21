package main

func Cheker(s string) (a, b int) {
	for i := range s {
		if s[i] >= 65 && s[i] <= 90 {
			a ++
		}else {
			b ++
		}
	}
	return
}
