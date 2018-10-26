package server

import (
	"fmt"
	"log"
	"net/http"
	"video"
)

// Server : An http server
type Server struct {
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not found")
}

// Init : Init http server
func Init() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/upload/", uploadHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8010", nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/Upload/"):]
	video1 := &video.Video{
		Title: title,
		Media: []byte("This is a sample Video."),
	}
	error := video1.Upload()
	if error == nil {
		fmt.Fprintf(w, "<h1>%s</h1>", "success")
	} else {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Oops!", error)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, error := video.RequestVideo(title)
	if error == nil {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Media)
	} else {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Oops!", "Not Found")
	}
}
