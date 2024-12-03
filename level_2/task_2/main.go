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
	var rows = ReadRows("./level_2/task_2/input.txt")
	for _, row := range rows {
		checkResult := IsSuccess(&row)
		if !checkResult.result {
			newRow := append(row.items[:checkResult.position], row.items[checkResult.position+1:]...)
			checkResult = IsSuccess(&Row{items: newRow, level: CheckType(&newRow)})
			if checkResult.result {
				isSuccessCount++
				continue
			}
		} else {
			isSuccessCount++
			continue
		}
	}

	fmt.Println(isSuccessCount)
}

func IsSuccess(row *Row) CheckResult {
	if row.level.result == Wrong {
		return CheckResult{
			result:   false,
			position: row.level.position,
		}
	}

	for i := 0; i < len(row.items)-1; i++ {
		if (row.items)[i] == (row.items)[i+1] {
			return CheckResult{
				result:   false,
				position: i,
			}
		}

		diff := Diff(row.items[i], row.items[i+1])
		if diff > 3 {
			return CheckResult{
				result:   false,
				position: i,
			}
		}
	}

	return CheckResult{
		result:   true,
		position: 0,
	}
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

func CheckType(items *[]int) CheckTypeResult {
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
			return CheckTypeResult{
				result:   Wrong,
				position: i,
			}
		}
	}

	return CheckTypeResult{
		result:   Normal,
		position: 0,
	}
}

type Row struct {
	items []int
	level CheckTypeResult
}

type CheckResult struct {
	result   bool
	position int
}

type CheckTypeResult struct {
	result   int
	position int
}

const (
	Normal = iota
	Wrong  = iota
)
