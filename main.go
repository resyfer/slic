package main

import (
  "github.com/gookit/color"
)

func main() {
  traverseDirectory(".")

  for elem, val := range LANG {
    if val != 0 {
      color.HEX(COLORS[elem]).Printf("%v : %0.2f%% \n", elem, float64(val) / float64(TOTAL_SIZE) * 100)
    }
  }
}