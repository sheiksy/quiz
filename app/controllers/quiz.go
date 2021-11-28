package controllers

import (
	"fmt"
	"quiz/app/models"
	"time"
)

func Quiz() {
	num1 := models.Numerics()
	num2 := models.Numerics()
	fmt.Printf("%d + %d = ", num1, num2)
	time.Sleep(time.Second * 3)
	fmt.Println(num1 + num2)
}
