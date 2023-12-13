package main

import (
	"bufio"
	"fmt"
	"os"
  "regexp"
)

var data []string = []string{
  "LLR",
  "",
  "AAA = (BBB, BBB)",
  "BBB = (AAA, ZZZ)",
  "ZZZ = (ZZZ, ZZZ)",
}
var conv [][]string
var convert []string
var point string = "AAA"
var count int
var found bool = false

func main(){
  data = read()

  // Regex to find 3 strings that are only capital letter
	re := regexp.MustCompile(`[A-Z]{3}`)
  for i := 2; i < len(data); i++ {
	  convert = re.FindAllString(data[i], -1)
    conv = append(conv, convert)
  }

  direction := data[0]

  // Loop through directions
  for i := 0; i < len(direction); i++ {
    // Loop through map
    for _, step := range conv {
      // If first value is point
      if step[0] == point {
        // Select new point depending on direction
        if string(direction[i]) == "L" {
          point = step[1]
        } else if string(direction[i]) == "R" {
          point = step[2]
        }
        count++
        if point == "ZZZ" {
          found = true
          break
        }
        break
      }
    }
    // If at end of directions but ZZZ not found
    if i == len(direction) - 1 && found == false{
      i = -1
    }
  }
  fmt.Println(count)
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

