package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

var data []string = []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
var digits = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var result int

func main() {
  // Load dummy data
  // for i := range data {
  //   result += first(data[i]) * 10 + last(data[i])
  // }
  // fmt.Println(result)

  // Load file data
  fmt.Println(read())
}

func read() int {
  // Load data
  file, err := os.Open("../data")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  // Load lines
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    // Calculate result
    result += first(line) * 10 + last(line)
  }

  return result
}

// Find first number
func first(line string) int {
  var load string = ""

  // Loop over all character of line
	for i := 0; i < len(line); i++ {
    // If character is between 0 and 9
		if line[i] >= '0' && line[i] <= '9' {
      // Convert from utf-8 hex to number and exit function
			return int(line[i] - '0')
		}

    // Add character to new string
		load += string(line[i])

    // Check if end of string is in digits, return index number, exit function
		for i, digit := range digits {
			if strings.HasSuffix(load, digit) {
				return i
			}
		}
	}
	return 0
}

//Find last number
func last(line string) int {
  var load string = ""

  // Loop over all character of line starting at the end of the string
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
      // Convert from utf-8 hex to number and exit function
			return int(line[i] - '0')
		}

    // Add character to the start of the string
		load = string(line[i]) + load

    // Check if start of string is in digits, return index number, exit function
		for i, digit := range digits {
			if strings.HasPrefix(load, digit) {
				return i
			}
		}
	}
	return 0
}
