package main

import (
	"fmt"
	"log"
	"net/http"
)

type movement struct {
	X_dif int
	Y_dif int
}

func main() {

	movements := make(chan movement, 500) //makes a channel for movements, with a depth of 500

	http.HandleFunc("/control/", control)

	http.HandleFunc("/move/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Recieved post request.")
		axis := req.FormValue("axis")
		step := req.FormValue("step")

		if axis != "x" || axis != "y" {
			w.WriteHeader(500)
		}

		log.Printf("Moving %d units on %s", step, axis)
		w.Write([]byte("Success!  "))
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
