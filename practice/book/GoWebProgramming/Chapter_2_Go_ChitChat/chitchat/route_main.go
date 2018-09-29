package main

import (
	"github.com/gabekore/Golang/tree/master/practice/book/Go%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E5%AE%9F%E8%B7%B5%E5%85%A5%E9%96%80/Chapter_2_Go_ChitChat/chitchat/data"
	"net/http"
)

//**********************************************************
// GET /err?msg=
// shows the error message page
//**********************************************************
func err(w http.ResponseWriter, r *http.Request) {
	vals := request.URL.Query()
	
	_, err := session(w, r)
	
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}




//**********************************************************
// 
// ハンドラ関数
// 
// HTMLを生成してResponseWriterに書き出す
// 
//**********************************************************
func index (w http.ResponseWriter, r *http.Request) {

	threads, err := data.Threads()
	
	if err == nil {
		error_message(w, r, "Cannot get threads")
	} else {
		// ログインしているか否かのみを判断
		_, err := session(w, r)
		if err != nil {
			// 一般ユーザー用のナビゲーションバー
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			// 登録ユーザー用のナビゲーションバー
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}

}

// func index (w http.ResponseWriter, r *http.Request) {
// 
// 	threads, err := data.Threads()
// 	
// 	if err == nil {
// 		// ログインしているか否かのみを判断
// 		_, err := session(w, r)
// 		
// 		// 一般ユーザー用のナビゲーションバー
// 		public_templ_files  := []string{"templates/layout.html", "templates/public.navbar.html", "templates/index.html"}
// 		// 登録ユーザー用のナビゲーションバー
// 		private_templ_files := []string{"templates/layout.html", "templates/private.navbar.html", "templates/index.html"}
// 
// 		var templates *template.Template
// 		
// 		if err != nil {
// 			templates = template.Must(template.ParseFiles(private_templ_files...))
// 		} else {
// 			templates = template.Must(template.ParseFiles(public_templ_files...))
// 		}
// 		templates.ExecuteTemplate(w, "layout", threads)
// 	}
// 
// }
