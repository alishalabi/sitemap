package main

import (
  "flag"
  // "fmt"
  "net/http"
  "io"
  "os"

  // "github.com/alishalabi/link_parser"
)


func main() {
  urlFlag := flag.String("url", "https://makeschool.com", "Website for which you want to build a sitemap")
  flag.Parse()

  // Initiate GET request
  resp, err := http.Get(*urlFlag)
  if err != nil {
    panic(err)
  }
  // Must ALWAYS close response body, or memory leak will occur
  defer resp.Body.Close() // defer will run when function ends

  io.Copy(os.Stdout, resp.Body)







}
