package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	A, e := ReadInput()
	var result bool = false
	for i := range A {
		if A[i] == e {
			result = true
		}
	}

	fmt.Println(result)
}

func ReadInput() ([]int, int) {
	var A []int
	var e int
	var input string
	var values []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	values = strings.Split(input, " | ")
	json.Unmarshal([]byte(values[0]), &A)
	e, _ = strconv.Atoi(values[1])
	return A, e
}
