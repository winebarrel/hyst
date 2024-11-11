package main

import (
	"log"
	"os"

	"github.com/winebarrel/hyst"
)

func init() {
	log.SetFlags(0)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	h := &hyst.Hyst{Writer: os.Stdout}

	err := hyst.ParseFlag(h)

	if err != nil {
		log.Fatal(err)
	}

	input, err := hyst.ReadStdin()

	if err != nil {
		log.Fatal(err)
	}

	values := hyst.Split(input)
	calculated, err := h.Calc(values)

	if err != nil {
		log.Fatal(err)
	}

	err = h.Draw(calculated)

	if err != nil {
		log.Fatal(err)
	}
}
