package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func errorHandle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func testHeader(headerKey string, headerValue string, debug string, except string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:"+os.Args[1]+"/?debug="+debug, nil)
	errorHandle(err)
	req.Header.Set(headerKey, headerValue)
	resp, err := client.Do(req)
	errorHandle(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errorHandle(err)
	if string(body) != except {
		log.Fatalln(except)
	} else {
		log.Printf("OK %s %s %s %s", debug, except, headerKey, headerValue)
	}
}

func testUserAgent(UserAgent string, debug string, except string) {
	testHeader("User-Agent", UserAgent, debug, except)
}

func testXForwardedFor(XForwardedFor string, debug string, except string) {
	testHeader("X-Forwarded-For", XForwardedFor, debug, except)
}

func main() {
	testUserAgent("mobile", "mobile", "true")
	testUserAgent("desktop", "mobile", "false")

	testUserAgent("MicroMessenger", "wechat", "true")
	testUserAgent("Line", "wechat", "false")

	testUserAgent("Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0; Xbox; Xbox One)", "platform", "Windows")
	testUserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko", "platform", "Windows")
	testUserAgent("Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Mobile Safari/535.19", "platform", "Android")
	testUserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2693.2 Safari/537.36", "platform", "Linux")
	testUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.114 Safari/537.36", "platform", "Mac")
	testUserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1", "platform", "iPhone")

	testXForwardedFor("", "ip", "127.0.0.1")
	testXForwardedFor("8.8.8.8", "ip", "8.8.8.8")
	testXForwardedFor("8.8.8.8, 114.114.114.114", "ip", "114.114.114.114")
}
