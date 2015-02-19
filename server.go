package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type movement struct {
	X_dif int
	Y_dif int
}

func main() {

	movements := make(chan movement, 500) //makes a channel for movements, with a depth of 500

	http.HandleFunc("/control/", control)
	http.HandleFunc("/move_x/{toMove}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		toMove, err := strconv.Atoi(vars["toMove"])
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Error: %s is probably not an int", vars["toMove"])
		} else {
			move_x(movements, toMove)
			fmt.Fprintf(w, "Moved %d on x", toMove)
		}
	})

	http.HandleFunc("/move_y/{toMove}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		toMove, err := strconv.Atoi(vars["toMove"])
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Error: %s is probably not an int", vars["toMove"])
		} else {
			move_y(movements, toMove)
			fmt.Fprintf(w, "Moved %d on y", toMove)
		}
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err.Error())
	}
}

func control(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Oh hi there!")
}

func move_x(movements chan movement, toMove int) {

}

func move_y(movements chan movement, toMove int) {

}
