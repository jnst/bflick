package main

import "testing"

func TestToUrls(t *testing.T) {
	expect := []string{"http://localhost:8080/ja-jp/ex/price%v"}
	names := []string{"Eth", "Bch", "Mona"}
	for _, v := range ToUrls(names) {

	}
}
