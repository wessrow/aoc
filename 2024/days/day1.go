package days

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func loadList(listName string) ([]float64, error) {
	file, err := os.Open(listName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var returnList []float64
	for scanner.Scan() {
		n, _ := strconv.ParseFloat(scanner.Text(), 64)
		returnList = append(returnList, n)
	}

	return returnList, nil
}

func countInList(searchIn []float64, searchFor float64) int {
	counter := 0
	for _, v := range searchIn {
		if v == searchFor {
			counter = counter + 1
		}
	}
	return counter
}

func Day1() int {
	list1, err := loadList("misc/list1.txt")
	if err != nil {
		log.Fatal(err)
	}
	list2, err := loadList("misc/list2.txt")
	if err != nil {
		log.Fatal(err)
	}

	sort.Float64s(list1)
	sort.Float64s(list2)

	var totalDistance int

	for _, v := range list1 {
		// distance := math.Max(v, list2[i]) - math.Min(v, list2[i])
		// totalDistance = totalDistance + distance
		counter := countInList(list2, v)
		simScore := counter * int(v)

		totalDistance = totalDistance + simScore

	}

	return totalDistance
}
