package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numbers := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, int(num))
	}

	loop:
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] + numbers[j] == 2020 {
				fmt.Printf("%d + %d == 2020\n", numbers[i], numbers[j])
				fmt.Printf("%d * %d == %d\n", numbers[i], numbers[j], numbers[i] * numbers[j])
				break loop
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
