package main

import (
	"flag"
	"sort"
	"strings"

	"github.com/gookit/color"
)

var VERSION = "1.0.0"
var containedLanguages = []Language {}

type Language struct {
  name string
  percentage float64
}

func main() {
  
  ignore := flag.String("i", "", "Names of Files/Directories (only names, not paths) to ignore (write comma-separated names eg. -i=foo,bar.txt ) ")
  directory := flag.String("d", ".", "Relative Path of Directory of Search (eg. -d=./src) ")
  flag.Parse()

  color.HEX("#F1E05A").Printf("Welcome to SLIC (Show Languages In Code) v%v\n\n", VERSION)
  color.HEX("#9400FF").Printf("The Languages Are:\n")


  addIgnoreItems(*ignore)

  traverseDirectory(*directory)

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

// Add ignore items
func addIgnoreItems(ignoreItems string) {

  if ignoreItems == "" {
    return
  }

  IGNORE = append(IGNORE, strings.Split(ignoreItems, ",")...)

}

// Print the langues
func printLanguages() {
  for _, elem := range containedLanguages {
    color.HEX(COLORS[elem.name]).Printf("%v : %0.2f%% \n", elem.name, elem.percentage)
  }
}