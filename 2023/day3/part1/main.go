package main

import (
  "fmt"
  "strconv"
  "os"
  "bufio"
)

var data []string = []string{
  "467..114..",
  "...*......",
  "..35..633.",
  "......#...",
  "617*......",
  ".....+.58.",
  "..592.....",
  "......755.",
  "...$.*....",
  ".664.598..",
}
// All coordinates around character
var directions = [][]int{
  {-1, -1}, {-1, 0}, {-1, 1},
  {0, -1}, {0, 1},
  {1, -1}, {1, 0}, {1, 1},
}
var num string
var symbol bool
var result int

func main() {

  // Load data file
  data = read()

  // Loop every line
  for i, line := range data {
    num = ""
    symbol = false
    // Loop every character in line
    for j := 0; j < len(line); j++ {
      // If character is number
      if checkNum(line[j]){
        // Add character to string
        num += string(line[j])

        // Check if symbol is near character
        if calc(data, i, j) {
          symbol = true
        }

        // Check if next number the end of the string or is NOT a number
        if j+1 == len(line) || !checkNum(line[j+1]){
          // If a symbol was found
          if symbol == true {
            // Convert to an actual number
            conv, err := strconv.Atoi(num)
            if err != nil {
              fmt.Println("Error converting", err)
            }
            // Add up the results and reset variables
            result += conv
            num = ""
            symbol = false
          } else {
            num = ""
          }
        }
      }
    }
  }
  fmt.Println(result)
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

// Check if character is a number
func checkNum(char byte) bool {
  return char >= '0' && char <= '9'
}
// Check if character is a "."
func checkDot(char byte) bool {
  return char == '.'
}
func calc(data []string, row, col int) bool {
  // Iterate through all directions
	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]

		// Check if the new position is within the bounds of the data
		if newRow >= 0 && newRow < len(data) && newCol >= 0 && newCol < len(data[0]) {
			// Check if the character at the new position is a symbol or number
      if !checkDot(data[newRow][newCol]) {
        if !checkNum(data[newRow][newCol]) {
				  return true
        }
			}
		}
	}
	return false
}
