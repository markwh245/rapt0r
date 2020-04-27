package files

import (
	"log"
	"os"
	"bufio"
)

func ReadFile(file string) []string {
	var lines []string

	fileData, err := os.Open(file)

	if err != nil {
		log.Fatalf("[+] Error open file: %v\n", err)
		os.Exit(1)
	}

	defer fileData.Close()

	scanner := bufio.NewScanner(fileData)
	
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}

	return lines
}
