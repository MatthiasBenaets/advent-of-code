package main

import (
	"fmt"
	"strconv"
	"strings"
)

var data []string = []string{
  "Time:        44     82     69     81",
  "Distance:   202   1076   1138   1458",
  // "Time:      7  15   30",
  // "Distance:  9  40  200",
}

var intValues []int
var counter int
var final int = 1

func main(){
  for _, line := range data {
    split := strings.Fields(line)
    strValues := split[1:]

    for _, strVal := range strValues {
      intVal, err := strconv.Atoi(strVal)
      if err != nil {
        return
      }
      intValues = append(intValues, intVal)
    }
  }

  // Loop races
  for i:=0; i<=3; i++ {
    var x int = 0
    var y int = 0
    counter = 0
    // Loop each button press
    for j:=1; j<=intValues[i]; j++ {
      // Time left
      x = intValues[i] - j
      // Calculate distance
      y = x * j
      // If further
      if y > intValues[i+4] {
        counter++
      }
    }
    final *= counter
  }
  fmt.Println(final)
}
