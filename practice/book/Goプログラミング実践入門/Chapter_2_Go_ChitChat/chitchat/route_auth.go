package main


import (
	"github.com/gabekore/Golang/tree/master/practice/book/Go%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E5%AE%9F%E8%B7%B5%E5%85%A5%E9%96%80/Chapter_2_Go_ChitChat/chitchat/data"
	"net/http"
)


//**********************************************************
// GET /login
// Show the login page
//**********************************************************
func login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}

//**********************************************************
// GET /signup
// Show the signup page
//**********************************************************
func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "login.layout", "public.navbar", "signup")
}

//**********************************************************
// POST /signupAccount
// Create the user account
//**********************************************************
func signupAccount(w http.ResponseWriter, r *http.Request) {
	err := request.ParseForm()
	
	if err != nil {
		danger(err, "Cannot parse form")
	}
	
	user := data.User {
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	
	// ステータスコード302 Found (※Moved Temporarily)
	// リクエストで要求された情報の中身が別の場所に移動したことを示します。
	// ※その移動は一時的なもの（一時的でないなら301）
	http.Redirect(w, r, "/login", 302)
}

//**********************************************************
// POST /authenticate
// Authenticate the user given the email and password
//**********************************************************
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	
	// data.UserByEmail()は、メアドを与えられると構造体Userを返す
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	
	// data.Encrypt()は、文字列を暗号化する
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		
		// 構造体Sessionを作成
		// type Session struct {
		//     Id			int
		//     Uuid			string
		//     Email		string
		//     UserId		int
		//     CreatedAt	time.time
		// }
		// ※Uuidはランダムに生成されたユニークなID
		//   このUuidをブラウザに保存したい
		sesion, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		
		// 構造体Cookieを作る
		// 有効期限を設定していないのでブラウザ終了時に自動で消去される
		// HttpOnlyは、クッキーへのアクセスをHTTPとHTTPSのみ許可する
		// ※JavaScript等の非HTTPのAPIでは禁止となる
		cookie := http.Cookie {
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		
		// レスポンスにクッキーを追加
		http.SetCookie(w, &cookie)
		
		http.Redirect(w, r, "/", 302)
	} else {
		// ステータスコード302 Found (※Moved Temporarily)
		// リクエストで要求された情報の中身が別の場所に移動したことを示します。
		// ※その移動は一時的なもの（一時的でないなら301）
		http.Redirect(w, r, "/login", 302)
	}
}

//**********************************************************
// GET /logout
// Logs the user out
//**********************************************************
func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := request.Cookie("_cookie")
	
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := data.Session{ Uuid: cookie.Value }
		session.DeleteByUUID()
	}

	// ステータスコード302 Found (※Moved Temporarily)
	// リクエストで要求された情報の中身が別の場所に移動したことを示します。
	// ※その移動は一時的なもの（一時的でないなら301）
	http.Redirect(w, r, "/", 302)
}

