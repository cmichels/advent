package main

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_stepOne(t *testing.T) {
	logrus.SetLevel(logrus.InfoLevel)
	data := []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}

	answer := stepOne(data)
	assert.Equal(t, 161, answer)
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
func Test_parseMultiplier(t *testing.T) {

	logrus.SetLevel(logrus.InfoLevel)

	data := "mul(2,4)"

	expected := parseMultipliers(data)

	assert.Equal(t, 8, expected)
}

func Test_runStepOne(t *testing.T) {
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

func Test_stepTwo(t *testing.T) {
	logrus.SetLevel(logrus.InfoLevel)
	data := []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}

	answer := stepTwo(data)
	expected := 2*4 + 8*5 + 2*4 + 8*5 + 2*4 + 8*5
	logrus.Debugf("expected: [%d]", expected)
	assert.Equal(t, expected, answer)
}
func Test_stepTwo_debug(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	data := []string{
		"}}+{where()mul(873,602) mul(954,447)^where()~mul(548,799)-<what()mul(588,631)^who()'@( [mul(143,388)how(445,327))$ select()who()mul(746,719)mul(963,262)}'*+why()<?&/select()don't()[%]% ^^mul(933,492)don't() ^-who()(%how()]mul(583,700))!,where()mul(999,507)](mul(821,504)^%/;:-mul(471,220)who()&'who(161,37)<'mul(843,213),<mul(298,802)-()@how()where()@+[:mul(910,85)'when()($select()%-*mul(532,213)*from()select()$,what()mul(956,380)mul(326,87)mul(243,775)!/:from()&^^mul(118,409)!)what()select()mul(610,89)mul(432,774)from()# from()<mul(400,864):+&mul(957,923)'don't()+select()why()mul(496,383),&why()(-;)+!#mul(79,135)~{mul(500,619),#what(361,47)@&;@+mul(580,98)<>#why():$from()(!mul(660,615):where()mul(596,918)how()[what()how()who()]@mul(594,699)$>$&why()mul(951,813)mul(948,41)]mul(699,225)>mul(840,275):@~^@%^select()mul(695,594)what()->-mul(409,636)@ {++mul(999,290)when(339,689),[+[+?mul(558,848)how(),what()&}]{mul(344,337)$)$]+,$,,why()do()::select()^:mul(796,16)mul(526,718) mul(66,417)#@>~do()why()from()/select()when():&mul(8,111);!where()do()}@mul(433,239)*}how() {^!do()#^mul(167,905)-&%(mul(979,711)when()]who()'?mul(523,172)[^what(844,455)'what()[/~@mul(928,776)where()]&/mul(587,380)who()/when()how()from(),}#}mul(97,952)select(837,505)~@>~%mul(486,198)^!mul(968,779)do()[mul(206,321);,where()where())how()}/mul(6,867)-&<!*select()mul(615,984)~?who()#<{how()?@mul(617,643){ @,'why()+>what()(mul(683,583)^(mul(244,555)^where()(how()what()?-[,who()mul(827,305);&,!$*?mul(793,452)^ )+(%-where()<mul(318,238)<$-what()mul(802,394)?how()<mul(488,879)do()$}<:&mul(429,358)~how()mul(36,174)mul(87,398)+%what()]mul(822,311)mul(379,261)~when()mul(618,193)@)!))-what()how(645,829)[mul(964,102);who()don't(),[what(),,?select()why()@mul(121,293)select()@who()!mul(770,96)+;%from(467,168)mul(616,890)@%:)[mul(943,145) [from()&where(283,130){*where()mul(649,49)select())mul(877,73)from()do()where()?where()>who()from()@who(){*mul(592,919)/:what():who()mul(213,139)what()$!?*how(934,24)$-who()(mul(661,7)?from(),}/who()mul(876,585)how()why():when()[@@mul(433,918)why()why(),*{who()when()who()mul(144,14/#(<mul(656,306)&what()&{*@%select(709,225)/mul(843,135)who()$[don't()<mul(576,966)^,?>don't()@(@{~why()[who()mul(394,492)where()/mul(508,797)-mul(150,832)+$@why()(why()~mul(527,236)<don't()!#{how()*!,!mul(177,945)'mul(5,181)#!#'&?what()&mul(10,492):%/'*mul(822,748){!*]who()when()<!<mul(21,920)+do()-who()from(505,694)where()from()$who()where()where()/mul(827,561)how()when()mul(770,221)%when()who()mul(872,151)mul(453,84)['&+:mul(258,293)@#,where()mul(264,689)mul(90[mul(94,665)-who()don't()what()>!:@!^&!mul(324>?!from()where())mul(731,184)#(}who()'mul(666,377)how()(;mul(346,141)^select()/@[$how()@]?mul(702,502,^where()<}[mul(872,308)+^!<@from() {:mul(537,974)%when()when()%mul(563,805)select()>how()>^';mul(643,555))where(290,606)who()?mul(656,803),?$/*?when()}@ mul(228,902)~/from(){where()what()>@~mul(849,93)~+select()?-^}<don't()mul(844,117)(when()~,}# /:+mul(24,282)@where()~from()($what()what()*<mul(886,107)[#;mul(725,121) %[>mul(115,425)${what()mul(219,305)how()])from()<mul(401,458)-from()+mul(92,705}",
	}
	answer := stepTwo(data)
	assert.Equal(t, 100, answer)
}

func Test_runStepTwo(t *testing.T) {
	logrus.SetLevel(logrus.InfoLevel)
	vals, err := parseData()

	if err != nil {
		logrus.Errorf("error parsing data: [%s]", err)
		return
	}

	answer := stepTwo(vals)
	assert.Equal(t, 161289189, answer)
	logrus.Infof("step2: [%d]", answer)
}
