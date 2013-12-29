package main

import (
  "fmt"
  "time"
  "math/rand"
  "sync"
)

// returns a random int, of maximum max value
func get_random_int(max int) int {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  return r.Intn(max)
}

func process_element(i int, acumulator chan int) {
  ri := get_random_int(4000)
  time.Sleep(time.Duration(ri) * time.Millisecond)
  acumulator <- ri
}

func main() {
  values := [5]int{1,2,3,4,5}

  fmt.Println(values)
}