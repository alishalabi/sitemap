package main

import (
  "flag"
  "fmt"
  "net/http"
  "net/url"
  "strings"
  "io"
  // "os"

  "github.com/alishalabi/link_parser"
)

func hrefs(r io.Reader, base string) []string {
  links, _ := link.Parse(r)
  // Variable ret = "return value"
  var ret []string
  for _, l := range links {
    switch{
    // Case 1: Path no domain (/path) - collect
    case strings.HasPrefix(l.Href, "/"):
      ret = append(ret, base + l.Href)
    // Case 2: Path with domain (https://makeschool.com/path)- collect
    case strings.HasPrefix(l.Href, "http"):
      ret = append(ret, l.Href)
    // Case 3: Fragment (#fragment) - ignore
    // Case 4: Email (mailto:myemail@gmail.com) - ignore
    // Do not do anything for Cases 3 or 4. If desired, can add helper text later
    }
  }
  return ret
}

// Initiate GET request
func get(urlString string) []string {
  resp, err := http.Get(urlString)
  if err != nil {
    panic(err)
  }
  // Must ALWAYS close response body, or memory leak will occur
  defer resp.Body.Close() // defer will run when function ends

  // Link cleanup
  reqUrl := resp.Request.URL
  baseURL := &url.URL {
    Scheme: reqUrl.Scheme,
    Host: reqUrl.Host,
  }
  base := baseURL.String()

  return filter(hrefs(resp.Body, base), goodPrefix(base))
}

func filter(links []string, keepFn func(string) bool) []string {
  // Variable ret = "return value"
  var ret []string
  for _, link := range links {
    if keepFn(link) {
      ret = append(ret, link)
    }
  }
  return ret
}

func goodPrefix(pfx string) func(string) bool {
  return func(link string) bool {
    // Variable pfx = "prefix value"
    return strings.HasPrefix(link, pfx)
  }
}

func main() {
  urlFlag := flag.String("url", "https://makeschool.com", "Website for which you want to build a sitemap")
  flag.Parse()

  pages := get(*urlFlag)
  for _, page := range pages {
    fmt.Println(page)
  }

}
