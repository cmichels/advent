package main

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)



func Test_stepOne(t *testing.T){
  logrus.SetLevel(logrus.InfoLevel)
  data := []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"} 

  answer := stepOne(data)
  assert.Equal(t, 161, answer )
}


func Test_match(t *testing.T) {

  logrus.SetLevel(logrus.InfoLevel)
  data := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"


  matches := matchMultipliers(data, mulPattern)
  logrus.Tracef("matches: [%s]", matches)


  assert.Equal(t, 4, len(matches))
}

func Test_matchDoDont(t *testing.T) {

  logrus.SetLevel(logrus.InfoLevel)
  data := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"


  matches := matchMultipliers(data, doPattern)
  logrus.Tracef("matches: [%s]", matches)


  assert.Equal(t, 6, len(matches))
}

func Test_filterMultipliers(t *testing.T) {

  logrus.SetLevel(logrus.InfoLevel)
  data := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
  matches := matchMultipliers(data, doPattern)
  logrus.Tracef("matches: [%s]", matches)
  assert.Equal(t, 6, len(matches))

  expected := filterMultipliers(matches)
  logrus.Tracef("filtered: [%s]", expected)
  assert.Equal(t, 2, len(expected))

}
func Test_parseMultiplier(t *testing.T){

  logrus.SetLevel(logrus.InfoLevel)

  data := "mul(2,4)"


  expected := parseMultipliers(data)


  assert.Equal(t, 8, expected)
}


func Test_runStepOne(t *testing.T)  {
  logrus.SetLevel(logrus.InfoLevel)
  vals, err := parseData()

  if err != nil {
    logrus.Errorf("error parsing data: [%s]", err)
    return
  }

  answer := stepOne(vals)
  assert.Equal(t, 161289189, answer)
  logrus.Infof("step1: [%d]", answer)
}

func Test_stepTwo(t *testing.T){
  logrus.SetLevel(logrus.InfoLevel)
  data := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}

  answer := stepTwo(data)
  assert.Equal(t, 48, answer )
}

func Test_runStepTwo(t *testing.T)  {
  logrus.SetLevel(logrus.DebugLevel)
  vals, err := parseData()

  if err != nil {
    logrus.Errorf("error parsing data: [%s]", err)
    return
  }

  answer := stepTwo(vals)
  assert.Equal(t, 161289189, answer)
  logrus.Infof("step2: [%d]", answer)
}
