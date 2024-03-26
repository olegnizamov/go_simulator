/*
	У вас есть переменная n, которая содержит входные пользовательские данные.
	Напишите код, который находит сумму цифр числа n и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sumOfDigits(n int) string {
	str := strconv.Itoa(n)
	var sum int

	for i := 0; i < len(str); i++ {
		current, _ := strconv.Atoi(string(str[i]))
		sum = sum + current
	}

	return fmt.Sprintf("Сумма цифр числа %d равна %d", n, sum)
}

func main() {
	n := ReadInput()
	result := sumOfDigits(n)
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
