package main

import (
  "github.com/gookit/color"
)

func main() {
  traverseDirectory(".")

  for elem, val := range LANG {
    if val != 0 {
      color.HEX(COLORS[elem]).Printf("%v : %v%% \n", elem, float64(val) / float64(TOTAL_SIZE) * 100)
    }
  }
}