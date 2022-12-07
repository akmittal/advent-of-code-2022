package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./day1/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	max, curr := 0, 0
	for i := 0; i < len(lines); i++ {

		if lines[i] == "" {
			if curr > max {
				max = curr
			}
			curr = 0
		} else {
			num, _ := strconv.Atoi(lines[i])
			curr += num
		}

	}
	fmt.Println(max)

}
