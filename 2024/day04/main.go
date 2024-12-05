package main

import (
	"slices"

	"github.com/sirupsen/logrus"
)

var xmas = []string{"X", "M", "A", "S"}

type Pair struct {
	x int
	y int
}

var bottomLeft = Pair{x: -1, y: -1}
var topRight = Pair{x: 1, y: 1}
var topLeft = Pair{x: 1, y: -1}
var bottomRight = Pair{x: -1, y: 1}

const (
	xIndex = 0
	mIndex = 1
	aIndex = 2
	sIndex = 3
)

func stepTwo(data [][]string) int {

	findVal := xmas[aIndex]
	total := 0

	for y := 1; y < len(data)-1; y++ {
		for x := 1; x < len(data[y])-1; x++ {
			logrus.Tracef("(%d,%d)=%s", x, y, data[y][x])
			if findVal == data[y][x] {
				logrus.Debugf("finding xmas (%d,%d)",x,y)
				if matches := isMasX(data, x, y); matches > 0 {
					logrus.Debugf(" merry xmas - (%d,%d)=%d", x, y,matches)
					total += matches
				}
			}
		}

	}
	return total
}

func isMasX(data [][]string, x, y int) int {

	tl := data[y + topLeft.y][x + topLeft.x]
	tr := data[y + topRight.y][x + topRight.x]
	bl := data[y + bottomLeft.y][x + bottomLeft.x]
	br := data[y + bottomRight.y][x + bottomRight.x]

	letters := []string{tl, tr, bl, br}
  logrus.Debugf("  letters %+v", letters)
	if slices.Contains(letters, "A") || slices.Contains(letters, "X"){
		return 0
	}

	mas := map[string]int{"M": 0, "S": 0}
	for _, l := range letters {
		mas[l]++
	}

  logrus.Debugf("  mas %+v", mas)

  if mas["M"] == 2 && mas["S"] == 2{
    if tl != br && tr != bl{
      return 1
    }
  }

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
