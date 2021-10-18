package main

import (
	"context"
	"fmt"
	"html/template"
	"log"

	"net/http"

	"github.com/go-session/session"
)

func home(page http.ResponseWriter, req *http.Request) {
	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}

	temp, err := template.ParseFiles("temp/html/home.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	temp.ExecuteTemplate(page, "home_page", nil)

	active, ok := store.Get("active_login")

	if ok {
	Active := fmt.Sprintf("User logged: %v", active)
	log.Println(Active)
	}

}