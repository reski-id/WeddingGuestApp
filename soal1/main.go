package main

import "fmt"

func CetakTabelPerkalian(number int) string {
	var str string
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			str += fmt.Sprint(j*i, "\t")
		}
		str += fmt.Sprintln()
	}
	return str
}

func main() {
	fmt.Println(CetakTabelPerkalian(5))
}
