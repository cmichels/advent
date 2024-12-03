package main

import (
	_ "embed"
	"strings"
)

//go:embed data.txt
var data string

func parseData() ([]string, error){


  var values []string
  lines := strings.Split(data, "\n")
  for _, line := range lines{
    if line == ""{
      continue
    }
    values = append(values, line)
  }
  return values,nil
}

