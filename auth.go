package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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
		http.Redirect(page, req, "/home/", 302)
	}else{
		http.Redirect(page, req, "/", 302)
	}

	if login == "" {
	fmt.Fprintf(page, "Login cant be nil")
	}else{
		if password == "" {
			fmt.Fprintf(page, "Password cant be nil")
		}else{
			md5_password := GetMD5Hash(password)
			users := fmt.Sprintf("/users/user_%s.json", login)
			dat, err := ioutil.ReadFile(users)
			if err != nil {
				fmt.Fprintf(page, "Ошибка авторизации, такого пользователя не существует.")
			}
			user := UserFull{}
			err = json.Unmarshal(dat, &user)
			if err != nil {
				return
			}
			if user.Password == md5_password {
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
