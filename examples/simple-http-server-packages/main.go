package main

import (
	"fmt"
	"github.com/cbrgm/githubevents/examples/simple-http-server-packages/plugins"
	"github.com/cbrgm/githubevents/githubevents"
	"net/http"
)

func main() {
	handle := githubevents.New("")

	// return handleFuncs from other packages
	// and use them ad "plugins"
	handle.OnIssueCommentCreated(
		plugins.NewResponder("ping!"),
		plugins.NewResponder("pong!"),
	)

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
