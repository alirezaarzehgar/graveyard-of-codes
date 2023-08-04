package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync/atomic"
)

var useAtomic = true

func add(num *int64, delta int64) {
	if useAtomic {
		atomic.AddInt64(num, delta)
		return
	}
	*num++
}

func main() {
	at := flag.Bool("atomic", true, "atomic or race")
	flag.Parse()
	useAtomic = *at

	fmt.Println(useAtomic)

	var count int64 = 0
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 50; i++ {
			add(&count, 1)
		}
		for i := 0; i < 50; i++ {
			add(&count, -1)
		}

		add(&count, 1)
	})

	http.HandleFunc("/getCount", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, count)
	})
	http.ListenAndServe(":8000", nil)
}
