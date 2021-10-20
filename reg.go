package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math/rand"

	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-session/session"
	_ "github.com/go-sql-driver/mysql"
)

func reg(page http.ResponseWriter, req *http.Request) {
	login := req.FormValue("login")
	email := req.FormValue("email")
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
			if email == "" {
				fmt.Fprintf(page, "Email cant be nil")
				http.Redirect(page, req, "/", 302)
			}else{
				if password == ""{
					fmt.Fprintf(page, "Password cant be nil")
					http.Redirect(page, req, "/", 302)
				}else{
					//reg

					db, err := sql.Open("mysql", "site:xLb43XEDkr8R4O@tcp(185.219.40.250)/site")


					if err != nil {
						panic(err.Error())
					}


					defer db.Close()

					query := fmt.Sprintf("SELECT * FROM `users` WHERE `login` = ?")
					row, err := db.Query(query, login)
					if err != nil {
						if err == sql.ErrNoRows {
							query_email := fmt.Sprintf("SELECT * FROM `users` WHERE `email` = ?")
							_, err := db.Query(query_email, email)

							if err != nil {
								if err == sql.ErrNoRows{
									db, err := sql.Open("mysql", "site:xLb43XEDkr8R4O@tcp(185.219.40.250)/site")


									if err != nil {
										panic(err.Error())
									}

									md5_userkey := rand.Intn(9999999999)

									md5_password := GetMD5Hash(password)
									var null = "0"

									res, err := db.Exec("INSERT INTO `users` (`login`, `password`, `userkey`, `ban`, `group`, `developer`, `undesirable`) VALUES(?, ?, ?, ?, ?, ?, ?)", login, md5_password, md5_userkey, null, null, null, null)
									if err != nil {

									}
									log.Println("Create user login: " + login + " | password: "+ password + " | email :" + email)
									spew.Dump(res)

									store.Set("active_login", login)
									err = store.Save()
									if err != nil {
										fmt.Fprint(page, err)
										return
									}
									http.Redirect(page, req, "/home/", 302)
									auth := fmt.Sprintf("User auth: %s", login)
									log.Println(auth)


								}else{
									fmt.Fprintf(page, "Пользователь с такой почтой уже есть.")
									http.Redirect(page, req, "/", 302)
								}
							}
							spew.Dump(row)
							log.Println(err)
						}else{
							fmt.Fprintf(page, "Пользователь с таким логином уже есть.")
							http.Redirect(page, req, "/", 302)
						}

					}
					log.Println(err)
				}
			}
		}


}
