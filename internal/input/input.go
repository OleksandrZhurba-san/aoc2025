package input

import (
	"os"
	"strings"
)

func Read(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n"), nil
}
