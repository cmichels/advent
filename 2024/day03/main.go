package main

import (
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
)

const mulPattern = `mul\((\d{1,3}),(\d{1,3})\)`
const doPattern = `(mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\))`

func stepTwo(data []string) int {
	var total int

	for _, line := range data {
    logrus.Debugf("line: [%s]", line)
		matches := matchMultipliers(line, doPattern)

    logrus.Tracef("matches: [%s]", matches)
    filtered := filterMultipliers(matches)
    logrus.Tracef("filtered: [%s]", filtered)
		for _, val := range filtered {
			val := parseMultipliers(val)
			total += val
		}
	}
	return total
}

func stepOne(data []string) int {

	var total int

	for _, line := range data {
		matches := matchMultipliers(line, mulPattern)
		for _, val := range matches {
			val := parseMultipliers(val)
			total += val
		}
	}
	return total
}

func matchMultipliers(line string, expression string) []string {

	// Compile the regex
	re := regexp.MustCompile(expression)

	// Find all matches
	return re.FindAllString(line, -1)
}

func filterMultipliers(multipliers []string) []string{
  enabled := true
  var result []string

  for _, m := range multipliers{

    if m == "don't()"{
      enabled = false
      logrus.Debug("disabled")
      continue
    }
    if m == "do()"{
      logrus.Debug("enabled")
      enabled = true
      continue
    }


    if enabled{
      logrus.Debugf("m:[%s], calc: [%d]", m, parseMultipliers(m))
      result = append(result, m)
    }
  }
  return result
}

func parseMultipliers(val string) int {
	re := regexp.MustCompile(mulPattern)

	matches := re.FindAllStringSubmatch(val, -1)

	for _, m := range matches {
		logrus.Tracef("match: [%s]", m)
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		return a * b
	}

	return 0

}

