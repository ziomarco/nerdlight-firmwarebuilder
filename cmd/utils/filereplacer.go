package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReplaceInFile(filePath string, searchTerm string, replaceTerm string) (bool, error) {
	input, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
		return err != nil, err
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, searchTerm) {
			lines[i] = strings.ReplaceAll(line, searchTerm, replaceTerm)
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
		return err != nil, err
	}
	return err != nil, err
}
