package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	columnFirst, columnSecond := ReadColumns("./level_1/task_1/input.txt")
	result := CountResult(&columnFirst, &columnSecond)

	fmt.Println(result)
}

func ReadColumns(path string) (columnFirst []int, columnSecond []int) {
	file, err := os.Open(path)
	columnFirst = []int{}
	columnSecond = []int{}

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	content, err := io.ReadAll(file)
	for _, row := range strings.Split(string(content), "\n") {
		items := strings.Split(row, ",")

		val, _ := strconv.Atoi(items[0])
		columnFirst = append(columnFirst, val)

		val, _ = strconv.Atoi(items[1])
		columnSecond = append(columnSecond, val)
	}

	return
}

func CountResult(columnFirst *[]int, columnSecond *[]int) int {
	sort.Ints(*columnFirst)
	sort.Ints(*columnSecond)

	result := 0

	for index, num := range *columnFirst {
		result += Abs(num - (*columnSecond)[index])
	}

	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
