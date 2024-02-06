package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	c := make(chan string)
	go oneSecPrint(ctx, 13, c)
	go twoSecPrint(ctx, 6, c)

	res := ""
	doBreak := false
	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			res += "context cancelled"
			doBreak = true
		case result := <-c:
			res += result
		}
		if doBreak {
			break
		}
	}
	w.Write([]byte(res))
}

func oneSecPrint(ctx context.Context, reps int, c chan string) {
	for i := 0; i < reps; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second * 1)
			fmt.Println("1 second delay")
		}
	}
	c <- "done one sec"
}

func twoSecPrint(ctx context.Context, reps int, c chan string) {
	for i := 0; i < reps; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("2 second delay")
		}
	}
	c <- "done two sec"
}
