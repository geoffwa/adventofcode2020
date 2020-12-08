package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<letter>[a-z]): (?P<password>[a-z]+)`)
	scanner := bufio.NewScanner(file)
	validCount := 0
	for scanner.Scan() {
		match := r.FindStringSubmatch(scanner.Text())
		if match == nil {
			continue
		}
		firstIdx, err := strconv.ParseInt(match[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		secondIdx, err := strconv.ParseInt(match[2], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		letter, password := match[3], match[4]
		count := 0
		if password[firstIdx-1] == letter[0] {
			count++
		}
		if password[secondIdx-1] == letter[0] {
			count++
		}
		if count == 1 {
			validCount++
		}
	}

	fmt.Printf("%d valid passwords\n", validCount)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
