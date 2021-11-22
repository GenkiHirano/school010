package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// Student構造体の各フィールドを、構造体タグを使って、XORMのメタ情報を付与してください。

func main() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync2(new(Student))
	if err != nil {
		log.Fatal(err)
	}

	// 各関数を実行します
	// Insert(*engine)
	// // Get(*engine)
	// Update(*engine)
	// // Delete(*engine)
	// WhereAnd(*engine)
}

// 問題2
// Insert関数を実装してください。
func Insert(engine xorm.Engine) {

}

// 問題3
// Get関数を実装してください。
func Get(engine xorm.Engine) {

}

// 問題4
// Update関数を実装してください。
func Update(engine xorm.Engine) {

}

// 問題5
// Delete関数を実装してください。
func Delete(engine xorm.Engine) {

}

// 問題6
// WhereAnd関数を実装してください。
func WhereAnd(engine xorm.Engine) {

}
