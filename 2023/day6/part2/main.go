package main

import (
	"fmt"
	"strconv"
	"strings"
)

var data []string = []string{
  "Time:        44     82     69     81",
  "Distance:   202   1076   1138   1458",
}

var intValues []int
var final int

func main(){
  for _, line := range data {
    // Remove spaces
    fix := strings.ReplaceAll(line," ","")
    // Split at ":"
    split := strings.Split(fix,":")
    strValues := split[1:]

    // Convert to number
    for _, strVal := range strValues {
      intVal, err := strconv.Atoi(strVal)
      if err != nil {
        return
      }
      intValues = append(intValues, intVal)
    }
  }

  var x int = 0
  var y int = 0
  counter := 0
  // Loop each button press
  for j:=1; j<=intValues[0]; j++ {
    // Time left
    x = intValues[0] - j
    // Calculate distance
    y = x * j
    // If further
    if y > intValues[1] {
      counter++
    }
  }
  final = counter
  fmt.Println(final)
}
