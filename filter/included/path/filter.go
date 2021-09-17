package path

import (
	"os"
	"strings"
)

type f struct{}

func (f) Filter(keyIn, valueIn string) (keyOut, valueOut string, err error) {
	if keyIn != "PATH" {
		return keyIn, valueIn, nil
	}

	paths := make(map[string]bool)
	path := os.Getenv("PATH")
	valueIn = strings.ReplaceAll(valueIn, "$PATH", "")
	path = valueIn + ":" + path

	for _, s := range strings.Split(path, ":") {
		if s != "" {
			paths[s] = true
		}
	}

	b := new(strings.Builder)

	needsSeparator := false
	for k := range paths {
		if needsSeparator {
			b.WriteString(":")
		} else {
			needsSeparator = true
		}
		b.WriteString(k)
	}

	return keyIn, b.String(), nil
}
