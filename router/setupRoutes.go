package router

import (
	"mime"
	"net/http"
	"os"

	"github.com/ranon-rat/chatting-in-office/core"
)

func SetupRoutes() {

	mime.AddExtensionType(".js", "application/javascript")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/channel", func(w http.ResponseWriter, r *http.Request) {
		f, _ := os.ReadFile("static/channel.html")
		w.Write(f)
	})
	http.HandleFunc("/ws", core.JoiningChannel)
	http.ListenAndServe(":8080", nil)
}
