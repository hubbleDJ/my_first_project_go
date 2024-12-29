package main

import (
	"fmt"
	"my_first_project_go/keyboard"
)

func main() {
	myFloat, err := keyboard.GetFloat64("Введите число: ")

	if err != nil {
		print(err)
	} else {
		fmt.Printf("Ваше число: %0.2f\n", myFloat)
	}
}
