package model

import (
	"database/sql"
	"log"
)

type Usertype struct {
	Username string
	Password string
}

func LoginUser(username, password string) (*Usertype, error) {
	result := &Usertype{}
	row := db.QueryRow(`
		SELECT USERNAME, PASSWORD
		FROM USER
		WHERE USERNAME=$1
	`, username)
	err := row.Scan(&result.Username, &result.Password)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			log.Println("No rows to return")
			return nil, nil
		default:
			log.Println("Some error occurred")
			return nil, err
		}
	}
	if result.Password == password {
		return result, nil
	} else {
		log.Println("Wrong credentials")
		return nil, nil
	}

}

func RegisterUser(username, password string) (*Usertype, error) {
	result := &Usertype{}
	row := db.QueryRow(`
		SELECT USERNAME, PASSWORD
		FROM USER
		WHERE USERNAME=$1
	`, username)
	err := row.Scan(&result.Username, &result.Password)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			log.Println("No rows to return")
			return nil, nil
		default:
			log.Println("Some error occurred")
			return nil, err
		}
	}
	if result != nil {
		log.Println("User already exists")
		return nil, nil
	} else {
		newrow := db.QueryRow(`
			INSERT INTO USER VALUES($1,$2)
		`, username, password)
		err = newrow.Scan(&result.Username, &result.Password)
		if err == nil {
			return result, nil
		} else {
			return nil, err
		}
	}

}

func NewUserType() *Usertype {
	return &Usertype{}
}
