package main

import (
	"encoding/json"
	"net/http"
)

type greeting struct {
	Hello string
}

func main() {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {

		greet := greeting{
			Hello: "world",
		}
		output, err := json.Marshal(greet)
		if err != nil {
			return
		}
		w.Write(output)

	})

	http.ListenAndServe(":8080", nil)

}