package hyst

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	. "hyst"
	"testing"
)

func TestCalc(t *testing.T) {
	assert := assert.New(t)
	h := &Hyst{Width: 0}
	expected := map[string]int{"foo": 2, "bar": 1, "zoo": 3}
	actual, _ := h.Calc([]string{"foo", "foo", "bar", "zoo", "zoo", "zoo"})
	assert.Equal(expected, actual)
}

func TestCalcWithWidth(t *testing.T) {
	assert := assert.New(t)
	h := &Hyst{Width: 3}
	expected := map[string]int{"0": 2, "3": 4, "6": 3}
	actual, _ := h.Calc([]string{"0", "2.5", "3", "3.1", "4.5", "5.9", "6", "7.3", "8.6"})
	assert.Equal(expected, actual)
}

func TestCalcWithComplete(t *testing.T) {
	assert := assert.New(t)
	h := &Hyst{Width: 3}
	expected := map[string]int{"3": 0, "6": 0, "9": 0, "15": 0, "18": 0, "0": 2, "12": 4, "21": 3}
	actual, _ := h.Calc([]string{"0", "2.5", "12", "12.1", "13.5", "14.9", "21", "22.3", "23.6"})
	assert.Equal(expected, actual)
}

func TestDraw(t *testing.T) {
	assert := assert.New(t)
	buf := &bytes.Buffer{}
	h := &Hyst{Width: 0, Writer: buf}

	expected := `bar  300  ##################################################
zoo  200  #################################
baz  150  #########################
foo  100  ################
`

	h.Draw(map[string]int{"foo": 100, "bar": 300, "zoo": 200, "baz": 150})
	assert.Equal(expected, buf.String())
}

func TestDrawWithWidth(t *testing.T) {
	assert := assert.New(t)
	buf := &bytes.Buffer{}
	h := &Hyst{Width: 3, Writer: buf}

	expected := ` 3  100  ################
 6  300  ##################################################
 9  150  #########################
12  200  #################################
`

	h.Draw(map[string]int{"3": 100, "6": 300, "12": 200, "9": 150})
	assert.Equal(expected, buf.String())
}
