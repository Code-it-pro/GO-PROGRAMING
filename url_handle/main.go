package main

import (
	"fmt"
	"net/url"
)

// main demonstrates how to parse and manipulate URLs in Go
// It shows how to extract URL components and modify them
func main() {
	fmt.Println("Starting with Urls")

	// Define a sample URL string with query parameters
	mylink := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26"

	// Parse the URL string into a *url.URL struct for easier manipulation
	myurl, error := url.Parse(mylink)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Url : ", myurl)

	// Extract and display individual URL components
	fmt.Println("Scheme:", myurl.Scheme)  // Protocol (https)
	fmt.Println("Host:", myurl.Host)      // Domain name (www.youtube.com)
	fmt.Println("Path:", myurl.Path)      // URL path (/watch)
	fmt.Println("Query:", myurl.RawQuery) // Query parameters (v=..., list=..., index=...)

	// Modify the Host component of the URL
	myurl.Host = "gagandeep.com"

	// Convert the modified URL back to a string
	finalUrl := myurl.String()
	fmt.Println("New Url:", finalUrl)
	fmt.Println("Url : ", myurl)
}

/*
EXPLANATION:
============

This program demonstrates URL parsing and manipulation in Go:

1. URL Parsing: url.Parse() breaks down a URL string into its components
   - Scheme: The protocol (https, http, ftp, etc.)
   - Host: The domain name (www.youtube.com)
   - Path: The resource path (/watch)
   - RawQuery: Query parameters (?v=..., list=..., index=...)

2. URL Components: Individual parts can be accessed via the *url.URL struct
   - Original URL: https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=...&index=26
   - Scheme: https
   - Host: www.youtube.com
   - Path: /watch
   - RawQuery: v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26

3. URL Manipulation: You can modify URL components and rebuild the URL
   - Changed Host from "www.youtube.com" to "gagandeep.com"
   - Other components (scheme, path, query) remain unchanged
   - url.String() reconstructs the complete URL from modified components

4. Use Cases:
   - Parsing API endpoints
   - Extracting URL parameters
   - Building dynamic URLs
   - URL validation and manipulation
   - Redirecting or rewriting URLs
*/
