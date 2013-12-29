package main

import (
  "fmt"
  "time"
  "math/rand"
  "sync"
)

var mutex sync.Mutex

// returns a random int, of maximum max value
func get_random_int(max int) int {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  return r.Intn(max)
}

func process_element(i int, wg *sync.WaitGroup, result_acumulator map[int]int) {
  fmt.Println("checking value",i)

  ri := get_random_int(1000)
  time.Sleep(time.Duration(ri) * time.Millisecond)

  mutex.Lock()

  result_acumulator[i] = ri
  wg.Done()

  mutex.Unlock()
  fmt.Println("returning value for",i,ri)
}

func main() {
  values     := [5]int{1,2,3,4,5}
  results    := make(map[int]int)

  wg         := sync.WaitGroup{}
//  acumulator := make(chan int,5)

  for _, val := range values {
    wg.Add(1)
    go process_element(val, &wg, results)
  }

  wg.Wait()

  for key,val := range results {
    fmt.Println("key",key,"val",val)
  }
}