package main

import (
	"fmt"
	"github.com/cbrgm/githubevents/githubevents"
	"net/http"
)

func main() {
	handle := githubevents.New("")

	// pass the eventHandler to funcs that define callbacks
	newPing(handle)
	newPong(handle)

	http.HandleFunc("/hook", func(w http.ResponseWriter, r *http.Request) {
		err := handle.HandleEventRequest(r)
		if err != nil {
			fmt.Println("error")
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func newPing(handle *githubevents.EventHandler) {
	handle.OnBeforeAny(
		func(deliveryID string, eventName string, event any) error {
			fmt.Println("ping!")
			return nil
		},
	)
}

func newPong(handle *githubevents.EventHandler) {
	handle.OnBeforeAny(
		func(deliveryID string, eventName string, event any) error {
			fmt.Println("pong!")
			return nil
		},
	)
}
