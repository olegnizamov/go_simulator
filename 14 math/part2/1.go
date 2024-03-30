/*
	У вас есть переменная n, которая содержит входные пользовательские данные.
	Напишите код, который находит сумму первой и последней цифры числа n и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumOfFirstAndLastDigits(n int) string {
	numStr := strconv.Itoa(n)
	firstDigit, _ := strconv.Atoi(string(numStr[0]))
	lastDigit, _ := strconv.Atoi(string(numStr[len(numStr)-1]))
	return fmt.Sprintf("Сумма первой и последней цифры числа %s равна %s", strconv.Itoa(n), strconv.Itoa(firstDigit+lastDigit))
}

func main() {
	n := ReadInput()
	result := sumOfFirstAndLastDigits(n)
	fmt.Println(result)
}

func ReadInput() int {
	var n int
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	n, _ = strconv.Atoi(input)
	return n
}
