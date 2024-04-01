/*
	У вас есть переменные page, message, которые принимают входные пользовательские данные.
	Напишите код, который увеличивает значение переменной message  на значение page и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func incrementString(page int, message string) string {
	nums := strings.Split(message, "-")
	num, _ := strconv.Atoi(nums[1])
	result := page + num
	return fmt.Sprintf("%s-%d", nums[0], result)
}

func main() {
	page, message := ReadInput()
	result := incrementString(page, message)
	fmt.Println(result)
}

func ReadInput() (int, string) {
	var page int
	var message string
	var values []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	page, _ = strconv.Atoi(values[0])
	message = values[1]
	return page, message
}
