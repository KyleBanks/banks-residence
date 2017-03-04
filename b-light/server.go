package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	gpioPinNumber = 4
	httpPort      = ":8080"
)

func main() {
	l, err := newLight(gpioPinNumber)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	http.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		l.toggleState()

		out := "ON"
		if l.State != On {
			out = "OFF"
		}
		fmt.Fprintf(w, out)
	})

	log.Printf("Listening on Port %v", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
