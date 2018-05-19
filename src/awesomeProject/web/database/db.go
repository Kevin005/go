package main

import (
	_ "github.com/go-sql-driver/mysql" //https://github.com/go-sql-driver/mysql
	"database/sql"
	"net/http"
	"io/ioutil"
	"io"
	"encoding/json"
)

func main() {

	// 通过 HandlerFunc 把函数转换成 Handler 接口的实现对象
	hh := http.HandlerFunc(helloHandler)
	http.Handle("/xiaoshi", hh)
	http.ListenAndServe(":9090", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var feedback Feedback
	//获取到body传过来的值
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &feedback); err != nil {
		w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	insertDb(feedback)
	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		panic(err)
	}
}

type Feedback struct {
	Id      int    `json:"id"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

func insertDb(t Feedback) int64 {
	db, err := sql.Open("mysql", "root:12345678@/user?charset=utf8")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT t_feedback SET email=? ,content=?")
	checkErr(err)
	res, err := stmt.Exec(t.Email, t.Content)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	db.Close()
	return id
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
