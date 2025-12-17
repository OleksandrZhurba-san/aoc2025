package input

import (
	"os"
)

func Read(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func GetArgs(defaultPath string) (string) {
	if len(os.Args) > 1 {
		defaultPath = os.Args[1]
	}
	return defaultPath
}
