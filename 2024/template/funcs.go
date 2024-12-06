package main

import (
	_ "embed"
	"strings"
)

//go:embed data.txt
var data string

func parseData() ([][]int, error){

  var values [][]int
  lines := strings.Split(data, "\n")
  for _, line := range lines{
    if line == ""{
      continue
    }
  }
  return values,nil
}

