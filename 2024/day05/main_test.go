package main

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var pageOrder = [][]int{
  {47,53},
  {97,13},
  {97,61},
  {97,47},
  {75,29},
  {61,13},
  {75,53},
  {29,13},
  {97,29},
  {53,29},
  {61,53},
  {97,53},
  {61,29},
  {47,13},
  {75,47},
  {97,75},
  {47,61},
  {75,61},
  {47,29},
  {75,13},
  {53,13},

}

var updates = [][]int{
  {75,47,61,53,29},
  {97,61,53,29,13},
  {75,29,13},
  {75,97,47,61,53},
  {61,13,29},
  {97,13,75,29,47},
}


func Test_stepOne(t *testing.T){
  logrus.SetLevel(logrus.InfoLevel)

  answer := stepOne(pageOrder, updates)
  assert.Equal(t, 143, answer )
}


func Test_runStepOne(t *testing.T)  {
  logrus.SetLevel(logrus.InfoLevel)
  pageOrder, updates := parseData()

  answer := stepOne(pageOrder, updates)
  assert.Equal(t, 4578, answer)
  logrus.Infof("step1: [%d]", answer)
}

func Test_stepTwo(t *testing.T){
  logrus.SetLevel(logrus.InfoLevel)

  answer := stepTwo(pageOrder, updates)
  assert.Equal(t, 123, answer )
}

func Test_runStepTwo(t *testing.T)  {
  logrus.SetLevel(logrus.InfoLevel)
  pageOrder, updates := parseData()

  answer := stepTwo(pageOrder,updates)
  assert.Equal(t, 6179, answer)
  logrus.Infof("step2: [%d]", answer)
}
