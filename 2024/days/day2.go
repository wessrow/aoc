package days

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadReports(listName string) ([][]int, error) {
	file, err := os.Open(listName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var returnList [][]int
	for scanner.Scan() {
		lineList := strings.Split(scanner.Text(), " ")
		var lineIntList []int
		for _, entry := range lineList {
			intEntry, _ := strconv.ParseInt(entry, 10, 64)
			lineIntList = append(lineIntList, int(intEntry))
		}
		returnList = append(returnList, lineIntList)
	}

	return returnList, nil
}

func testDescending(listLine []int, dampening int) bool {
	for i, v := range listLine {
		if i == 0 {
			continue
		}
		if v >= listLine[i-1] {
			if dampening == 1 {
				return false
			}
			listLine = append(listLine[:i], listLine[i+1:]...)
			return testDescending(listLine, dampening+1)
		}
		if listLine[i-1]-v > 3 {
			if dampening == 1 {
				return false
			}
			listLine = append(listLine[:i], listLine[i+1:]...)
			return testDescending(listLine, dampening+1)
		}
	}
	log.Println(dampening, listLine)
	return true
}

func testAscending(listLine []int, dampening int) bool {
	for i, v := range listLine {
		if i == 0 {
			continue
		}
		if v <= listLine[i-1] {
			if dampening == 1 {
				return false
			}

			listLine = append(listLine[:i], listLine[i+1:]...)

			test := testAscending(listLine, dampening+1)
			return test
		}

		if v-listLine[i-1] > 3 {
			if dampening == 1 {
				return false
			}
			listLine = append(listLine[:i], listLine[i+1:]...)

			test := testAscending(listLine, dampening+1)
			return test
		}
	}

	return true
}

func safeReport(listLine []int) bool {
	ascendingCopy := append([]int(nil), listLine...)
	descendingCopy := append([]int(nil), listLine...)
	if testAscending(ascendingCopy, 0) || testDescending(descendingCopy, 0) {
		return true
	}
	return false
}

func Day2() {
	input, err := loadReports("misc/day2-reports.txt")
	if err != nil {
		log.Fatal(err)
	}

	counter := 0

	for _, listLine := range input {
		if safeReport(listLine) {
			counter = counter + 1
		}
	}

	log.Println(counter)
}
