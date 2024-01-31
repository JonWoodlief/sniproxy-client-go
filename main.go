package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

const proxyAddr = "localhost:18443"

func main() {
	// Create a new HTTP client with the proxy
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				dialer := net.Dialer{}
				return dialer.DialContext(ctx, network, proxyAddr)
			},
		},
	}

	// Create a new request
	req, err := http.NewRequest("GET", "https://ibm.com/healthcheck", nil)
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
