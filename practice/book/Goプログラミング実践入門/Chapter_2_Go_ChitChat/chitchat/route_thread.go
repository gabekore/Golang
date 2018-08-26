package main


import (
	"fmt"
	"github.com/gabekore/Golang/tree/master/practice/book/Go%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E5%AE%9F%E8%B7%B5%E5%85%A5%E9%96%80/Chapter_2_Go_ChitChat/chitchat/data"
	"net/http"
)


//**********************************************************
// GET /threads/new
// Show the new thread form page
//**********************************************************
func newThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}





//**********************************************************
// GET /threads/create
// Create the user account
//**********************************************************
func createThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	
	if err != nill {
		http.Redirect(w, r, "login", 302)
	} else {
		err = r.ParseForm()
		if err != nill {
			danger(err, "Cannot parse form")
		}
		
		user, err := sess.User()
		if err != nill {
			danger(err, "Cannot get user from session")
		}
		
		topic := r.PostFormValue("topic")
		if _, err := user.createThread(topic);  err != nill {
			danger(err, "Cannot create thread")
		}
		
		http.Redirect(w, r, "/", 302)
	}
}




//**********************************************************
// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
//**********************************************************
func readThread(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	
	if err != nil {
		error_message(w, r, "Cannot read thread")
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, &thread, "layout", "public.navbar", "public.thread")
		} else {
			generateHTML(w, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}





//**********************************************************
// POST /thread/post
// Create the post
//**********************************************************
func postThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		
		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			err_message(w, r, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?=id=", uuid)
		
		http.Redirect(w, r, url, 302)
	}
}

