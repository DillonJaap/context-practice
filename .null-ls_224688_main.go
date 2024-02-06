package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/")
}

func handler(w http.ResponseWriter, r *http.Request) {
}

func testOne(ctx context.Conetext) {
	time.Sleep()
}
