/*
	У вас есть переменная minutes, которая принимает входные пользовательские данные.
	Напишите код, который переводит минуты в секунды и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	minutes := ReadInput()
	var result int = minutes * 60
	fmt.Println(result)
}

func ReadInput() int {
	var minutes int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	minutes, _ = strconv.Atoi(scanner.Text())

	return minutes
}
