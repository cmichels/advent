package main

import (
	"slices"
	"sort"

	"github.com/sirupsen/logrus"
)

func stepTwo(pageOrder, updates [][]int) int {

	reordered := reversePages(pageOrder)
	logrus.Debugf("reverse %+v", reordered)
	total := 0

	var badGuys [][]int

	for _, u := range updates {
		isValid := true
		for i, p := range u {
			logrus.Tracef("find %d", p)
			order, ok := reordered[p]
			logrus.Tracef(" order: %+v", order)

			if !ok {
				continue
			}

			var remaining []int

			if i >= len(u)-1 {
				remaining = []int{p}
			} else {
				remaining = u[i+1:]
			}

			logrus.Tracef("  remaining: %+v", remaining)
			rs := make(map[int]int)
			for i, r := range remaining {
				rs[r] = i
			}
			logrus.Tracef("   rs: %+v", rs)

			for _, o := range order {
				if _, ok := rs[o]; ok {
					isValid = false
				}
			}
		}
		logrus.Tracef(" isValid: %t", isValid)
		if !isValid {
			badGuys = append(badGuys, u)
		}
	}

	var goodGuys [][]int
	for i, bg := range badGuys {
	// for i, bg := range badGuys[len(badGuys)-1:] {

		tempList := slices.Clone(bg)
		logrus.Debugf("%d=%+v", i, bg)
		for _, v := range bg {
			val, ok := reordered[v]
			if !ok {
				continue // no rules apply
			}

			logrus.Debugf("v: %d -- val: %+v", v, val)
			for _, order := range val {
				curIndex := slices.Index(tempList, order)
				if curIndex == -1 {
					continue
				}
				targetIdx := slices.Index(tempList, v)
				logrus.Debugf(" order: %d, %d, %d", order, curIndex, targetIdx)
				if curIndex > targetIdx {
					tempList = slices.Delete(tempList, targetIdx, targetIdx+1)
					logrus.Debugf("   templist %+v", tempList)
					tempList = slices.Insert(tempList, curIndex, v)
					logrus.Debugf("   templist %+v", tempList)
				}
			}
		}

		goodGuys = append(goodGuys, tempList)
	}

	logrus.Debugf("goodguys%+v", goodGuys)
  for _, gg := range goodGuys{
    middle := len(gg)/2
    total += gg[middle]
  }
	return total
}

func toMap(data []int) map[int]int {
	rs := make(map[int]int)
	for i, r := range data {
		rs[r] = i
	}

	return rs
}

func toSlice(data map[int]int) []int {

	var result []int
	kv := make([]struct {
		k int
		v int
	}, 0, len(data))
	for k, v := range data {
		kv = append(kv, struct {
			k int
			v int
		}{k, v})
	}

	sort.Slice(kv, func(i, j int) bool {
		return kv[i].v < kv[j].v
	})

	for _, v := range kv {
		result = append(result, v.k)
	}

	return result

}

func stepOne(pageOrder, updates [][]int) int {

	reordered := reversePages(pageOrder)
	logrus.Debugf("reverse %+v", reordered)
	total := 0

	for _, u := range updates {
		isValid := true
		for i, p := range u {
			logrus.Debugf("find %d", p)
			order, ok := reordered[p]
			logrus.Debugf(" order: %+v", order)

			if !ok {
				continue
			}

			var remaining []int

			if i >= len(u)-1 {
				remaining = []int{p}
			} else {
				remaining = u[i+1:]
			}

			logrus.Debugf("  remaining: %+v", remaining)
			rs := make(map[int]int)
			for _, r := range remaining {
				rs[r] = 0
			}
			logrus.Debugf("   rs: %+v", rs)

			for _, o := range order {
				if _, ok := rs[o]; ok {
					isValid = false
				}
			}
		}
		logrus.Debugf(" isValid: %t", isValid)
		if isValid {
			middle := len(u) / 2
			total += u[middle]
		}
	}
	return total
}

func reversePages(pageOrder [][]int) map[int][]int {

	reverseOrder := map[int][]int{}
	for _, po := range pageOrder {
		page1 := po[0]
		page2 := po[1]

		if _, ok := reverseOrder[page2]; !ok {
			reverseOrder[page2] = []int{}
		}
		reverseOrder[page2] = append(reverseOrder[page2], page1)
	}

	return reverseOrder
}
