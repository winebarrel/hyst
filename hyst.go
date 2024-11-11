package hyst

import (
	"fmt"
	"io"
)

const (
	Width = 50
)

type Hyst struct {
	Width  int
	Writer io.Writer
}

func (h *Hyst) Calc(values []string) (map[string]int, error) {
	var ints []int

	if h.Width > 0 {
		nums, err := StrsToNums(values)

		if err != nil {
			return nil, err
		}

		ints = SummarizeByWidth(nums, h.Width)
		values = IntsToStrs(ints)
	}

	m := map[string]int{}

	for _, v := range values {
		cnt, ok := m[v]

		if ok {
			m[v] = cnt + 1
		} else {
			m[v] = 1
		}
	}

	if h.Width > 0 {
		completed := CompleteInts(ints, h.Width)

		for _, i := range completed {
			v := fmt.Sprintf("%d", i)
			m[v] = 0
		}
	}

	return m, nil
}

func (h *Hyst) Draw(calculated map[string]int) error {
	keys := MapKeys(calculated)
	keyLen := MaxStrLen(keys)
	vals := MapVals(calculated)
	max := Max(vals)
	valLen := len(fmt.Sprintf("%d", max))

	if h.Width > 0 {
		SortAsNums(keys)
	} else {
		keys = ReverseSortKeyByVal(calculated)
	}

	for _, k := range keys {
		cnt := calculated[k]
		bar := Bar(cnt, Width, max)
		line := fmt.Sprintf("%*s  %*d  %s\n", keyLen, k, valLen, cnt, bar)
		_, err := io.WriteString(h.Writer, line)

		if err != nil {
			return err
		}
	}

	return nil
}
