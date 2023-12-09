package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	timeFunction(solve)
}

// Start timer
func timeFunction(function func()) {
	start := time.Now()
  solve()
	fmt.Println(time.Since(start))
}

// Error checker
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func solve() {
	// Read file
	data, err := os.ReadFile("../data")
  // Check for errors
	check(err)
  // Remove trailing spaces and splits content into lines
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	part1 := 0
  // Declare regexp, will find all numbers
	re := regexp.MustCompile(`(\d+)`)
  // Initialize map, keys are strings and values are slices of int
	gears := make(map[string][]int)

  // For each line
	for index, line := range lines {
    // Search for number in a line and return the start and end index
		for _, match := range re.FindAllStringIndex(line, -1) {
      // Convert number found from the index
			num, _ := strconv.Atoi(line[match[0]:match[1]])
			isPartNo := false
      // Loop vertically: top , middle, bottom,
			for i := max(index-1, 0); i <= min(index+1, len(lines)-1); i++ {
        // Loop horizontally: left, middle, right
				for j := max(match[0]-1, 0); j <= min(match[1], len(line)-1); j++ {
          // Symbol is character in this location
					symbol := string(lines[i][j])
          // If symbol is not one of the below and thus valid
					if !strings.Contains("0123456789.", symbol) {
            // Enable so it can be added up in the end
						isPartNo = true
            // If symbol is a gear
						if symbol == "*" {
              // Convert argument into string which are coordinates of the gear
							gearing := fmt.Sprint(i, "-", j)
              // If length of gears[gearing] = 0 so no entry
							if len(gears[gearing]) == 0 {
                // One number
								gears[gearing] = []int{1, num}
							} else {
                // Two numbers
								gears[gearing][0]++
                // Multiply number
								gears[gearing][1] *= num
							}
						}
					}
				}
			}
      // Add up value
			if isPartNo {
				part1 += num
			}
		}
	}

	part2 := 0
  // Loop through gears
	for _, gear := range gears {
    // If 2 gears, add up all the multipications
		if gear[0] == 2 {
			part2 += gear[1]
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
