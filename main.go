package main

import "fmt"

func main() {
  traverseDirectory(".")

  for elem, val := range LANG {
    if val != 0 {
      fmt.Printf("%v : %v%% \n", elem, float64(val) / float64(TOTAL_SIZE) * 100)
    }
  }
}