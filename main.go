package main

import (
  "log"
  "fmt"
  "github.com/jimgong92/geddit/reddit"
)

func main() {
  items, err := reddit.Get("all")
  if (err != nil) {
    log.Fatal(err)
  }
  for _, item := range items {
    fmt.Println(item)
  }
}