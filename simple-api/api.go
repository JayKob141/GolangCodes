package main

import (
  "fmt"
  "log"
  "encoding/json"
  "strconv"
  "net/http"
  "database/sql"
  _ "github.com/lib/pq"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	EnterDate string `json:"enterDate"`
	LeaveDate string `json:"leaveDate"`
}

type ApiDb struct {
  db *sql.DB
}

func crearApiDbObject () ApiDb {
 connStr := "user=postgres password='example' dbname='postgres' host=localhost port=5432 sslmode=disable"
  db, err := sql.Open("postgres", connStr)
  apiDb := ApiDb{ db: db }

  if err != nil {
     fmt.Println(err.Error())
  } else {
     fmt.Println("Connected to database")
  }

  return apiDb
}

func writeUser(newUser User) User{
 apiDb := crearApiDbObject()
 db := apiDb.db
 defer db.Close()

 result, err := db.Exec("INSERT INTO users (name) values ($1)", newUser.Name )
 if err != nil {
	fmt.Println("Error executing query")
	fmt.Println(err.Error())
 } else {
   id, err2 := result.LastInsertId()
   if err2 != nil {
     fmt.Println("Cannot obtain LastInsertId")
     fmt.Println(err2.Error())
   }
   newUser.Id = strconv.FormatInt( id, 10)
 }

 return newUser
}

func readUsers() []User {
 apiDb := crearApiDbObject()
 db := apiDb.db
 defer db.Close()

 users := make([]User, 0)

 rows, errSelect := db.Query("SELECT * FROM users")
 if errSelect != nil {
	fmt.Println("Error executing select query")
	fmt.Println(errSelect.Error())
 }

 defer rows.Close()
 columns, _ := rows.Columns()


 fmt.Println("Columns of query:")
 fmt.Println(columns)

 for rows.Next() {
	 var user User
	 rows.Scan(&user.Id, &user.Name, &user.EnterDate, &user.LeaveDate)
	 fmt.Printf("%s %s %s %s\n", user.Id, user.Name, user.EnterDate, user.LeaveDate)
	 users = append(users, user)
 }

 return users
}

func listEntries(w http.ResponseWriter, r *http.Request) {
	users := readUsers()

	b, _ := json.Marshal(users)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func registerEnter(w http.ResponseWriter, r *http.Request) {
  newUser := User{ Name: "From golang" }
  newUser = writeUser(newUser)

	b, _ := json.Marshal(newUser)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

func registerLeave(w http.ResponseWriter, r *http.Request) {
	// user := User{ Id: "123", Name: "Jacob Meneses", EnterDate: "2018-01-01", LeaveDate: "2018-02-02" }

	// b, _ := json.Marshal(user)

	// w.Header().Set("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(b)
}

func main() {
	http.HandleFunc("/list-entries/", listEntries)
	http.HandleFunc("/register-enter/", registerEnter)
	http.HandleFunc("/register-leave/", registerLeave)

	fmt.Println("Running server on: http://localhost:8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
