package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	columnFirst, columnSecond := ReadColumns("./level_1/task_2/input.txt")

	mapItems := PrepareDictionary(&columnSecond)

	result := CountResult(&columnFirst, &mapItems)

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

func PrepareDictionary(items *[]int) map[int]int {
	result := map[int]int{}

	for _, item := range *items {
		result[item]++
	}

	return result
}

func CountResult(columnFirst *[]int, mapItems *map[int]int) int {
	result := 0

	for _, num := range *columnFirst {
		count := (*mapItems)[num]
		result += num * count
	}

	return result
}
