package main

import (
	"bufio"
	"fmt"
	"greeting"
	"os"
	"strconv"
	"strings"
)

func getFloat64(terminalText string) (float64, error) {
	fmt.Print(terminalText)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0.0, err
	}

	input = strings.TrimSpace(input)
	returnFloat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0.0, err
	}

	return returnFloat, nil
}

func main() {
	myFloat, err := getFloat64("Введите число: ")

	if err != nil {
		print(err)
	} else {
		fmt.Printf("Ваше число: %0.2f\n", myFloat)
	}

	greeting.Hello()
}
