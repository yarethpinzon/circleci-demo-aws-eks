package main

import (
	"fmt"
	"net/http"
	"os"
)

func mainHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "FUNCIONA YSP IN THE HOUSE MATHER FUCKERS! (Version info: %s, build date: %s)", os.Getenv("VERSION_INFO"), os.Getenv("BUILD_DATE"))
	})
}

func main() {
	http.HandleFunc("/", mainHandler())
	http.ListenAndServe(":8080", nil)
}
