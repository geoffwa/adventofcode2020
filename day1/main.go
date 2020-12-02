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

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			a, b := numbers[i], numbers[j]
			if a + b == 2020 {
				fmt.Printf("%d + %d == 2020\n", a, b)
				fmt.Printf("%d * %d == %d\n",a, b, a * b)
			}
			for k := j + 1; k < len(numbers); k++ {
				c := numbers[k]
				if a + b + c == 2020 {
					fmt.Printf("%d + %d + %d == 2020\n", a, b, c)
					fmt.Printf("%d * %d * %d == %d\n",a, b, c, a * b * c)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
