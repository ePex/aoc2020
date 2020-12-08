package utils

import (
	"bufio"
	"log"
	"os"
)

func FileInt(filePath string) (values []int) {
	return ToIntArray(FileString(filePath))
}

func FileString(filePath string) (values []string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return values
}