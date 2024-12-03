package main

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)



func Test_stepOne(t *testing.T){
  logrus.SetLevel(logrus.InfoLevel)

  vals := [][]int{
    {7,6,4,2,1},
    {1,2,7,8,9},
    {9,7,6,2,1},
    {1,3,2,4,5},
    {8,6,4,4,1},
    {1,3,6,7,9},
  }


  total := stepOne(vals)
  assert.Equal(t, 2, total)
}


func Test_runStepOne(t *testing.T)  {
  logrus.SetLevel(logrus.InfoLevel)
  vals, err := parseData()

  if err != nil {
    logrus.Errorf("error parsing data: [%s]", err)
    return
  }

  answer := stepOne(vals)
  assert.Equal(t, 631, answer)
  logrus.Infof("step1: [%d]", answer)
}
func Test_stepTwo(t *testing.T){
  logrus.SetLevel(logrus.TraceLevel)

  vals := [][]int{
    {7,6,4,2,1},
    {1,2,7,8,9},
    {9,7,6,2,1},
    {1,3,2,4,5},
    {8,6,4,4,1},
    {1,3,6,7,9},
  }


  total := stepTwo(vals)
  assert.Equal(t, 4, total)
}
func Test_runStepTwo(t *testing.T)  {
  logrus.SetLevel(logrus.InfoLevel)
  vals, err := parseData()

  if err != nil {
    logrus.Errorf("error parsing data: [%s]", err)
    return
  }

  answer := stepTwo(vals)
  assert.Equal(t, 665, answer)
  logrus.Infof("step2: [%d]", answer)
}
