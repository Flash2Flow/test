package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"


	"net/http"


	"github.com/go-session/session"
	_ "github.com/go-sql-driver/mysql"
)

func auth(page http.ResponseWriter, req *http.Request) {

	login := req.FormValue("login")

	password := req.FormValue("password")


	store, err := session.Start(context.Background(), page, req)
	if err != nil {
		fmt.Fprint(page, err)
		return
	}
	temp, err := template.ParseFiles("temp/html/home.html", "static/js/script.js")

	if err != nil {
		fmt.Fprintf(page, err.Error())
	}

	temp.ExecuteTemplate(page, "home_page", nil)

	active, ok := store.Get("active_login")

	if ok {
		Active := fmt.Sprintf("User logged: %v", active)
		log.Println(Active)
		http.Redirect(page, req, "/", 302)
	}


		if login == ""{
			fmt.Fprintf(page, "Login cant be nil")
			http.Redirect(page, req, "/", 302)
		}else{
			if password == ""{
				fmt.Fprintf(page, "Password cant be nil")
				http.Redirect(page, req, "/", 302)
			}else{
				//auth
				db, err := sql.Open("mysql", "site:xLb43XEDkr8R4O@tcp(185.219.40.250)/site")


				if err != nil {
					panic(err.Error())
				}


				defer db.Close()

				query := fmt.Sprintf("SELECT * FROM `users` WHERE `login` = ?")
				rows, err := db.Query(query, login)
				if err != nil {
					if err == sql.ErrNoRows{
						fmt.Fprintf(page, "Такого пользователя не существует.")
						http.Redirect(page, req, "/", 302)
					}else{
						log.Println(err)
					}
				}
				defer rows.Close()

				for rows.Next() {
					var user UserFull
					err = rows.Scan(user.Id, user.Login, user.Password, user.Developer, user.Group, user.Ban, user.Undesirable)
					if err != nil {
						log.Println(err)
					}

					if password == user.Password {

						store.Set("active_login", login)
						err = store.Save()
						if err != nil {
							fmt.Fprint(page, err)
							return
						}
						http.Redirect(page, req, "/home/", 302)
						auth := fmt.Sprintf("User auth: %s", login)
						log.Println(auth)
					}
				}
			}
		}


}
