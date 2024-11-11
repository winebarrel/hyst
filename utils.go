package hyst

import (
	"bufio"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SummarizeByWidth(nums []float64, width int) []int {
	ints := make([]int, 0, len(nums))

	for _, n := range nums {
		trancated := int(math.Trunc(n))
		m := trancated % width
		ints = append(ints, trancated-m)
	}

	return ints
}

func ReadStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := io.ReadAll(reader)

	if err != nil {
		return "", err
	}

	return string(input), nil
}

func Split(str string) []string {
	splitted := strings.Split(str, "\n")
	values := make([]string, 0, len(splitted))

	for _, v := range splitted {
		if v != "" {
			values = append(values, v)
		}
	}

	return values
}

func StrsToNums(strs []string) ([]float64, error) {
	nums := make([]float64, 0, len(strs))

	for _, s := range strs {
		f, err := strconv.ParseFloat(s, 64)

		if err != nil {
			return nil, err
		}

		nums = append(nums, f)
	}

	return nums, nil
}

func IntsToStrs(ints []int) []string {
	strs := make([]string, 0, len(ints))

	for _, i := range ints {
		s := strconv.Itoa(i)
		strs = append(strs, s)
	}

	return strs
}

func Max(ints []int) int {
	max := 0

	for _, i := range ints {
		if i > max {
			max = i
		}
	}

	return max
}

func Min(ints []int) int {
	min := ints[0]

	for _, i := range ints {
		if i < min {
			min = i
		}
	}

	return min
}

func Bar(cnt int, width int, max int) string {
	if cnt < 1 {
		return ""
	}

	fraction := float64(width) / float64(max)

	if fraction > 1.0 {
		fraction = 1.0
	}

	length := int(float64(cnt) * fraction)
	bar := make([]byte, 0, length)

	for i := 0; i < length; i++ {
		bar = append(bar, '#')
	}

	return string(bar)
}

func MapKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func MaxStrLen(strs []string) int {
	max := 0

	for _, s := range strs {
		l := len(s)

		if l > max {
			max = l
		}
	}

	return max
}

func MapVals(m map[string]int) []int {
	vals := make([]int, 0, len(m))

	for _, v := range m {
		vals = append(vals, v)
	}

	return vals
}

type AsNum []string

func (an AsNum) Len() int {
	return len(an)
}

func (an AsNum) Swap(i, j int) {
	an[i], an[j] = an[j], an[i]
}

func (an AsNum) Less(i, j int) bool {
	fi, err := strconv.ParseFloat(an[i], 64)

	if err != nil {
		panic(err)
	}

	fj, err := strconv.ParseFloat(an[j], 64)

	if err != nil {
		panic(err)
	}

	return fi < fj
}

func SortAsNums(strs []string) {
	sort.Sort(AsNum(strs))
}

type Pair struct {
	Key   string
	Value int
}

type ByValue []Pair

func (ps ByValue) Len() int {
	return len(ps)
}

func (ps ByValue) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps ByValue) Less(i, j int) bool {
	return ps[i].Value < ps[j].Value
}

func ReverseSortKeyByVal(m map[string]int) []string {
	ps := make([]Pair, 0, len(m))
	keys := make([]string, 0, len(m))

	for k, v := range m {
		pair := Pair{k, v}
		ps = append(ps, pair)
	}

	sort.Sort(sort.Reverse(ByValue(ps)))

	for _, p := range ps {
		keys = append(keys, p.Key)
	}

	return keys
}

func CompleteInts(ints []int, width int) []int {
	m := make(map[int]bool, len(ints))
	completed := []int{}

	for _, i := range ints {
		m[i] = true
	}

	max := Max(ints)
	min := Min(ints)

	for i := min; i < max; i += width {
		if _, ok := m[i]; !ok {
			completed = append(completed, i)
		}
	}

	return completed
}
