package main

import (
	"fmt"
	"strconv"
	"strings"
  "os"
  "bufio"
)

var data []string = []string{
  "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
  "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
  "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
  "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
  "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
  "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}
var line []string
var result, final int

func main(){
  data = read()

  // For each string in data
  for i := range data {
    left, right := parseData(data[i])
    result = checker(left, right)
    final += result
  }
  fmt.Println(final)
}

func read() []string{
  // Load data
  file, err := os.Open("../data")
  if err != nil {
    fmt.Println("Error opening file:", err)
  }
  defer file.Close()

  // Read the file line by line
  scanner := bufio.NewScanner(file)

  var data []string
  // Create a string array
  for scanner.Scan() {
    line := scanner.Text()
    data = append(data, line)
  }

  return data
}

// Parse data into left and right array of integers
func parseData(data string) ([]int, []int){
  numbers :=  strings.Split(data, "|")
  left := parseNum(numbers[0])
  right := parseNum(numbers[1])
  return left, right
}

// Convert to integers
func parseNum(data string) []int{
  var numbers []int
  fields := strings.Fields(data)
  for _, fields := range fields {
    num, err := strconv.Atoi(fields)
    if err == nil {
      numbers = append(numbers, num)
    }
  }
  return numbers
}

// Count if number left is part of number right
func checker(left, right []int) int {
  var counter int = 0
  result = 0
  for _, num := range right {
    if contains(left, num) {
      if counter == 0 {
        result = 1
        counter += 1
      } else {
        result *= 2
      }
    }
  }
  return result
}

// Check if number is part of numbers
func contains(numbers []int, num int) bool{
  for _, e := range numbers {
    if e == num {
      return true
    }
  }
  return false
}
