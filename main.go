package main

import (
  "flag"
  "fmt"
  "net/http"
  "net/url"
  "strings"
  // "io"
  // "os"

  "github.com/alishalabi/link_parser"
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

  // io.Copy(os.Stdout, resp.Body)

  // links, _ := link.Parse(resp.Body)
  // for _, l := range links {
  //   fmt.Println(l)
  // }

  // Link cleanup
  reqUrl := resp.Request.URL
  baseURL := &url.URL {
    Scheme: reqUrl.Scheme,
    Host: reqUrl.Host,
  }
  base := baseURL.String()
  // fmt.Println(base)

  links, _ := link.Parse(resp.Body)

  var hrefs []string
  for _, l := range links {
    switch{
    // Case 1: Path no domain (/path) - collect
    case strings.HasPrefix(l.Href, "/"):
      hrefs = append(hrefs, base + l.Href)
    // Case 2: Path with domain (https://makeschool.com/path)- collect
    case strings.HasPrefix(l.Href, "http"):
      hrefs = append(hrefs, l.Href)
    // Case 3: Fragment (#fragment) - ignore
    // Case 4: Email (mailto:myemail@gmail.com) - ignore
    // Do not do anything for Cases 3 or 4. If desired, can add helper text later
    }
  }
  for _, href := range hrefs {
    fmt.Println(href)
  }

}
