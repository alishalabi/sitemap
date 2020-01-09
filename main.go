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
  // Utilizing html anchor parsing package
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

func bfs(urlString string, maxDepth int) [] string {
  // Keep track of all visited webpages, using struct instead of bool bc less memory
  seenLinks := make(map[string]struct{})
  // Variable q = queue
  var q map[string]struct{}
  // Variable nq = next queue
  nq := map[string]struct{} {
    urlString: struct{}{},
  }
  for i := 0; i <= maxDepth; i++ {
    q, nq = nq, make(map[string]struct{})
    for url, _ := range q {
      // ok tests to see if a value is found for key in map
      // If page has already been seem, do nothing
      if _, ok := seenLinks[url]; ok {
        continue
      }
      seenLinks[url] = struct{}{}
      for _, link := range get(url) {
        nq[link] = struct{}{}
      }
    }
  }

  // Optimizing space for ret
  ret := make([]string, 0, len(seenLinks))
  for url, _ := range seenLinks{
    ret = append(ret, url)
  }
  return ret
}

func main() {
  urlFlag := flag.String("url", "https://makeschool.com", "Website for which you want to build a sitemap")

  maxDepth := flag.Int("depth", 3, "Maximum recursion depth when traversing links")
  flag.Parse()

  pages := bfs(*urlFlag, *maxDepth)

  // pages := get(*urlFlag)
  for _, page := range pages {
    fmt.Println(page)
  }

}
