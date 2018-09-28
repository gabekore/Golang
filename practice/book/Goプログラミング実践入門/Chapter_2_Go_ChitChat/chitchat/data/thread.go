package data

import(
	"time"
)

// threadsテーブルと紐付く
type Thread struct {
	Id			int
	Uuid		string
	Topic		string
	UserId		int
	CreatedAt	time.Time
}


func Threads() (threads []Thread, err error) {

	// SQL文作成
	sql := " SELECT "
		 + "   id "
		 + "   , uuid "
		 + "   , topi "
		 + "   , user_id "
		 + "   , created_at "
		 + " FROM "
		 + "   threads "
		 + " ORDER BY "
		 + "   created_at DESC "

	// SQL実行
	rows, err := Db.Query(sql)

	// 取得結果が0件なら処理終了
	if err != nill {
		return
	}
	
	// 取得結果が1件以上あるなら処理する
	for rows.Next() {
		
		// 一時的に構造体を作成
		th := Thread{}
		
		// 一時的な構造体にデータを入れる
		err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt)
		if  err != nil {
			return
		}

		// 一時的な構造体を引数のthreads構造体にappendする
		threads = append(threads, th)
		
	}

	rows.Close()
	
	return

}

func (thread *Thread) NumReplies() (count int) {

	// SQL文作成
	sql := " SELECT "
		 + "   count(*) "
		 + " FROM "
		 + "   posts "
		 + " WHERE "
		 + "   thread_id = $1 "

	// SQL実行
	rows, err := Db.Query(sql)

	// 取得結果が0件なら処理終了
	if err != nill {
		return
	}
	
	for rows.Next() {
		// 一時的な構造体にデータを入れる
		err = rows.Scan(&count)
		if  err != nil {
			return
		}
	}
	
	rows.Close()
	
	return
}