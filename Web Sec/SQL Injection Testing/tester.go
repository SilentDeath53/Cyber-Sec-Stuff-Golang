package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// DB connection parameters
const (
	dbDriver   = "mysql"
	dbUser     = "your_username"
	dbPassword = "your_password"
	dbName     = "your_database"
)

// User struct represents a user entity
type User struct {
	ID       int
	Username string
	Password string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, password)

		db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName))
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Username, &user.Password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		if len(users) > 0 {
			fmt.Fprintln(w, "Login successful!")
		} else {
			fmt.Fprintln(w, "Login failed!")
		}
	} else {
		fmt.Fprintf(w, `<html>
			<body>
				<form action="/login" method="post">
					<input type="text" name="username" placeholder="Username"><br>
					<input type="password" name="password" placeholder="Password"><br>
					<input type="submit" value="Login">
				</form>
			</body>
		</html>`)
	}
}

func main() {
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
