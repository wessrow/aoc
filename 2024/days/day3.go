package days

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func loadData(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func extractEntries(data string) ([][]string, error) {
	r, err := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|(do\(\))|(don't\(\))`)
	if err != nil {
		return nil, err
	}
	return r.FindAllStringSubmatch(data, -1), nil
}

func loopList(multList [][]string) {
	var total int64
	var add bool = true
	for _, v := range multList {
		if v[0] == "don't()" {
			add = false
			continue
		}
		if v[0] == "do()" {
			add = true
			continue
		}
		x, _ := strconv.ParseInt(v[1], 10, 64)
		y, _ := strconv.ParseInt(v[2], 10, 64)
		xy := x * y
		if add {
			total = total + xy
		}

	}
	log.Println(total)
}

func Day3() {
	data, _ := loadData("misc/day3-corrupted.txt")
	entries, _ := extractEntries(data)

	loopList(entries)

}
