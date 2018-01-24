package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

func ToUrls(names []string) []string {
	var urls []string
	for _, name := range names {
		urls = append(urls, fmt.Sprintf("http://localhost:8080/ja-jp/ex/price%v", name))
	}
	return urls
}

func main() {
	url := "http://localhost:8080/ja-jp/ex/priceLsk"

	maxConnection := make(chan bool, 200)
	wg := &sync.WaitGroup{}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		println(resp.Status)
	} else {
		println(resp.Status)
	}
}
