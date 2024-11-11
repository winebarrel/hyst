package hyst_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/hyst"
)

func TestSplit(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"foo", "bar", "zoo"}
	actual := hyst.Split("\nfoo\nbar\nzoo\n")
	assert.Equal(expected, actual)
}

func TestStrsToNums(t *testing.T) {
	assert := assert.New(t)
	expected := []float64{1, 2, 3}
	actual, _ := hyst.StrsToNums([]string{"1", "2", "3"})
	assert.Equal(expected, actual)
}

func TestSummarizeByWidth(t *testing.T) {
	assert := assert.New(t)
	expected := []int{0, 0, 3, 3, 3, 3, 6, 6, 6}
	actual := hyst.SummarizeByWidth([]float64{0, 2.5, 3, 3.1, 4.5, 5.9, 6, 7.3, 8.6}, 3)
	assert.Equal(expected, actual)
}

func TestIntsToStrs(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"1", "2", "3"}
	actual := hyst.IntsToStrs([]int{1, 2, 3})
	assert.Equal(expected, actual)
}

func TestMax(t *testing.T) {
	assert := assert.New(t)
	expected := 3
	actual := hyst.Max([]int{2, 3, 1, 0})
	assert.Equal(expected, actual)
}

func TestMin(t *testing.T) {
	assert := assert.New(t)
	expected := 1
	actual := hyst.Min([]int{2, 3, 1, 4})
	assert.Equal(expected, actual)
}

func TestBar(t *testing.T) {
	assert := assert.New(t)
	width := 50
	max := 300
	assert.Equal("", hyst.Bar(0, width, max))
	assert.Equal("#", hyst.Bar(10, width, max))
	assert.Equal("#####", hyst.Bar(30, width, max))
	assert.Equal("############################", hyst.Bar(170, width, max))
	assert.Equal("##################################################", hyst.Bar(300, width, max))
}

func TestMapKeys(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"bar", "foo", "zoo"}
	actual := hyst.MapKeys(map[string]int{"foo": 2, "bar": 1, "zoo": 3})
	sort.Strings(actual)
	assert.Equal(expected, actual)
}

func TestMaxStrLen(t *testing.T) {
	assert := assert.New(t)
	expected := 6
	actual := hyst.MaxStrLen([]string{"foofo", "barbar", "zoo"})
	assert.Equal(expected, actual)
}

func TestMapVals(t *testing.T) {
	assert := assert.New(t)
	expected := []int{1, 2, 3}
	actual := hyst.MapVals(map[string]int{"foo": 2, "bar": 1, "zoo": 3})
	sort.Ints(actual)
	assert.Equal(expected, actual)
}

func TestSortAsNums(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"1.1", "2.2", "3.3"}
	actual := []string{"3.3", "1.1", "2.2"}
	hyst.SortAsNums(actual)
	assert.Equal(expected, actual)
}

func TestReverseSortKeyByVal(t *testing.T) {
	assert := assert.New(t)
	expected := []string{"foo", "zoo", "bar"}
	actual := hyst.ReverseSortKeyByVal(map[string]int{"foo": 3, "bar": 1, "zoo": 2})
	assert.Equal(expected, actual)
}

func TestCompleteInts(t *testing.T) {
	assert := assert.New(t)
	expected := []int{3, 6, 9, 15, 18, 21, 24, 27}
	actual := hyst.CompleteInts([]int{0, 12, 30}, 3)
	assert.Equal(expected, actual)
}
