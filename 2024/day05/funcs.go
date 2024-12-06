package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed pageOrders.txt
var pageOrderFile string
//go:embed updates.txt
var updateFile string


func parseData() ([][]int, [][]int){


  var pageOrders [][]int
  lines := strings.Split(pageOrderFile, "\n")
  for _, line := range lines{
    if line == ""{
      continue
    }
    
    vals := strings.Split(line, "|")
    a, _ := strconv.Atoi(vals[0])
    b, _ := strconv.Atoi(vals[1])

    pageOrders = append(pageOrders, []int{a,b})
  }


  var updates [][]int
  uLines := strings.Split(updateFile, "\n")
  for _, line := range uLines{
    if line == ""{
      continue
    }
    var pages []int
    vals := strings.Split(line, ",")
    for _, val := range vals{

      a, _ := strconv.Atoi(val)
      pages = append(pages, a)
    }
    updates = append(updates, pages)

  }


  return pageOrders,updates
}

