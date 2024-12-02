package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string

func parseData() ([]int, []int, error){

  var list1 []int
  var list2 []int
  lines := strings.Split(data, "\n")
  for _, line := range lines{
    vals := strings.Split(line, " ")
    if vals[0] == ""{
      continue
    }
    if val1, err := strconv.Atoi(vals[0]); err != nil{
      return list1, list2, err
    }else{
      list1 = append(list1, val1)
    }
    if val2, err := strconv.Atoi(vals[len(vals)-1]); err != nil{
      return list1, list2, err
    }else{
      list2 = append(list2, val2)
    }
  }
  return list1, list2,nil
}

