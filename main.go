package main

import (
  "flag"
  "fmt"
  "net/http"
  // "io"
  // "os"

  "github.com/alishalabi/link_parser"
)


func main() {
  urlFlag := flag.String("url", "https://twitch.tv", "Website for which you want to build a sitemap")
  flag.Parse()

  // Initiate GET request
  resp, err := http.Get(*urlFlag)
  if err != nil {
    panic(err)
  }
  // Must ALWAYS close response body, or memory leak will occur
  defer resp.Body.Close() // defer will run when function ends

  // io.Copy(os.Stdout, resp.Body)

  links, _ := link.Parse(resp.Body)
  for _, l := range links {
    fmt.Println(l)
  }

  // Case 1: Path no domain (/path)
  // Case 2: Path with domain (https://makeschool.com/path)
  // Case 3: Fragment (#fragment)
  // Case 4: Email (mailto:myemail@gmail.com)


}
