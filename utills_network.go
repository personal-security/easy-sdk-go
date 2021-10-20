package easysdk

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

func GetRequestTimeout(url string, timeout int, token string) ([]byte, int, error) {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	//resp, err := client.Get(url)

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + token

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	if token != "" {
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
	}

	// Send req using http Client
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}

func GetRequest(url string) ([]byte, int, error) {
	return GetRequestTimeout(url, 2, "")
}

func GetRequestWithToken(url string, token string) ([]byte, int, error) {
	return GetRequestTimeout(url, 2, token)
}

func SendGetTorRequest(url string) ([]byte, error) {
	const PROXY_ADDR = torAddr

	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", PROXY_ADDR, nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		return nil, err
	}

	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	// create a request

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't create request:", err)
		return nil, err
	}
	// use the http client to fetch the page
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't GET page:", err)
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading body:", err)
		return nil, err
	}

	return b, nil
}

func SendPostTorRequest(url string, postData url.Values) ([]byte, error) {
	const PROXY_ADDR = torAddr

	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", PROXY_ADDR, nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		return nil, err
	}

	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	// create a request

	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(postData.Encode())) // URL-encoded payload
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(postData.Encode())))
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't create request:", err)
		return nil, err
	}
	// use the http client to fetch the page
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't POST page:", err)
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading body:", err)
		return nil, err
	}

	return b, err
}
