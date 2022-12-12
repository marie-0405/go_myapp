package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("The number of bites: %d", n))
	})

	http.ListenAndServe(":8080", nil)
}
