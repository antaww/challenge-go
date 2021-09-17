package main

import "fmt"

func NightClub(age int, genre string) {
	if genre == "fille" && age >= 18 {
		fmt.Println("gratuit")
	} else if genre != "fille" && genre != "homme" {
		fmt.Println("vous ne pouvez pas rentrer")
	} else if age >= 18 {
		fmt.Println("payer 20€")
	}
}

func main() {
	NightClub(18, "homme")
	NightClub(18, "fille")
	NightClub(20, "chien")
	Muliplication(4, 2)
}

func Muliplication(nb1 int, nb2 int) {
	fmt.Println(nb1 * nb2)
}
