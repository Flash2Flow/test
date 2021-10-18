package main

import (
	"context"
	"fmt"
	"github.com/go-session/session"
	"log"
	"net/http"
	"os"
	"html/template"

)


func main() {
	//design pages
	http.HandleFunc("/", home)

	//tech pages
	http.HandleFunc("/auth", auth)

	port := os.Getenv("PORT")
	log.Println("start server with port: " +port)
	http.ListenAndServe(":"+port, nil)
}



func home(page http.ResponseWriter, req *http.Request) {
	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}
	temp, err := template.ParseFiles("index.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	temp.ExecuteTemplate(page, "home_page", nil)

	active, ok := store.Get("active_login")
	if ok {
		fmt.Fprintf(page, "Nice session %s", active)
		fmt.Fprintf(page, "Gay")
	}else{
		fmt.Fprintf(page, "Error session" )
	}


}