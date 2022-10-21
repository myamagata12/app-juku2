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
    http.HandleFunc("/tozono", tozono)
    http.ListenAndServe(":8005", nil)
    
    http.HandleFunc("/myamagata", myamagata)
    http.ListenAndServe(":8003", nil)

}

func tozono(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "tozono:tozono_App_2022@tcp(10.35.0.59:3306)/seweb")
                if err != nil {
                        panic(err.Error())
                }
    defer db.Close()

     rows, err := db.Query("SELECT body FROM tozono_web WHERE id=1")
                if err != nil {
                        panic(err.Error())
                }

    defer rows.Close()

    var body string
    rows.Next()
    rows.Scan(&body)
// log.Print(body)
    t , err := template.ParseFiles("tozono.tpl")
                if err != nil {
                        panic(err.Error())
                }
    t.Execute(w, body)
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
