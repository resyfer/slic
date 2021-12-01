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

    // Ignore specified items
    if doIgnore(name) {
      continue
    }

    if item.IsDir() {
      
      // If directory, go down recursively
      traverseDirectory(directory + "/" + item.Name())

    } else {

      //If not a directory, add to language counter.
      
      extensionIndex := strings.IndexRune(name, '.')

      extension := name[extensionIndex + 1:] //Grabbing extension from name

      // Increasing the respective size counter of language
      language, found := EXTENSIONS[extension]
      itemInfo, _ := item.Info()
      if found {
        LANG[language]+= itemInfo.Size()
      } else {
        LANG["Other"]+= itemInfo.Size()
      }
      TOTAL_SIZE+= itemInfo.Size() // Increase total size counter
    }
  }
}