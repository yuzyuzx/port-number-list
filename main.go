package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

const portNumberText = "./port-number.txt"

func main() {
  data, _ := os.Open(portNumberText)
  defer data.Close()

  scanner := bufio.NewScanner(data)

  for scanner.Scan() {
    s := scanner.Text()
    s = strings.ReplaceAll(s, "\t", ",")
    fmt.Println(s)
  }

}
