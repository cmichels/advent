package main

import (
	"slices"

	"github.com/sirupsen/logrus"
)

const (
	less    = 1
	greater = 2
	equal   = 3
	invalid = 4
)

func stepTwo(data [][]int) int {

	var safeRows int

	logrus.Infof("processing [%d] rows", len(data))
	for _, r := range data {

		logrus.Tracef("row: [%v]", r)
		if len(r) == 0 {
			continue
		}

		if isSafe(r) {
			safeRows++
			continue
		} else {
			for i := 0; i < len(r); i++ {
				first := r[:i]
				last := r[i+1:]
				var cleansed []int
				cleansed = append(cleansed, first...)
				cleansed = append(cleansed, last...)
				logrus.Tracef("cleansed: %v", cleansed)
				if isSafe(cleansed) {
					safeRows++
					break
				}
			}
		}
	}

	return safeRows
}

func isSafe(r []int) bool {
	curVal := r[0]
	row := r[1:]
	var direction []int

	for _, v := range row {
		if v == curVal {
			direction = append(direction, equal)
		} else if abs(v, curVal) > 3 {
			direction = append(direction, invalid)
		} else if v > curVal {
			direction = append(direction, greater)
		} else {
			direction = append(direction, less)
		}
		curVal = v
	}

	if slices.Contains(direction, equal) || slices.Contains(direction, invalid) {
		return false
	}
	maxVal := slices.Max(direction)
	minVla := slices.Min(direction)

	if maxVal == minVla {
		return true
	}
	return false

}

func stepOne(data [][]int) int {

	var safeRows int

	logrus.Tracef("processing [%d] rows", len(data))
	for _, r := range data {

		if len(r) == 0 {
			continue
		}
		if isSafe(r) {

			safeRows++
		}
	}

	return safeRows
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
