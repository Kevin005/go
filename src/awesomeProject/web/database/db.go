package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"net/http"
	"io"
)


func helloHandler(w http.ResponseWriter, req *http.Request) {
	insertDb()
	io.WriteString(w, "hello acd")
}

func main() {

	// 通过 HandlerFunc 把函数转换成 Handler 接口的实现对象
	hh := http.HandlerFunc(helloHandler)
	http.Handle("/xiaoshi", hh)
	http.ListenAndServe(":8080", nil)
}

func insertDb(){
	db, err := sql.Open("mysql", "root:123456@/baidu?charset=utf8")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT t_feedback SET email=?,content=?")
	checkErr(err)

	res, err := stmt.Exec("研发部门", "2016-03-06")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
