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
	isSuccessCount := 0
	var rows = ReadRows("./level_2/task_1/input.txt")
	for _, row := range rows {
		if IsSuccess(&row) {
			isSuccessCount++
		}
	}

	fmt.Println(isSuccessCount)
}

func IsSuccess(row *Row) bool {
	if row.level == Wrong {
		return false
	}

	for i := 0; i < len(row.items)-1; i++ {
		if (row.items)[i] == (row.items)[i+1] {
			return false
		}

		diff := Diff(row.items[i], row.items[i+1])
		if diff > 3 {
			return false
		}
	}

	return true
}

func Diff(a, b int) int {
	if a < b {
		return Abs(b - a)
	}

	return Abs(a - b)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ReadRows(path string) (rows []Row) {
	file, err := os.Open(path)
	rows = []Row{}

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
		items := strings.Split(row, " ")
		var columns []int

		for _, rowItems := range items {
			val, _ := strconv.Atoi(rowItems)
			columns = append(columns, val)
		}

		rows = append(rows, Row{
			items: columns,
			level: CheckType(&columns),
		})
	}

	return
}

func CheckType(items *[]int) int {
	increasing := false
	decreasing := false

	for i := 1; i < len(*items); i++ {
		if (*items)[i] > (*items)[i-1] {
			increasing = true
		}

		if (*items)[i] < (*items)[i-1] {
			decreasing = true
		}

		if increasing && decreasing {
			return Wrong
		}
	}

	return Normal
}

type Row struct {
	items []int
	level int
}

const (
	Normal = iota
	Wrong  = iota
)
