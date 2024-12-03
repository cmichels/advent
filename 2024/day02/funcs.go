package main

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
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
    stringVals := strings.Split(line, " ")
    vals := toInts(stringVals)
    values = append(values, vals)
  }
  return values,nil
}


func toInts(data []string) []int{
  ints := make([]int, len(data))
  for i, s := range data{
    if v, err := strconv.Atoi(s); err != nil{
      logrus.Errorf("error parsing val [%s]. caused by: [%s]", s, err)
    }else{
      ints[i] =  v
    }
  }
  
  return ints
}

