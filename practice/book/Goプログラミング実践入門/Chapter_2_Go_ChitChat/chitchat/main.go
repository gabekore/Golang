package main

import (
	"net/http"
)


//**********************************************************
// 
// エントリポイント
// 
//**********************************************************
func main() {
	// マルチプレクサ生成
	//		ふたつ以上の入力をひとつの信号として出力する機構で
	mux := http.NewServeMux()
	
	//----------------------------------------------------------
	// ハンドラ作成
	// 「/static/」で始まる全てのリクエストURLについて、
	// 		1. URLから文字列「/static/」を取り去る
	// 		2. publicディレクトリを起点として、残った文字列を名前に持つファイルを探す
	// 
	// 例）
	//		リクエストURL
	//			http://localhost/static/css/bootstrap.min.css
	//		処理
	//			<アプリケーションのルート>/css/bootstrap.min.css
	//			↑このファイルを探す
	//----------------------------------------------------------
	// 起点となるディレクトリを作成
	// config.Staticはjsonファイル
	files := http.FileServer( http.Dir(config.Static) )
	// リクエストURLから「/static/」という文字列を取り去る
	// 		strip：剥く、剥ぐ
	mux.Handle( "/static/", http.StripPrefix("/static/", files) )

	// URLとハンドラ関数を紐付ける
	mux.HandleFunc("/"              , index)			// route_main.go
	mux.HandleFunc("err"            , err)				// route_main.go

	mux.HandleFunc("login"          , login)			// route_auth.go
	mux.HandleFunc("logout"         , logout)			// route_auth.go
	mux.HandleFunc("/signup"        , signup)			// route_auth.go
	mux.HandleFunc("/signup_account", signupAccount)	// route_auth.go
	mux.HandleFunc("/authenticate"  , authenticate)		// route_auth.go

	mux.HandleFunc("/thread/new"    , newThread)		// route_thread.go
	mux.HandleFunc("/thread/create" , createThread)		// route_thread.go
	mux.HandleFunc("/thread/post"   , postThread)		// route_thread.go
	mux.HandleFunc("thread/read"    , readThread)		// route_thread.go
	
	server := &http.Server {
		Addr    : "0.0.0.0:8080",
		Handler : mux,
	}
	
	server.ListenAndServe()
}


//**********************************************************
// 
// ハンドラ関数
// 
// HTMLを生成してResponseWriterに書き出す
// 
//**********************************************************
func index (w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html", "templates/navbar.html", "templates/index.html",}

	templates := template.Must(template.ParseFiles(files...))

	threads, err := data.Threads()
	
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}