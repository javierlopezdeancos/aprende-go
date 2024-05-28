package main

import (
	"fmt"
	"sync"
)

func write(texto string, wg *sync.WaitGroup) {
  fmt.Println(texto)
  defer wg.Done()
}

func main() {
  var wg sync.WaitGroup

  fmt.Println("hey")

  wg.Add(1)
  go write("hey again", &wg)

  wg.Wait()
}
