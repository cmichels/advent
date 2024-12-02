package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func Test_StepOne(t *testing.T){
  list1 := []int{3,4,2,1,3,3}
  list2 := []int{4,3,5,3,9,3}

  total, err := StepOne(list1,list2)
  assert.Nil(t, err)
  assert.Equal(t, 11, total)
}

func Test_StepTwo(t *testing.T){
  list1 := []int{3,4,2,1,3,3}
  list2 := []int{4,3,5,3,9,3}

  total, err := StepTwo(list1,list2)
  assert.Nil(t, err)
  assert.Equal(t, 31, total)
}

func Test_runStepOne(t *testing.T){
  list1, list2, err := parseData()

  if err != nil {
    println("error1: ", err.Error())
    return
  }


  total, err := StepOne(list1, list2)

  if err != nil {
    println("error2: ", err.Error())
    return
  }


  println("1 totals: ", total)

}
func Test_runStepTwo(t *testing.T){
  list1, list2, err := parseData()

  if err != nil {
    println("2: error1: ", err.Error())
    return
  }


  total, err := StepTwo(list1, list2)

  if err != nil {
    println("2: error2: ", err.Error())
    return
  }


  println("2 totals: ", total)

}
