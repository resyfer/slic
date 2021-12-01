package main

// Ignore the files/directories (Along with default Additions)
var IGNORE = []string{
  // Environmental Variables
  ".env",

  // Git
  ".git",

  // Go
  "go.mod",
  "go.sum",

  // JavaScript
  "node_modules",
  "package.json",
  "package-lock.json",

  // License
  "LICENSE",

  // Readme
  "README.md",

}

// Check if item needs to be ignored or not
func doIgnore(item string) bool {
  
  for i:=0; i<len(IGNORE); i++ {
    if item == IGNORE[i] {
      return true
    }
  }

  return false
}