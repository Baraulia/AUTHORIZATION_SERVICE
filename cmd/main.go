package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		logrus.Fatal("ListenAndServe: ", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello team!")
}
