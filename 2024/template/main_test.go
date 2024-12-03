package main

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)



func Test_stepOne(t *testing.T){
  logrus.SetLevel(logrus.InfoLevel)
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
