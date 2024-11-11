package hyst

import (
	"errors"
	"flag"
)

func ParseFlag(h *Hyst) error {
	flag.IntVar(&h.Width, "w", 0, "Width")
	flag.Parse()

	if h.Width < 0 {
		return errors.New("Width must be >= 0")
	}

	return nil
}
