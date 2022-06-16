package file

import (
	"os"
	"strings"
)

func ReadFile(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	rawString := string(file)

	return strings.Split(strings.TrimRight(rawString, "\r\n"), "\n")
}

func SplitOnce(str string, seperator string) (string, string) {
	slice := strings.Split(str, seperator)

	return slice[0], slice[1]
}
