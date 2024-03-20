/*
	У вас есть переменные x1, x2, x3, которые содержат входные пользовательские данные.
	Напишите код, который находит максимальное и минимальное число из x1, x2, x3 и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x1, x2, x3 := ReadInput()
	var result string
	var maxElement int = max(x1, x2, x3)
	var minElement int = min(x1, x2, x3)
	result = "минимальное: " + strconv.Itoa(minElement) + ", максимальное: " + strconv.Itoa(maxElement)
	fmt.Println(result)
}

func ReadInput() (int, int, int) {
	var x1, x2, x3 int
	var input string
	var values []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	values = strings.Split(input, " ")
	x1, _ = strconv.Atoi(values[0])
	x2, _ = strconv.Atoi(values[1])
	x3, _ = strconv.Atoi(values[2])

	return x1, x2, x3
}
