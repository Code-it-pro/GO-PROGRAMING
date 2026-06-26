package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Starting with Urls")
	mylink := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26"

	// fmt.Printf("Url datatype: %T", mylink)
	myurl, error := url.Parse(mylink)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Url : ", myurl)
	// fmt.Printf("Url datatype: %T", myurl)

	fmt.Println("Scheme:", myurl.Scheme)
	fmt.Println("Scheme:", myurl.Host)
	fmt.Println("Path:", myurl.Path)
	fmt.Println("Query:", myurl.RawQuery)

	myurl.Host = "gagandeep.com"
	finalUrl := myurl.String()
	fmt.Println("New Url:", finalUrl)
	fmt.Println("Url : ", myurl)

}
