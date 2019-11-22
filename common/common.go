package common

import (
	"io/ioutil"
	"strings"
)

func PanicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFully(filepath string) string {
	dat, err := ioutil.ReadFile(filepath)
	PanicOnError(err)
	return string(dat)
}

func DelimitByNewLine(s string) []string {
	lines := strings.Split(s, "\n")
	for idx, freq := range lines {
		lines[idx] = strings.TrimSpace(freq)
	}
	return lines
}
