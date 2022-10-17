package main

import(
  "net/http"
  "html/template"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  _ "fmt"
//   "log"
)

func main() {
    http.HandleFunc("/myamagata", myamagata)
    http.HandleFunc("/yichikawa", yichikawa)
    http.ListenAndServe(":8004", nil)
}

func myamagata(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "myamagata:myamagata_App_2022@tcp(10.35.0.59:3306)/seweb")
                if err != nil {
                        panic(err.Error())
                }
    defer db.Close()

     rows, err := db.Query("SELECT body FROM myamagata_web WHERE id=1")
                if err != nil {
                        panic(err.Error())
                }

    defer rows.Close()

    var body string
    rows.Next()
    rows.Scan(&body)
// log.Print(body)
    t , err := template.ParseFiles("myamagata.tpl")
                if err != nil {
                        panic(err.Error())
                }
    t.Execute(w, body)
}

func yichikawa(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "yichikawa:yichikawa_App_2022@tcp(10.35.0.59:3306)/seweb")
                if err != nil {
                        panic(err.Error())
                }
    defer db.Close()

     rows, err := db.Query("SELECT body FROM yichikawa_web WHERE id=1")
                if err != nil {
                        panic(err.Error())
                }

    defer rows.Close()

    var body string
    rows.Next()
    rows.Scan(&body)
// log.Print(body)
    t , err := template.ParseFiles("main.tpl")
                if err != nil {
                        panic(err.Error())
                }
    t.Execute(w, body)
}

