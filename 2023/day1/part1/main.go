package main

import (
  "fmt"
  "unicode"
  "strconv"
  "os"
  "bufio"
)

var data [4]string = [4]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
var result int

func main() {
  // for i := range data {
  //   result += count(data[i])
  // }
  // fmt.Println(result)
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
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    // Calculate result
    result += count(line)
  }

  return result
}

func count(line string) int {
  var sum int
  var firstNum int8 = 0
  var secondNum int8 = 0
  var firstPick bool = false
  var secondPick bool = false

  // Loop over each character of the line
  for _ , char := range line {
    // If character is a number, convert to actual number
    if unicode.IsNumber(char) {
      num, err := strconv.Atoi(string(char))
      if err != nil {
        fmt.Println("Error converting", err)
        return 0
      }
      // Check if first number or not
      switch {
        case !firstPick:
          firstNum = int8(num*10)
          firstPick = true
        default:
          secondNum = int8(num)
          secondPick = true
      }
      // If no additional number is found, first number becomse also second number
      if !secondPick {
        secondNum = int8(num)
      }
    }
  }
  sum = int(firstNum + secondNum)

  return sum
}
