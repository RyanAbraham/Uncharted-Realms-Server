package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID          string
	IDBin       int
	Email       string
	MatchesWon  int
	MatchesLost int
}

func (u *User) GetUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id_text, email FROM user WHERE id_bin=%d", u.IDBin)
	return db.QueryRow(statement).Scan(&u.ID, &u.Email)
}

func (u *User) UpdateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE user SET matches_won='%d', matches_lost = %d WHERE id_bin=%d", u.MatchesWon, u.MatchesLost, u.IDBin)
	_, err := db.Exec(statement)
	return err
}

func (u *User) CreateUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO user(id_bin, email) VALUES ('%d', '%s')", u.IDBin, u.Email)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	// Verify insertion
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.IDBin)

	if err != nil {
		return err
	}

	return nil
}
