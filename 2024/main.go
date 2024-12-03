package main

import (
	"aoc2024/days"
	"os"
)

func main() {
	n := os.Args[1]

	switch n {
	case "1":
		days.Day1()
	case "2":
		days.Day2()
	case "3":
		days.Day3()
	}
}
