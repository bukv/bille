package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CreateHttpServer() {
	req := mux.NewRouter()
	req.HandleFunc("/", ProcessSendWebPage)
	req.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("styles/"))))
	req.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))

	http.Handle("/", req) // http://localhost:8090/
	http.ListenAndServe(":8090", nil)
}
