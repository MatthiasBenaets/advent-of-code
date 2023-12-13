package main

import (
	"bufio"
	"fmt"
	"os"
  "regexp"
)

var data []string = []string{
  "LR",
  "",
  "11A = (11B, XXX)",
  "11B = (XXX, 11Z)",
  "11Z = (11B, XXX)",
  "22A = (22B, XXX)",
  "22B = (22C, 22C)",
  "22C = (22Z, 22Z)",
  "22Z = (22B, 22B)",
  "XXX = (XXX, XXX)",

}
var conv [][]string
var convert []string
var point []string
var count int
var found bool = false
var counter int

func main(){
  data = read()

  // Regex to find 3 strings that are only capital letter
	re := regexp.MustCompile(`[A-Z0-9]{3}`)
  for i := 2; i < len(data); i++ {
	  convert = re.FindAllString(data[i], -1)
    conv = append(conv, convert)
  }

  direction := data[0]

  for _, step := range conv {
    if string(step[0][2]) == "A" {
      point = append(point, string(step[0]))
    }
  }

  // Loop through directions
  for i := 0; i < len(direction); i++ {
    // Loop through all direction points
    for j := 0; j < len(point); j++{
      // Loop through map
      for _, step := range conv {
        // If first value is point
        if step[0] == point[j] {
          // Select new point depending on direction
          if string(direction[i]) == "L" {
            point[j] = step[1]
          } else if string(direction[i]) == "R" {
            point[j] = step[2]
          }
          count++
          break
        }
      }

      var zcount int
      // Check if points end with Z
      for k := 0; k < len(point); k++{
        if string(point[k][2]) == "Z" {
          zcount++
        }
      }
      counter++
      // If all Z, path found
      if zcount == len(point) {
        found = true;
        break
      }
    }
    // If at end of directions but not found
    if i == len(direction) - 1 && found == false{
      i = -1
    }

    // Found = done
    if found == true {
      break
    }
  }
  fmt.Println(count - 1)
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

