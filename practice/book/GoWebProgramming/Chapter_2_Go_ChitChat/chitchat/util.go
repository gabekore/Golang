package main

//**********************************************************
// Checks if the user is logged in and has a session, 
// if not err is not nil
//**********************************************************
func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	
	// リクエストからクッキーを取得
	cookie, err := r.Cookie("_cookie")
	// ユーザーがログインしているならクッキーがあるはず
	if err == nil {
		// データベースを検索
		// ユニークIDが存在してるか？
		sess = data.Session{ Uuid: cookie.Value }
		
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}



//**********************************************************
// HTML作成
//**********************************************************
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {

	var files []string
	
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	
	templates := template.Must(template.ParseFiles(files...))
	
	// テンプレートのlayout.htmlを実行する
	templates.ExecuteTemplate(w, "layout", data)

}