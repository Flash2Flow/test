package main

import (
	"fmt"
	"net/http"
	"context"
	"github.com/go-session/session"
)

func auth(page http.ResponseWriter, req *http.Request) {
	login := req.FormValue("login")
	password := req.FormValue("password")

	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}

	if login == "test"{
		if password == "1234"{
			store.Set("active_login", login)
			err = store.Save()
			if err != nil {
				fmt.Fprint(page, err)
				return
			}
			http.Redirect(page, req, "/", 302)
		}else{
			fmt.Fprintf(page, "Err password")
			http.Redirect(page, req, "/", 302)
		}
	}else{
		fmt.Fprintf(page, "Err login")
		http.Redirect(page, req, "/", 302)
	}
}