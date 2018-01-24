package main

import (
	"fmt"
	"net/http"
)

func SuccessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "200 OK")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "404 Not Found")
}

// Start bitFlyer mock server
func main() {
	fmt.Println("Server started: http://localhost:8080")
	http.HandleFunc("/ja-jp/ex/price", SuccessHandler)
	http.HandleFunc("/ja-jp/ex/EthPrice", SuccessHandler)
	http.HandleFunc("/ja-jp/ex/EtcPrice", SuccessHandler)
	http.HandleFunc("/ja-jp/ex/LtcPrice", SuccessHandler)
	http.HandleFunc("/ja-jp/ex/BchPrice", SuccessHandler)
	http.HandleFunc("/ja-jp/ex/MonaPrice", SuccessHandler)
	http.ListenAndServe(":8080", nil)
}
