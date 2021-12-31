package main

import (
	"quiz/app/controllers"
)

func main() {
	for i := 0; i < 10; i++ {
		controllers.Quiz()
	}
}
