package main

import (
	"fmt"
	"net/url"
)

func parseUrl(site string) (*url.URL, error) {
	url, err := url.Parse(site)

	return url, err
}

func main() {
	parsed, err := parseUrl("http://a b.com/")


	if err != nil {
		if urlError, ok := err.(*url.Error); ok {
			fmt.Printf("The op is %v and the url is %v\n", urlError.Op, urlError.URL)
		}
	}
	fmt.Printf("The url is %v\n", parsed)
}