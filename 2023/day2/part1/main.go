package main

import (
  "fmt"
  "strings"
  "strconv"
  "os"
  "bufio"
)

var data []string = []string{
"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

const red = 12
const green = 13
const blue = 14;
var result int

func main() {
  // Load dummy data
  // for i := range data {
  //   if calc(data[i]) == 1 {
  //     result = result + i + 1
  //   }
  // }
  // fmt.Println(result)

  // Load data file
  fmt.Println(read())
}

func read() int{
  // Load data
  file, err := os.Open("../data")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  // Load lines
  var counter int
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    counter++
    // Calculate results
    if calc(line) == 1 {
      result = result + counter
    }
  }

  return result
}

func calc(data string) int{
  var rolls = map[string]int{
    "red": 0,
    "green": 0,
    "blue": 0,
  }

  // Split string at ": "
  game := strings.Split(data, ": ")

  // Ignore column 0, split multiple times
  for _, minigames := range game[1:] {
    minigame := strings.Split(minigames, "; ")
    for _, roles := range minigame {
    role := strings.Split(roles, ", ")
      for _, value := range role {
        split := strings.Split(value, " ")
        // Convert to number
        number, err := strconv.Atoi(split[0])
        if err != nil {
          fmt.Println("Error converting", err)
          return 0
        }

        // Assign number to map color
        rolls[split[1]] = number

        // If any number is larger, invalid game
        if rolls["red"] > red || rolls["green"]>green || rolls["blue"] > blue {
          return 0
        }
      }
    }
  }
  return 1
}
