package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type movement struct {
	X_dif int
	Y_dif int
}

func main() {

	//movements := make(chan movement, 500) //makes a channel for movements, with a depth of 500

	http.HandleFunc("/control/", control)

	http.HandleFunc("/move/", func(w http.ResponseWriter, req *http.Request) {

		axis := req.FormValue("axis")
		step := req.FormValue("step")

		log.Printf("Recieved post request. %s, %s", axis, step)
		w.Header().Set("Access-Control-Allow-Origin", "*")

		message := []byte("Don't know yet")

		toMove, err := strconv.Atoi(step)
		if err != nil {
			w.WriteHeader(500)
			message = []byte("Step needs to be an integer")
		} else if axis != "x" && axis != "y" {
			w.WriteHeader(500)
			message = []byte("Axis needs to be x or y")
		} else {
			log.Printf("Moving %d units on %s", toMove, axis)
			message = []byte("Success! ")
		}
		w.Write(message)
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
