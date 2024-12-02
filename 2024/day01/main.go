package main

import (
	"math"
	"slices"
)


func StepOne(list1 []int, list2 []int) (int, error)  {


  slices.Sort(list1)
  slices.Sort(list2)

  var total float64;
  for i, v1 := range list1{
    
    v2 := list2[i]
    diff := float64(v1 - v2)
    total += math.Abs(diff)
  }
  return int(total) , nil
}


func StepTwo(list1 []int, list2 []int) (int, error)  {

  slices.Sort(list1)
  slices.Sort(list2)


  var total int
  for _, v1 := range list1{
    
    multiplier := 0
    for _, v2 := range list2{
      if v2 == v1{
        multiplier += 1
      }
    }

    total += v1 * multiplier
  }

  return total, nil

}

