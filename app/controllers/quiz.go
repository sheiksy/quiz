package controllers

import (
	"fmt"
	"quiz/app/models"
)

var answer int

func Quiz() {
	num1 := models.Numerics()
	num2 := models.Numerics()
	fmt.Printf("%d + %d = ", num1, num2)
	fmt.Scan(&answer)
	if answer == (num1 + num2) {
		fmt.Println("Correct!")
	} else {
		fmt.Printf("The correct answer is %d \n", num1+num2)
	}
}
