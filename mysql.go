package main // 声明 main 包
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", Index)

	db, error := sql.Open("mysql", "root:root@(172.16.150.202:3306)/godb")

	if error != nil {
		panic(error.Error())
	}

	defer db.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Blog:www.flysnow.org\nwechat:flysnow_org")
}
