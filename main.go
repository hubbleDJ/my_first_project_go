package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func getValueInPointer(myPointer *bool) bool {
	return *myPointer
}

func main() {
	// Ввод числа от пользователя
	fmt.Print("Введи число: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	inputValueForm := strings.TrimSpace(input)
	inputInt, _ := strconv.Atoi(inputValueForm)
	fmt.Println(inputInt)

	// Генерация случайного числа
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	randNum := rand.Intn(101)
	fmt.Println(randNum)

	myBoolVar := true
	fmt.Println(getValueInPointer(&myBoolVar))
}
