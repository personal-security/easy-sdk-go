package easysdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/proxy"
)

const torAddr = "127.0.0.1:9050"

func MattermostSendMessage(url, hook, message string, isCode bool, warning bool) (string, error) {
	const PROXY_ADDR = torAddr
	var URL = "http://" + url + "/hooks/" + hook

	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", PROXY_ADDR, nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		return "", err
	}

	type SendPayload struct {
		Text string `json:"text"`
	}

	if isCode {
		message = "```\n" + message + "\n```"
	}

	if warning {
		message = "<!all> \n" + message
	}

	sendPayload := &SendPayload{Text: message}
	jsonToSend, err := json.Marshal(sendPayload)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var jsonStr = []byte(jsonToSend)

	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	// create a request

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	//req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't create request:", err)
		return "", err
	}
	// use the http client to fetch the page
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't GET page:", err)
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading body:", err)
		return "", err
	}

	return string(b), nil
}
