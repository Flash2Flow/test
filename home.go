package main

import (
	"context"
	"html/template"

	"fmt"

	"net/http"

	"github.com/go-session/session"
)

func home(page http.ResponseWriter, req *http.Request) {
	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}

	active, ok := store.Get("active_login")

	if ok {
		temp, err := template.ParseFiles("temp/html/HomeBadAuth.html")

		if err != nil {
			fmt.Fprintf(page, err.Error())
		}

		temp.ExecuteTemplate(page, "home_page_nice", active)
	}else{
		temp, err := template.ParseFiles("temp/html/HomeNiceAuth.html")

		if err != nil {
			fmt.Fprintf(page, err.Error())
		}

		temp.ExecuteTemplate(page, "home_page_bad", nil)
	}

}
