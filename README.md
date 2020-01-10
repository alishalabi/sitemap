# Sitemap Builder
This project is designed to create a sitemap for any given website url. Sitemaps are vital for Search Engine Optimization, and are a requirement for popular search engines (such as Google).

This project makes use of a custom HTML anchor tag parser package (https://github.com/alishalabi/link_parser), which returns all anchor tags in a given section of HTML. A recursive Breadth First Search (bfs) is then called on each link to visit that link, returning unique urls that correspond to the base domain. Finally, the returned values are converted to a readable XML format, which can be written to file for sharing.

A sample sitemap for the website https://makeschool.com has been provided in the makeschool-sitemap.xml file of this project.

For more information on sitemaps, visit: https://en.wikipedia.org/wiki/Site_map

To submit your sitemap to improve your own SEO, submit the generated XML file to: https://search.google.com/search-console/welcome

## Installation
Project can be cloned locally using : `https://github.com/alishalabi/sitemap.git`

## Usage
- Once locally installed, open the `main.go` file
- In your `func main()`, update your `urlFlag` variable to your website, and update your `maxDepth` to your desired recursion depth
- Print your sitemap to your terminal by entering `$ go run main.go` in your terminal
- Write your sitemap to a shareable file by entering `$ go run main.go > my-sitemap.xml`

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Milestones
- [x] Be able to run a GET request on a website
- [x] Use our custom link parser to retrieve all anchor tags on that website (https://github.com/alishalabi/link_parser)
- [x] Normalize all links
- [x] Refactor code into modular functions
- [x] Filter links that do not direct to desired domain
- [x] Recursively search all links using a Breadth First Search
- [x] Return sitemap in sharable format (XML)

## Attribution
This project and its code structure were heavily influenced by Jon Calhoun's amazing tutorials on gophercises.com

## License
The MIT License (MIT)

Copyright (c) 2020 Ali Shalabi

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
