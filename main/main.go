package main

import (
	"fmt"
	"project2/nation"
)

func main() {
	fmt.Println("a person speaks ...")

	var person nation.Nation

	person.Name = "EnglishMan"
	person.Speaks("English")

}
