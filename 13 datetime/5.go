/*
	У вас есть переменные month, year, которые содержат входные пользовательские данные.
	Напишите код, подсчитывает количество дней в указанном месяце month года year и записывает результат в переменную result.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func diff(month, year string) int {

	switch month {
	case "1":
		return 31
	case "2":
		yearInt, _ := strconv.Atoi(year)
		if (yearInt%4 == 0 && yearInt%100 != 0) || yearInt%400 == 0 {
			return 29
		}
		return 28
	case "3":
		return 31
	case "4":
		return 30
	case "5":
		return 31
	case "6":
		return 30
	case "7":
		return 31
	case "8":
		return 31
	case "9":
		return 30
	case "10":
		return 31
	case "11":
		return 30
	case "12":
		return 31
	}

	return 0
}

func main() {
	month, year := ReadInput()
	result := diff(month, year)
	fmt.Println(result)
}

func ReadInput() (string, string) {
	var month, year string
	var values []string
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	month = values[0]
	year = values[1]
	return month, year
}
