package main

import (
	"net/http"

	"github.com/Fybrid/go-api-architecture/internal/http/app/router"
	"github.com/Fybrid/go-api-architecture/pkg"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	pkg.LoadEnv()

	//DB接続
	// db, err := sql.Open("mysql", "root:fybrid@tcp(127.0.0.1:3306)/estell")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//ルーティング
	router.NewRouter()

	http.ListenAndServe(":8085", nil)
}
