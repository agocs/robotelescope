package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/control/", handler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err.Error())
	}
}

func control(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Oh hi there!")
}
