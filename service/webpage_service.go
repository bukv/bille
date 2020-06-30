package service

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func ProcessSendWebPage(w http.ResponseWriter, req *http.Request) {
	var angle float64
	err := req.ParseForm() // Parse power an angle from form
	if err != nil {
		log.Println(err)
	}

	// Getting values from form
	for key, values := range req.Form {
		for _, value := range values {
			switch key {
			case "direction":
				angle, err = strconv.ParseFloat(value, 64)
				if err != nil {
					log.Println(err)
				}
				break
			case "force":
				power, err = strconv.Atoi(value)
				if err != nil {
					log.Println(err)
				}
				break
			case "update":
				break
			default:
				log.Println("invalid request")
				break
			}
		}
	}

	ImageСreationProcess(angle)
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, nil)
}
