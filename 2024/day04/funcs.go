package main

import (
	_ "embed"
	"strings"

	"github.com/sirupsen/logrus"
)

//go:embed data.txt
var data string

func parseData() ([][]string, error){
  var values [][]string
  lines := strings.Split(data, "\n")
  for _, line := range lines{
    if line == ""{
      continue
    }

    chars := strings.Split(line, "")

    values = append(values, chars)
    logrus.Tracef("values %+v", values)
  }


  return values,nil
}

