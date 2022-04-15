package main

import "fmt"

func computeMarsYears(earthYears int) int {
	earthDays := earthYears * 365
	marsYears := earthDays / 687
	return marsYears
}

func main() {
	fmt.Println("====EARTH TO MARS AGE CALCULATOR====")
	fmt.Println("How old are you?:")
	var myAge int
	fmt.Scan(&myAge)

	myMartianAge := computeMarsYears(myAge)
	fmt.Println("You are", myMartianAge, "years old on mars!")
}
