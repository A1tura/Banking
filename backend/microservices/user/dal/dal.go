package dal

import (
	"db"
	"fmt"
)

func EmailInUse(db *db.DB, email string) (bool, bool) {
	var emailInUse bool
	row := db.QueryRow(`SELECT CASE
           WHEN EXISTS (SELECT 1 FROM users WHERE email=$1) THEN true
           ELSE false
       END;`, email)

	if err := row.Scan(&emailInUse); err != nil {
		fmt.Println(err)
		return false, false
	}

	return true, emailInUse
}

func UsernameInUse(db *db.DB, email string) (bool, bool) {
	var usernameInUse bool
	row := db.QueryRow(`SELECT CASE
           WHEN EXISTS (SELECT 1 FROM users WHERE username=$1) THEN true
           ELSE false
       END;`, email)

	if err := row.Scan(&usernameInUse); err != nil {
		fmt.Println(err)
		return false, false
	}

	return true, usernameInUse
}

func CreateUser(db *db.DB, username, password_hash, email string) bool {
	_, err := db.Exec(`INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`, username, email, password_hash)

	if err != nil {
		return false
	}

	return true
}
