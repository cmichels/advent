package main

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)


func Test_stepOne(t *testing.T){
  logrus.SetLevel(logrus.InfoLevel)
  // logrus.SetLevel(logrus.TraceLevel)
  data := [][]string{
    //0   1   2   3   4   5   6   7   8   9
    {"M","M","M","S","X","X","M","A","S","M"},// 0
    {"M","S","A","M","X","M","S","M","S","A"},// 1
    {"A","M","X","S","X","M","A","A","M","M"},// 2
    {"M","S","A","M","A","S","M","S","M","X"},// 3
    {"X","M","A","S","A","M","X","A","M","M"},// 4
    {"X","X","A","M","M","X","X","A","M","A"},// 5
    {"S","M","S","M","S","A","S","X","S","S"},// 6
    {"S","A","X","A","M","A","S","A","A","A"},// 7
    {"M","A","M","M","M","X","M","M","M","M"},// 8
    {"M","X","M","X","A","X","M","A","S","X"},// 9
  }


  answer := stepOne(data)

  assert.Equal(t, 18, answer)
}


func Test_runStepOne(t *testing.T)  {
  logrus.SetLevel(logrus.InfoLevel)
  vals, err := parseData()

  if err != nil {
    logrus.Errorf("error parsing data: [%s]", err)
    return
  }

  answer := stepOne(vals)
  assert.Equal(t, 2571, answer)
  logrus.Infof("step1: [%d]", answer)
}

func Test_stepTwo(t *testing.T){
  logrus.SetLevel(logrus.DebugLevel)
  // logrus.SetLevel(logrus.TraceLevel)
  data := [][]string{
    //0   1   2   3   4   5   6   7   8   9
    {"M","M","M","S","X","X","M","A","S","M"},// 0
    {"M","S","A","M","X","M","S","M","S","A"},// 1
    {"A","M","X","S","X","M","A","A","M","M"},// 2
    {"M","S","A","M","A","S","M","S","M","X"},// 3
    {"X","M","A","S","A","M","X","A","M","M"},// 4
    {"X","X","A","M","M","X","X","A","M","A"},// 5
    {"S","M","S","M","S","A","S","X","S","S"},// 6
    {"S","A","X","A","M","A","S","A","A","A"},// 7
    {"M","A","M","M","M","X","M","M","M","M"},// 8
    {"M","X","M","X","A","X","M","A","S","X"},// 9
  }


  answer := stepTwo(data)

  assert.Equal(t, 9, answer)
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
