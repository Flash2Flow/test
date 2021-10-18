package main

import (
	"context"

	"html/template"

	"log"

	"fmt"

	"net/http"

	"github.com/go-session/session"
)

func home(page http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("temp/html/home.html")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	temp.ExecuteTemplate(page, "home_page", nil)

	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}


	active, ok := store.Get("active_login")

	if ok {
		ActiveLog := fmt.Sprintf("User active: %v", active)
		log.Println(ActiveLog)
		return
	}

}