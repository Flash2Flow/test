package main

import (
	"context"
	"fmt"
	"html/template"
	"log"

	"net/http"

	"github.com/go-session/session"
)

func home_active(page http.ResponseWriter, req *http.Request) {
	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}
	temp, err := template.ParseFiles("temp/html/home_active.html", "static/js/script.js")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	temp.ExecuteTemplate(page, "home_page_active", nil)

	active, ok := store.Get("active_login")

	if ok {
		Active := fmt.Sprintf("User logged: %v", active)
		log.Println(Active)
		temp.ExecuteTemplate(page, "home_page", active)
	}else{
		http.Redirect(page, req, "/", 302)
	}

}