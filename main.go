package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
)

func main() {
	// Set the proxy URL
	proxyURL, err := url.Parse("localhost:18443")
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return
	}

	// Create a new HTTP client with the proxy
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				addr = proxyURL.String() // Convert proxyURL to string
				dialer := net.Dialer{}
				return dialer.DialContext(ctx, network, addr)
			},
		},
	}

	// Create a new request
	req, err := http.NewRequest("GET", "https://globalcatalog.test.cloud.ibm.com/healthcheck", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Send the request through the proxy
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response body:", string(body))
}
