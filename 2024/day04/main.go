package main

import "github.com/sirupsen/logrus"

var xmas = []string{"X", "M", "A", "S"}

const (
	xIndex = 0
	mIndex = 1
	aIndex = 2
	sIndex = 3
)

func stepTwo(data [][]string) int {
	return 0
}

func stepOne(data [][]string) int {
	total := 0
	xVal := xmas[xIndex]
	for y, line := range data {
		for x, val := range line {
			if val == xVal {
				logrus.Tracef("(%d,%d)=%s", x, y, val)
				if matches := isXmas(data, x, y); matches > 0 {
					logrus.Debugf("merry xmas - (%d,%d)=%s", x, y, val)
					total += matches
				}
			}
		}
	}

	return total
}

func isXmas(data [][]string, x, y int) int {
	total := 0
	for searchX := -1; searchX <= 1; searchX++ {
		for searchY := -1; searchY <= 1; searchY++ {
			if check(data, mIndex, x, y, searchX, searchY) {
				total++
			}
		}
	}

	return total

}

func check(data [][]string, index, x, y, xMove, yMove int) bool {

	x = x + xMove
	y = y + yMove

	logrus.Tracef("nextVal (%d,%d)", x, y)
	if x < 0 || y < 0 {
		return false
	}

	if len(data) > y && len(data[y]) > x {
		logrus.Tracef("nextVal %s @ (%d,%d)", data[y][x], x, y)
		if data[y][x] == xmas[index] {

			if index == sIndex {
				return true
			}
			index++
			return check(data, index, x, y, xMove, yMove)
		}

	}

	return false

}
