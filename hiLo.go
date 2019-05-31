package main

import (
  "time"
  "fmt"
  "math/rand"
  "strconv"
  "os"
)

func newRand() *rand.Rand {
  t := time.Now().UnixNano()
  s := rand.NewSource(t)
  return rand.New(s)
}

func readInt(prompt string) (i int) {
  var err error
  for {
    fmt.Print(prompt)
    _, err = fmt.Scanf("%d\n", &i)
    if err == nil {
      break
    } else {
      fmt.Println("Error:", err)
    }
  }
  return
}

func play(min int, max int) {
  r := newRand()
  target := min + r.Intn(max - min)
  for {
    guess := readInt(fmt.Sprintf("Enter a number between %d and %d: ", min, max))
    if guess == target {
      fmt.Println("Correct!")
      break
    } else if guess < target {
      fmt.Println("Too low.")
    } else {
      fmt.Println("Too high.")
    }
  }
}

func main() {
  if len(os.Args) != 3 {
    fmt.Println("Arguments: min max")
    return
  }
  min, errMin := strconv.Atoi(os.Args[1])
  max, errMax := strconv.Atoi(os.Args[2])
  if errMin != nil || errMax != nil {
    fmt.Println("Error: min and max must be integers")
    return
  }
  play(min, max)
}
