package router

import (
	"net/http"
	"os"

	"github.com/ranon-rat/chatting-in-office/core"
)

func SetupRoutes() {

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		f, _ := os.ReadFile("static/message.html")
		w.Write(f)
	})
	http.HandleFunc("/ws", core.JoiningChannel)
	http.ListenAndServe(":8080", nil)
}
