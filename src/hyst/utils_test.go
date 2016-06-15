package hyst

import (
	"github.com/stretchr/testify/assert"
	. "hyst"
	"sort"
	"testing"
)

func TestSplit(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"foo", "bar", "zoo"}
	actual := Split("\nfoo\nbar\nzoo\n")
	assert.Equal(expected, actual)
}

func TestStrsToNums(t *testing.T) {
	assert := assert.New(t)
	expected := []float64{1, 2, 3}
	actual, _ := StrsToNums([]string{"1", "2", "3"})
	assert.Equal(expected, actual)
}

func TestSummarizeByWidth(t *testing.T) {
	assert := assert.New(t)
	expected := []int{0, 0, 3, 3, 3, 3, 6, 6, 6}
	actual := SummarizeByWidth([]float64{0, 2.5, 3, 3.1, 4.5, 5.9, 6, 7.3, 8.6}, 3)
	assert.Equal(expected, actual)
}

func TestIntsToStrs(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"1", "2", "3"}
	actual := IntsToStrs([]int{1, 2, 3})
	assert.Equal(expected, actual)
}

func TestMax(t *testing.T) {
	assert := assert.New(t)
	expected := 3
	actual := Max([]int{2, 3, 1, 0})
	assert.Equal(expected, actual)
}

func TestMin(t *testing.T) {
	assert := assert.New(t)
	expected := 1
	actual := Min([]int{2, 3, 1, 4})
	assert.Equal(expected, actual)
}

func TestBar(t *testing.T) {
	assert := assert.New(t)
	width := 50
	max := 300
	assert.Equal("", Bar(0, width, max))
	assert.Equal("#", Bar(10, width, max))
	assert.Equal("#####", Bar(30, width, max))
	assert.Equal("############################", Bar(170, width, max))
	assert.Equal("##################################################", Bar(300, width, max))
}

func TestMapKeys(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"bar", "foo", "zoo"}
	actual := MapKeys(map[string]int{"foo": 2, "bar": 1, "zoo": 3})
	sort.Strings(actual)
	assert.Equal(expected, actual)
}

func TestMaxStrLen(t *testing.T) {
	assert := assert.New(t)
	expected := 6
	actual := MaxStrLen([]string{"foofo", "barbar", "zoo"})
	assert.Equal(expected, actual)
}

func TestMapVals(t *testing.T) {
	assert := assert.New(t)
	expected := []int{1, 2, 3}
	actual := MapVals(map[string]int{"foo": 2, "bar": 1, "zoo": 3})
	sort.Ints(actual)
	assert.Equal(expected, actual)
}

func TestSortAsNums(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"1.1", "2.2", "3.3"}
	actual := []string{"3.3", "1.1", "2.2"}
	SortAsNums(actual)
	assert.Equal(expected, actual)
}

func TestReverseSortKeyByVal(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"foo", "zoo", "bar"}
	actual := ReverseSortKeyByVal(map[string]int{"foo": 3, "bar": 1, "zoo": 2})
	assert.Equal(expected, actual)
}

func TestCompleteInts(t *testing.T) {
	assert := assert.New(t)
	expected := []int{3, 6, 9, 15, 18, 21, 24, 27}
	actual := CompleteInts([]int{0, 12, 30}, 3)
	assert.Equal(expected, actual)
}
