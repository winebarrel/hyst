package hyst

import (
	"flag"
	"fmt"
)

func ParseFlag(h *Hyst) error {
	flag.IntVar(&h.Width, "w", 0, "Width")
	flag.Parse()

	if h.Width < 0 {
		fmt.Errorf("Width must be >= 0")
	}

	return nil
}
