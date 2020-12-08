package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type treeMap struct {
	trees []string
}

func (t *treeMap) isTreeAt(x, y int) bool {
	row := t.trees[y]
	return row[x%len(row)] == '#'
}

func (t *treeMap) depth() int {
	return len(t.trees)
}

func readTreeMapFile(filename string) *treeMap {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return &treeMap{trees: lines}
}

type slope struct {
	xIncr int
	yIncr int
}

func main() {
	filename := os.Args[1]
	treeMap := readTreeMapFile(filename)

	slopes := []slope{
		slope{1, 1},
		slope{3, 1},
		slope{5, 1},
		slope{7, 1},
		slope{1, 2},
	}
	for _, slope := range slopes {
		treeCount := 0
		for y, x := 0, 0; y < treeMap.depth(); y, x = y+slope.yIncr, x+slope.xIncr {
			if treeMap.isTreeAt(x, y) {
				treeCount++
			}
		}
		fmt.Printf("hit %d trees with xIncr: %d, yIncr: %d\n", treeCount, slope.xIncr, slope.yIncr)
	}

}
