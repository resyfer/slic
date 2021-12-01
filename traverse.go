package main

import (
	"os"
	"strings"
)

var TOTAL_SIZE int64 = 0

func traverseDirectory(directory string) {
  contents, _ := os.ReadDir(directory)

  for _, item := range contents {
    
    name := item.Name()

    // Ignore some items
    if doIgnore(name) {
      continue
    }

    // If directory, go down recursively. If not, add to language counter.
    if item.IsDir() {
      traverseDirectory(directory + "/" + item.Name())
    } else {
      extensionIndex := strings.IndexRune(name, '.')

      extension := name[extensionIndex + 1:]

      // Increasing the respective counter of language
      language, found := EXTENSIONS[extension]
      itemInfo, _ := item.Info()
      if found {
        LANG[language]+= itemInfo.Size()
      } else {
        LANG["Other"]+= itemInfo.Size()
      }
      TOTAL_SIZE+= itemInfo.Size()
    }
  }
}