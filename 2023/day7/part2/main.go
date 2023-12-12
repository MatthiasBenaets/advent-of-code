package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var data []string = []string{
  "32T3K 765",
  "T55J5 684",
  "KK677 28",
  "KTJJT 220",
  "QQQJA 483",
}
var value string
var result int

type Hand struct {
  Score int
  Value string
  Multi int
}

var customOrder = map[byte]int{
	'J': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T':10, 'Q': 11, 'K': 12, 'A': 13,
}

var count2Pairs = 0
var countFullHouse = 0
var countHighCard = 0

func main(){
  data = read()

  // Create array of Hand
  myHand := []Hand{}

  // Loop through all data
  for _, line := range data {
    score := 0
    // Split at space
    hands := strings.Split(line, " ")
    // Sort by hand
    value = sorting(hands[0])
    // Count how often there is an x
    countPairs := countOccurrences(value, 2)
		countThreeOfAKind := countOccurrences(value, 3)
		countFourOfAKind := countOccurrences(value, 4)
		countFiveOfAKind := countOccurrences(value, 5)
    countJack := countJacks(value)

    // fmt.Printf("jacks counted: %v \n", countJack)

    // Give a score value and convert if needed
    switch {
    case countFiveOfAKind == 1:
      score = 6
    case countFourOfAKind == 1:
      score = 5
    case countPairs == 1 && countThreeOfAKind == 1:
      score = 4
    case countThreeOfAKind == 1:
      score = 3
    case countPairs == 2:
      score = 2
    case countPairs == 1:
      score = 1
    case countPairs == 0 && countThreeOfAKind == 0 && countFourOfAKind == 0  && countFiveOfAKind == 0:
      score = 0
    default:
      break
    }

    // Convert score depening on Jacks
    switch {
    case score == 5:
      if countJack == 1 {
        score = 6
      } else {
        score = 5
      }
    case score == 3:
      if countJack == 1 {
        score = 5
      } else if countJack == 2 {
        score = 6
      } else {
        score = 3
      }
    case score == 2:
      if countJack == 1 {
        score = 4
      } else {
        score = 2
      }
    case score == 1:
      if countJack == 1 {
        score = 3
      } else if countJack == 2 {
        score = 5
      } else if countJack == 3 {
        score = 6
      } else {
        score = 1
      }
    case score == 0:
      if countJack == 5 || countJack == 4 {
        score = 6
      } else if countJack == 3 {
        score = 5
      } else if countJack == 2 {
        score = 3
      } else if countJack == 1 {
        score = 1
      } else {
        score = 0
      }
    default:
      break
    }

    // Convert multiplier to number
    convert, err := strconv.Atoi(hands[1])
    if err != nil {
      return
    }

    newHand := Hand{score, hands[0], convert}
    myHand = append(myHand, newHand)

    // Sort hand (with score and converted multiplier). First by Score, then by value (per grouped score)
    sort.SliceStable(myHand, func(i, j int) bool {
      if myHand[i].Score != myHand[j].Score {
        return myHand[i].Score < myHand[j].Score
      }
      return customSort(myHand[i].Value, myHand[j].Value)
    })
  }

  // After sorting, add the position by the multiplier
  for i := range myHand {
    result += (i + 1)*myHand[i].Multi
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

// Sort hand
func sorting(values string) string {
  // Convert hand to array of runes
  value := []rune(values)
  // Sort the cards from small to big
	sort.Slice(value, func(i, j int) bool {
		return value[i] < value[j]
	})
	return string(value)
}

// Count how many of target in hand
func countOccurrences(sortedValue string, targetCount int) int {
	count := 0
	i := 0
	for i < len(sortedValue)-1 {
		j := i + 1
		for j < len(sortedValue) && sortedValue[i] == sortedValue[j] {
			j++
		}
		if j-i == targetCount {
      // Only count if not a Jack
      if string(sortedValue[i]) != "J" {
			  count++
      }
		}
		i = j
	}
	return count
}

func customSort(a, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		// Compare characters using the custom order map
		if customOrder[a[i]] != customOrder[b[j]] {
			return customOrder[a[i]] < customOrder[b[j]]
		}
		// Move to the next character
		i++
		j++
	}

	// If one string is a prefix of the other, shorter is considered "less"
	return len(a) < len(b)
}

// Count Jacks
func countJacks(value string) int{
  var jacks = 0
  // Convert hand to runes
  values := []rune(value)
  // Loop through all cards and look for Jacks
  for i:=0; i < len(values); i++ {
    if string(value[i]) == "J" {
      jacks++
    }
  }
  return jacks
}
