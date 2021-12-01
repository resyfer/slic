package main

import (
	"sort"

	"github.com/gookit/color"
)
var containedLanguages = []Language {}

func main() {
  traverseDirectory(".")

  addLanguagesToList()
  sortLanguages()
  printLanguages()

}

// Add them from LANG to make them ready for sorting.
func addLanguagesToList() {

  for language, size := range LANG {
    
    if size == 0 {
      continue
    }

    containedLanguages = append(containedLanguages, Language{
      name: language,
      percentage: (float64(size) / float64(TOTAL_SIZE)) * 100,
    })

  }

}

// Sort the langues
func sortLanguages() {
  sort.Slice(containedLanguages, func(i, j int) bool {
    return containedLanguages[i].percentage > containedLanguages[j].percentage
  })
}

// Print the langues
func printLanguages() {
  for _, elem := range containedLanguages {
    color.HEX(COLORS[elem.name]).Printf("%v : %0.2f%% \n", elem.name, elem.percentage)
  }
}